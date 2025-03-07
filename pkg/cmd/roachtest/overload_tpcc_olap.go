// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/cockroach/pkg/util/timeutil"
)

// tpccOlapQuery is a contrived query that seems to do serious damage to a
// cluster. The query itself is a hash join with a selective filter and a
// limited sort.
const tpccOlapQuery = `SELECT
    i_id, s_w_id, s_quantity, i_price
FROM
    stock JOIN item ON s_i_id = i_id
WHERE
    s_quantity < 100 AND i_price > 90
ORDER BY
    i_price DESC, s_quantity ASC
LIMIT
    100;`

type tpccOLAPSpec struct {
	Nodes       int
	CPUs        int
	Warehouses  int
	Concurrency int
}

func (s tpccOLAPSpec) run(ctx context.Context, t *test, c *cluster) {
	crdbNodes, workloadNode := setupTPCC(ctx, t, c, s.Warehouses, false /* zfs */, nil /* versions */)
	const queryFileName = "queries.sql"
	// querybench expects the entire query to be on a single line.
	queryLine := `"` + strings.Replace(tpccOlapQuery, "\n", " ", -1) + `"`
	c.Run(ctx, workloadNode, "echo", queryLine, "> "+queryFileName)
	t.Status("waiting")
	m := newMonitor(ctx, c, crdbNodes)
	rampDuration := time.Minute
	duration := 2 * time.Minute
	m.Go(func(ctx context.Context) error {
		t.WorkerStatus("running querybench")
		cmd := fmt.Sprintf(
			"./workload run querybench --db tpcc"+
				" --tolerate-errors=t"+
				" --concurrency=%d"+
				" --query-file %s"+
				" --histograms="+perfArtifactsDir+"/stats.json "+
				" --ramp=%s --duration=%s {pgurl:1-%d}",
			s.Warehouses, queryFileName, rampDuration, duration, c.spec.NodeCount-1)
		c.Run(ctx, workloadNode, cmd)
		return nil
	})
	m.Wait()
	verifyNodeLiveness(ctx, c, t, duration)
}

// Check that node liveness did not fail more than maxFailures times across
// all of the nodes.
func verifyNodeLiveness(ctx context.Context, c *cluster, t *test, runDuration time.Duration) {
	const maxFailures = 10
	adminURLs := c.ExternalAdminUIAddr(ctx, c.Node(1))
	now := timeutil.Now()
	response := getMetrics(t, adminURLs[0], now.Add(-runDuration), now, []tsQuery{
		{
			name:      "cr.node.liveness.heartbeatfailures",
			queryType: total,
		},
	})
	if len(response.Results[0].Datapoints) <= 1 {
		t.Fatalf("not enough datapoints in timeseries query response: %+v", response)
	}
	datapoints := response.Results[0].Datapoints
	finalCount := datapoints[len(datapoints)-1].Value
	initialCount := datapoints[0].Value
	if failures := finalCount - initialCount; failures > maxFailures {
		t.Fatalf("Node liveness failed %d times, expected no more than %d",
			failures, maxFailures)
	} else {
		t.logger().Printf("Node liveness failed %d times which is fewer than %d",
			failures, maxFailures)
	}
}

func registerTPCCOverloadSpec(r *testRegistry, s tpccOLAPSpec) {
	name := fmt.Sprintf("overload/tpcc_olap/nodes=%d/cpu=%d/w=%d/c=%d",
		s.Nodes, s.CPUs, s.Warehouses, s.Concurrency)
	r.Add(testSpec{
		Name:       name,
		Cluster:    makeClusterSpec(s.Nodes+1, cpu(s.CPUs)),
		Run:        s.run,
		MinVersion: "v19.2.0",
		Timeout:    20 * time.Minute,
		Skip:       "fails due to node liveness starvation",
	})
}

func registerOverload(r *testRegistry) {
	specs := []tpccOLAPSpec{
		{
			CPUs:        16,
			Concurrency: 256,
			Nodes:       3,
			Warehouses:  100,
		},
	}
	for _, s := range specs {
		registerTPCCOverloadSpec(r, s)
	}
}
