### Array functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>array_append(array: <a href="bool.html">bool</a>[], elem: <a href="bool.html">bool</a>) &rarr; <a href="bool.html">bool</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="bytes.html">bytes</a>[], elem: <a href="bytes.html">bytes</a>) &rarr; <a href="bytes.html">bytes</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="date.html">date</a>[], elem: <a href="date.html">date</a>) &rarr; <a href="date.html">date</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="decimal.html">decimal</a>[], elem: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="float.html">float</a>[], elem: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="inet.html">inet</a>[], elem: <a href="inet.html">inet</a>) &rarr; <a href="inet.html">inet</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="int.html">int</a>[], elem: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="interval.html">interval</a>[], elem: <a href="interval.html">interval</a>) &rarr; <a href="interval.html">interval</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="string.html">string</a>[], elem: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="time.html">time</a>[], elem: <a href="time.html">time</a>) &rarr; <a href="time.html">time</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="timestamp.html">timestamp</a>[], elem: <a href="timestamp.html">timestamp</a>) &rarr; <a href="timestamp.html">timestamp</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="timestamp.html">timestamptz</a>[], elem: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="timestamp.html">timestamptz</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: <a href="uuid.html">uuid</a>[], elem: <a href="uuid.html">uuid</a>) &rarr; <a href="uuid.html">uuid</a>[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: oid[], elem: oid) &rarr; oid[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_append(array: varbit[], elem: varbit) &rarr; varbit[]</code></td><td><span class="funcdesc"><p>Appends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="bool.html">bool</a>[], right: <a href="bool.html">bool</a>[]) &rarr; <a href="bool.html">bool</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="bytes.html">bytes</a>[], right: <a href="bytes.html">bytes</a>[]) &rarr; <a href="bytes.html">bytes</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="date.html">date</a>[], right: <a href="date.html">date</a>[]) &rarr; <a href="date.html">date</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="decimal.html">decimal</a>[], right: <a href="decimal.html">decimal</a>[]) &rarr; <a href="decimal.html">decimal</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="float.html">float</a>[], right: <a href="float.html">float</a>[]) &rarr; <a href="float.html">float</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="inet.html">inet</a>[], right: <a href="inet.html">inet</a>[]) &rarr; <a href="inet.html">inet</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="int.html">int</a>[], right: <a href="int.html">int</a>[]) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="interval.html">interval</a>[], right: <a href="interval.html">interval</a>[]) &rarr; <a href="interval.html">interval</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="string.html">string</a>[], right: <a href="string.html">string</a>[]) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="time.html">time</a>[], right: <a href="time.html">time</a>[]) &rarr; <a href="time.html">time</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="timestamp.html">timestamp</a>[], right: <a href="timestamp.html">timestamp</a>[]) &rarr; <a href="timestamp.html">timestamp</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="timestamp.html">timestamptz</a>[], right: <a href="timestamp.html">timestamptz</a>[]) &rarr; <a href="timestamp.html">timestamptz</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: <a href="uuid.html">uuid</a>[], right: <a href="uuid.html">uuid</a>[]) &rarr; <a href="uuid.html">uuid</a>[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: oid[], right: oid[]) &rarr; oid[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_cat(left: varbit[], right: varbit[]) &rarr; varbit[]</code></td><td><span class="funcdesc"><p>Appends two arrays.</p>
</span></td></tr>
<tr><td><code>array_length(input: anyelement[], array_dimension: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the length of <code>input</code> on the provided <code>array_dimension</code>. However, because CockroachDB doesn’t yet support multi-dimensional arrays, the only supported <code>array_dimension</code> is <strong>1</strong>.</p>
</span></td></tr>
<tr><td><code>array_lower(input: anyelement[], array_dimension: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the minimum value of <code>input</code> on the provided <code>array_dimension</code>. However, because CockroachDB doesn’t yet support multi-dimensional arrays, the only supported <code>array_dimension</code> is <strong>1</strong>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="bool.html">bool</a>[], elem: <a href="bool.html">bool</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="bytes.html">bytes</a>[], elem: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="date.html">date</a>[], elem: <a href="date.html">date</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="decimal.html">decimal</a>[], elem: <a href="decimal.html">decimal</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="float.html">float</a>[], elem: <a href="float.html">float</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="inet.html">inet</a>[], elem: <a href="inet.html">inet</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="int.html">int</a>[], elem: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="interval.html">interval</a>[], elem: <a href="interval.html">interval</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="string.html">string</a>[], elem: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="time.html">time</a>[], elem: <a href="time.html">time</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="timestamp.html">timestamp</a>[], elem: <a href="timestamp.html">timestamp</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="timestamp.html">timestamptz</a>[], elem: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: <a href="uuid.html">uuid</a>[], elem: <a href="uuid.html">uuid</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: oid[], elem: oid) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_position(array: varbit[], elem: varbit) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return the index of the first occurrence of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="bool.html">bool</a>[], elem: <a href="bool.html">bool</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="bytes.html">bytes</a>[], elem: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="date.html">date</a>[], elem: <a href="date.html">date</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="decimal.html">decimal</a>[], elem: <a href="decimal.html">decimal</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="float.html">float</a>[], elem: <a href="float.html">float</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="inet.html">inet</a>[], elem: <a href="inet.html">inet</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="int.html">int</a>[], elem: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="interval.html">interval</a>[], elem: <a href="interval.html">interval</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="string.html">string</a>[], elem: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="time.html">time</a>[], elem: <a href="time.html">time</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="timestamp.html">timestamp</a>[], elem: <a href="timestamp.html">timestamp</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="timestamp.html">timestamptz</a>[], elem: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: <a href="uuid.html">uuid</a>[], elem: <a href="uuid.html">uuid</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: oid[], elem: oid) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_positions(array: varbit[], elem: varbit) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Returns and array of indexes of all occurrences of <code>elem</code> in <code>array</code>.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="bool.html">bool</a>, array: <a href="bool.html">bool</a>[]) &rarr; <a href="bool.html">bool</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="bytes.html">bytes</a>, array: <a href="bytes.html">bytes</a>[]) &rarr; <a href="bytes.html">bytes</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="date.html">date</a>, array: <a href="date.html">date</a>[]) &rarr; <a href="date.html">date</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="decimal.html">decimal</a>, array: <a href="decimal.html">decimal</a>[]) &rarr; <a href="decimal.html">decimal</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="float.html">float</a>, array: <a href="float.html">float</a>[]) &rarr; <a href="float.html">float</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="inet.html">inet</a>, array: <a href="inet.html">inet</a>[]) &rarr; <a href="inet.html">inet</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="int.html">int</a>, array: <a href="int.html">int</a>[]) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="interval.html">interval</a>, array: <a href="interval.html">interval</a>[]) &rarr; <a href="interval.html">interval</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="string.html">string</a>, array: <a href="string.html">string</a>[]) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="time.html">time</a>, array: <a href="time.html">time</a>[]) &rarr; <a href="time.html">time</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="timestamp.html">timestamp</a>, array: <a href="timestamp.html">timestamp</a>[]) &rarr; <a href="timestamp.html">timestamp</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="timestamp.html">timestamptz</a>, array: <a href="timestamp.html">timestamptz</a>[]) &rarr; <a href="timestamp.html">timestamptz</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: <a href="uuid.html">uuid</a>, array: <a href="uuid.html">uuid</a>[]) &rarr; <a href="uuid.html">uuid</a>[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: oid, array: oid[]) &rarr; oid[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_prepend(elem: varbit, array: varbit[]) &rarr; varbit[]</code></td><td><span class="funcdesc"><p>Prepends <code>elem</code> to <code>array</code>, returning the result.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="bool.html">bool</a>[], elem: <a href="bool.html">bool</a>) &rarr; <a href="bool.html">bool</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="bytes.html">bytes</a>[], elem: <a href="bytes.html">bytes</a>) &rarr; <a href="bytes.html">bytes</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="date.html">date</a>[], elem: <a href="date.html">date</a>) &rarr; <a href="date.html">date</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="decimal.html">decimal</a>[], elem: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="float.html">float</a>[], elem: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="inet.html">inet</a>[], elem: <a href="inet.html">inet</a>) &rarr; <a href="inet.html">inet</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="int.html">int</a>[], elem: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="interval.html">interval</a>[], elem: <a href="interval.html">interval</a>) &rarr; <a href="interval.html">interval</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="string.html">string</a>[], elem: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="time.html">time</a>[], elem: <a href="time.html">time</a>) &rarr; <a href="time.html">time</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="timestamp.html">timestamp</a>[], elem: <a href="timestamp.html">timestamp</a>) &rarr; <a href="timestamp.html">timestamp</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="timestamp.html">timestamptz</a>[], elem: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="timestamp.html">timestamptz</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: <a href="uuid.html">uuid</a>[], elem: <a href="uuid.html">uuid</a>) &rarr; <a href="uuid.html">uuid</a>[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: oid[], elem: oid) &rarr; oid[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_remove(array: varbit[], elem: varbit) &rarr; varbit[]</code></td><td><span class="funcdesc"><p>Remove from <code>array</code> all elements equal to <code>elem</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="bool.html">bool</a>[], toreplace: <a href="bool.html">bool</a>, replacewith: <a href="bool.html">bool</a>) &rarr; <a href="bool.html">bool</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="bytes.html">bytes</a>[], toreplace: <a href="bytes.html">bytes</a>, replacewith: <a href="bytes.html">bytes</a>) &rarr; <a href="bytes.html">bytes</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="date.html">date</a>[], toreplace: <a href="date.html">date</a>, replacewith: <a href="date.html">date</a>) &rarr; <a href="date.html">date</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="decimal.html">decimal</a>[], toreplace: <a href="decimal.html">decimal</a>, replacewith: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="float.html">float</a>[], toreplace: <a href="float.html">float</a>, replacewith: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="inet.html">inet</a>[], toreplace: <a href="inet.html">inet</a>, replacewith: <a href="inet.html">inet</a>) &rarr; <a href="inet.html">inet</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="int.html">int</a>[], toreplace: <a href="int.html">int</a>, replacewith: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="interval.html">interval</a>[], toreplace: <a href="interval.html">interval</a>, replacewith: <a href="interval.html">interval</a>) &rarr; <a href="interval.html">interval</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="string.html">string</a>[], toreplace: <a href="string.html">string</a>, replacewith: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="time.html">time</a>[], toreplace: <a href="time.html">time</a>, replacewith: <a href="time.html">time</a>) &rarr; <a href="time.html">time</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="timestamp.html">timestamp</a>[], toreplace: <a href="timestamp.html">timestamp</a>, replacewith: <a href="timestamp.html">timestamp</a>) &rarr; <a href="timestamp.html">timestamp</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="timestamp.html">timestamptz</a>[], toreplace: <a href="timestamp.html">timestamptz</a>, replacewith: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="timestamp.html">timestamptz</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: <a href="uuid.html">uuid</a>[], toreplace: <a href="uuid.html">uuid</a>, replacewith: <a href="uuid.html">uuid</a>) &rarr; <a href="uuid.html">uuid</a>[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: oid[], toreplace: oid, replacewith: oid) &rarr; oid[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_replace(array: varbit[], toreplace: varbit, replacewith: varbit) &rarr; varbit[]</code></td><td><span class="funcdesc"><p>Replace all occurrences of <code>toreplace</code> in <code>array</code> with <code>replacewith</code>.</p>
</span></td></tr>
<tr><td><code>array_to_string(input: anyelement[], delim: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Join an array into a string with a delimiter.</p>
</span></td></tr>
<tr><td><code>array_to_string(input: anyelement[], delimiter: <a href="string.html">string</a>, null: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Join an array into a string with a delimiter, replacing NULLs with a null string.</p>
</span></td></tr>
<tr><td><code>array_upper(input: anyelement[], array_dimension: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the maximum value of <code>input</code> on the provided <code>array_dimension</code>. However, because CockroachDB doesn’t yet support multi-dimensional arrays, the only supported <code>array_dimension</code> is <strong>1</strong>.</p>
</span></td></tr>
<tr><td><code>string_to_array(str: <a href="string.html">string</a>, delimiter: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Split a string into components on a delimiter.</p>
</span></td></tr>
<tr><td><code>string_to_array(str: <a href="string.html">string</a>, delimiter: <a href="string.html">string</a>, null: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Split a string into components on a delimiter with a specified string to consider NULL.</p>
</span></td></tr></tbody>
</table>

### BOOL functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>ilike_escape(unescaped: <a href="string.html">string</a>, pattern: <a href="string.html">string</a>, escape: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Matches case insensetively <code>unescaped</code> with <code>pattern</code> using ‘escape’ as an escape token.</p>
</span></td></tr>
<tr><td><code>inet_contained_by_or_equals(val: <a href="inet.html">inet</a>, container: <a href="inet.html">inet</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Test for subnet inclusion or equality, using only the network parts of the addresses. The host part of the addresses is ignored.</p>
</span></td></tr>
<tr><td><code>inet_contains_or_contained_by(val: <a href="inet.html">inet</a>, val: <a href="inet.html">inet</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Test for subnet inclusion, using only the network parts of the addresses. The host part of the addresses is ignored.</p>
</span></td></tr>
<tr><td><code>inet_contains_or_equals(container: <a href="inet.html">inet</a>, val: <a href="inet.html">inet</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Test for subnet inclusion or equality, using only the network parts of the addresses. The host part of the addresses is ignored.</p>
</span></td></tr>
<tr><td><code>inet_same_family(val: <a href="inet.html">inet</a>, val: <a href="inet.html">inet</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Checks if two IP addresses are of the same IP family.</p>
</span></td></tr>
<tr><td><code>like_escape(unescaped: <a href="string.html">string</a>, pattern: <a href="string.html">string</a>, escape: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Matches <code>unescaped</code> with <code>pattern</code> using ‘escape’ as an escape token.</p>
</span></td></tr>
<tr><td><code>not_ilike_escape(unescaped: <a href="string.html">string</a>, pattern: <a href="string.html">string</a>, escape: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Checks whether <code>unescaped</code> not matches case insensetively with <code>pattern</code> using ‘escape’ as an escape token.</p>
</span></td></tr>
<tr><td><code>not_like_escape(unescaped: <a href="string.html">string</a>, pattern: <a href="string.html">string</a>, escape: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Checks whether <code>unescaped</code> not matches with <code>pattern</code> using ‘escape’ as an escape token.</p>
</span></td></tr>
<tr><td><code>not_similar_to_escape(unescaped: <a href="string.html">string</a>, pattern: <a href="string.html">string</a>, escape: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Checks whether <code>unescaped</code> not matches with <code>pattern</code> using ‘escape’ as an escape token.</p>
</span></td></tr>
<tr><td><code>similar_to_escape(unescaped: <a href="string.html">string</a>, pattern: <a href="string.html">string</a>, escape: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Matches <code>unescaped</code> with <code>pattern</code> using ‘escape’ as an escape token.</p>
</span></td></tr></tbody>
</table>

### Comparison functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>greatest(anyelement...) &rarr; anyelement</code></td><td><span class="funcdesc"><p>Returns the element with the greatest value.</p>
</span></td></tr>
<tr><td><code>least(anyelement...) &rarr; anyelement</code></td><td><span class="funcdesc"><p>Returns the element with the lowest value.</p>
</span></td></tr></tbody>
</table>

### Date and time functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>age(end: <a href="timestamp.html">timestamptz</a>, begin: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="interval.html">interval</a></code></td><td><span class="funcdesc"><p>Calculates the interval between <code>begin</code> and <code>end</code>.</p>
</span></td></tr>
<tr><td><code>age(val: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="interval.html">interval</a></code></td><td><span class="funcdesc"><p>Calculates the interval between <code>val</code> and the current time.</p>
</span></td></tr>
<tr><td><code>clock_timestamp() &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Returns the current system time on one of the cluster nodes.</p>
</span></td></tr>
<tr><td><code>clock_timestamp() &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Returns the current system time on one of the cluster nodes.</p>
</span></td></tr>
<tr><td><code>current_date() &rarr; <a href="date.html">date</a></code></td><td><span class="funcdesc"><p>Returns the date of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>current_timestamp() &rarr; <a href="date.html">date</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>current_timestamp() &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>current_timestamp() &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>date_trunc(element: <a href="string.html">string</a>, input: <a href="date.html">date</a>) &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Truncates <code>input</code> to precision <code>element</code>.  Sets all fields that are less
significant than <code>element</code> to zero (or one, for day and month)</p>
<p>Compatible elements: year, quarter, month, week, hour, minute, second,
millisecond, microsecond.</p>
</span></td></tr>
<tr><td><code>date_trunc(element: <a href="string.html">string</a>, input: <a href="time.html">time</a>) &rarr; <a href="interval.html">interval</a></code></td><td><span class="funcdesc"><p>Truncates <code>input</code> to precision <code>element</code>.  Sets all fields that are less
significant than <code>element</code> to zero.</p>
<p>Compatible elements: hour, minute, second, millisecond, microsecond.</p>
</span></td></tr>
<tr><td><code>date_trunc(element: <a href="string.html">string</a>, input: <a href="timestamp.html">timestamp</a>) &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Truncates <code>input</code> to precision <code>element</code>.  Sets all fields that are less
significant than <code>element</code> to zero (or one, for day and month)</p>
<p>Compatible elements: year, quarter, month, week, hour, minute, second,
millisecond, microsecond.</p>
</span></td></tr>
<tr><td><code>date_trunc(element: <a href="string.html">string</a>, input: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Truncates <code>input</code> to precision <code>element</code>.  Sets all fields that are less
significant than <code>element</code> to zero (or one, for day and month)</p>
<p>Compatible elements: year, quarter, month, week, hour, minute, second,
millisecond, microsecond.</p>
</span></td></tr>
<tr><td><code>experimental_follower_read_timestamp() &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Returns a timestamp which is very likely to be safe to perform
against a follower replica.</p>
<p>This function is intended to be used with an AS OF SYSTEM TIME clause to perform
historical reads against a time which is recent but sufficiently old for reads
to be performed against the closest replica as opposed to the currently
leaseholder for a given range.</p>
<p>Note that this function requires an enterprise license on a CCL distribution to
return without an error.</p>
</span></td></tr>
<tr><td><code>experimental_strftime(input: <a href="date.html">date</a>, extract_format: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>From <code>input</code>, extracts and formats the time as identified in <code>extract_format</code> using standard <code>strftime</code> notation (though not all formatting is supported).</p>
</span></td></tr>
<tr><td><code>experimental_strftime(input: <a href="timestamp.html">timestamp</a>, extract_format: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>From <code>input</code>, extracts and formats the time as identified in <code>extract_format</code> using standard <code>strftime</code> notation (though not all formatting is supported).</p>
</span></td></tr>
<tr><td><code>experimental_strftime(input: <a href="timestamp.html">timestamptz</a>, extract_format: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>From <code>input</code>, extracts and formats the time as identified in <code>extract_format</code> using standard <code>strftime</code> notation (though not all formatting is supported).</p>
</span></td></tr>
<tr><td><code>experimental_strptime(input: <a href="string.html">string</a>, format: <a href="string.html">string</a>) &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Returns <code>input</code> as a timestamptz using <code>format</code> (which uses standard <code>strptime</code> formatting).</p>
</span></td></tr>
<tr><td><code>extract(element: <a href="string.html">string</a>, input: <a href="date.html">date</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Extracts <code>element</code> from <code>input</code>.</p>
<p>Compatible elements: year, quarter, month, week, dayofweek, dayofyear,
hour, minute, second, millisecond, microsecond, epoch</p>
</span></td></tr>
<tr><td><code>extract(element: <a href="string.html">string</a>, input: <a href="time.html">time</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Extracts <code>element</code> from <code>input</code>.</p>
<p>Compatible elements: hour, minute, second, millisecond, microsecond, epoch</p>
</span></td></tr>
<tr><td><code>extract(element: <a href="string.html">string</a>, input: <a href="timestamp.html">timestamp</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Extracts <code>element</code> from <code>input</code>.</p>
<p>Compatible elements: year, quarter, month, week, dayofweek, dayofyear,
hour, minute, second, millisecond, microsecond, epoch</p>
</span></td></tr>
<tr><td><code>extract(element: <a href="string.html">string</a>, input: <a href="timestamp.html">timestamptz</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Extracts <code>element</code> from <code>input</code>.</p>
<p>Compatible elements: year, quarter, month, week, dayofweek, dayofyear,
hour, minute, second, millisecond, microsecond, epoch</p>
</span></td></tr>
<tr><td><code>extract_duration(element: <a href="string.html">string</a>, input: <a href="interval.html">interval</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Extracts <code>element</code> from <code>input</code>.
Compatible elements: hour, minute, second, millisecond, microsecond.</p>
</span></td></tr>
<tr><td><code>now() &rarr; <a href="date.html">date</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>now() &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>now() &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>statement_timestamp() &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Returns the start time of the current statement.</p>
</span></td></tr>
<tr><td><code>statement_timestamp() &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Returns the start time of the current statement.</p>
</span></td></tr>
<tr><td><code>timezone(timestamp: <a href="timestamp.html">timestamp</a>, timezone: <a href="string.html">string</a>) &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Treat given time stamp without time zone as located in the specified time zone</p>
</span></td></tr>
<tr><td><code>timezone(timestamptz: <a href="timestamp.html">timestamptz</a>, timezone: <a href="string.html">string</a>) &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Convert given time stamp with time zone to the new time zone, with no time zone designation</p>
</span></td></tr>
<tr><td><code>transaction_timestamp() &rarr; <a href="date.html">date</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>transaction_timestamp() &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr>
<tr><td><code>transaction_timestamp() &rarr; <a href="timestamp.html">timestamptz</a></code></td><td><span class="funcdesc"><p>Returns the time of the current transaction.</p>
<p>The value is based on a timestamp picked when the transaction starts
and which stays constant throughout the transaction. This timestamp
has no relationship with the commit order of concurrent transactions.</p>
</span></td></tr></tbody>
</table>

### ID generation functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>experimental_uuid_v4() &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Returns a UUID.</p>
</span></td></tr>
<tr><td><code>gen_random_uuid() &rarr; <a href="uuid.html">uuid</a></code></td><td><span class="funcdesc"><p>Generates a random UUID and returns it as a value of UUID type.</p>
</span></td></tr>
<tr><td><code>unique_rowid() &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns a unique ID used by CockroachDB to generate unique row IDs if a Primary Key isn’t defined for the table. The value is a combination of the insert timestamp and the ID of the node executing the statement, which guarantees this combination is globally unique. However, there can be gaps and the order is not completely guaranteed.</p>
</span></td></tr>
<tr><td><code>uuid_v4() &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Returns a UUID.</p>
</span></td></tr></tbody>
</table>

### INET functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>abbrev(val: <a href="inet.html">inet</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts the combined IP address and prefix length to an abbreviated display format as text.For INET types, this will omit the prefix length if it’s not the default (32 or IPv4, 128 for IPv6)</p>
<p>For example, <code>abbrev('192.168.1.2/24')</code> returns <code>'192.168.1.2/24'</code></p>
</span></td></tr>
<tr><td><code>broadcast(val: <a href="inet.html">inet</a>) &rarr; <a href="inet.html">inet</a></code></td><td><span class="funcdesc"><p>Gets the broadcast address for the network address represented by the value.</p>
<p>For example, <code>broadcast('192.168.1.2/24')</code> returns <code>'192.168.1.255/24'</code></p>
</span></td></tr>
<tr><td><code>family(val: <a href="inet.html">inet</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Extracts the IP family of the value; 4 for IPv4, 6 for IPv6.</p>
<p>For example, <code>family('::1')</code> returns <code>6</code></p>
</span></td></tr>
<tr><td><code>host(val: <a href="inet.html">inet</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Extracts the address part of the combined address/prefixlen value as text.</p>
<p>For example, <code>host('192.168.1.2/16')</code> returns <code>'192.168.1.2'</code></p>
</span></td></tr>
<tr><td><code>hostmask(val: <a href="inet.html">inet</a>) &rarr; <a href="inet.html">inet</a></code></td><td><span class="funcdesc"><p>Creates an IP host mask corresponding to the prefix length in the value.</p>
<p>For example, <code>hostmask('192.168.1.2/16')</code> returns <code>'0.0.255.255'</code></p>
</span></td></tr>
<tr><td><code>masklen(val: <a href="inet.html">inet</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Retrieves the prefix length stored in the value.</p>
<p>For example, <code>masklen('192.168.1.2/16')</code> returns <code>16</code></p>
</span></td></tr>
<tr><td><code>netmask(val: <a href="inet.html">inet</a>) &rarr; <a href="inet.html">inet</a></code></td><td><span class="funcdesc"><p>Creates an IP network mask corresponding to the prefix length in the value.</p>
<p>For example, <code>netmask('192.168.1.2/16')</code> returns <code>'255.255.0.0'</code></p>
</span></td></tr>
<tr><td><code>set_masklen(val: <a href="inet.html">inet</a>, prefixlen: <a href="int.html">int</a>) &rarr; <a href="inet.html">inet</a></code></td><td><span class="funcdesc"><p>Sets the prefix length of <code>val</code> to <code>prefixlen</code>.</p>
<p>For example, <code>set_masklen('192.168.1.2', 16)</code> returns <code>'192.168.1.2/16'</code>.</p>
</span></td></tr>
<tr><td><code>text(val: <a href="inet.html">inet</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts the IP address and prefix length to text.</p>
</span></td></tr></tbody>
</table>

### JSONB functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>array_to_json(array: anyelement[]) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the array as JSON or JSONB.</p>
</span></td></tr>
<tr><td><code>array_to_json(array: anyelement[], pretty_bool: <a href="bool.html">bool</a>) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the array as JSON or JSONB.</p>
</span></td></tr>
<tr><td><code>json_array_length(json: jsonb) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns the number of elements in the outermost JSON or JSONB array.</p>
</span></td></tr>
<tr><td><code>json_build_array(anyelement...) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Builds a possibly-heterogeneously-typed JSON or JSONB array out of a variadic argument list.</p>
</span></td></tr>
<tr><td><code>json_build_object(anyelement...) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Builds a JSON object out of a variadic argument list.</p>
</span></td></tr>
<tr><td><code>json_extract_path(jsonb, <a href="string.html">string</a>...) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments.</p>
</span></td></tr>
<tr><td><code>json_object(keys: <a href="string.html">string</a>[], values: <a href="string.html">string</a>[]) &rarr; jsonb</code></td><td><span class="funcdesc"><p>This form of json_object takes keys and values pairwise from two separate arrays. In all other respects it is identical to the one-argument form.</p>
</span></td></tr>
<tr><td><code>json_object(texts: <a href="string.html">string</a>[]) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Builds a JSON or JSONB object out of a text array. The array must have exactly one dimension with an even number of members, in which case they are taken as alternating key/value pairs.</p>
</span></td></tr>
<tr><td><code>json_remove_path(val: jsonb, path: <a href="string.html">string</a>[]) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Remove the specified path from the JSON object.</p>
</span></td></tr>
<tr><td><code>json_set(val: jsonb, path: <a href="string.html">string</a>[], to: jsonb) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments.</p>
</span></td></tr>
<tr><td><code>json_set(val: jsonb, path: <a href="string.html">string</a>[], to: jsonb, create_missing: <a href="bool.html">bool</a>) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments. If <code>create_missing</code> is false, new keys will not be inserted to objects and values will not be prepended or appended to arrays.</p>
</span></td></tr>
<tr><td><code>json_strip_nulls(from_json: jsonb) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns from_json with all object fields that have null values omitted. Other null values are untouched.</p>
</span></td></tr>
<tr><td><code>json_typeof(val: jsonb) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the type of the outermost JSON value as a text string.</p>
</span></td></tr>
<tr><td><code>jsonb_array_length(json: jsonb) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns the number of elements in the outermost JSON or JSONB array.</p>
</span></td></tr>
<tr><td><code>jsonb_build_array(anyelement...) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Builds a possibly-heterogeneously-typed JSON or JSONB array out of a variadic argument list.</p>
</span></td></tr>
<tr><td><code>jsonb_build_object(anyelement...) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Builds a JSON object out of a variadic argument list.</p>
</span></td></tr>
<tr><td><code>jsonb_extract_path(jsonb, <a href="string.html">string</a>...) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments.</p>
</span></td></tr>
<tr><td><code>jsonb_insert(target: jsonb, path: <a href="string.html">string</a>[], new_val: jsonb) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments. <code>new_val</code> will be inserted before path target.</p>
</span></td></tr>
<tr><td><code>jsonb_insert(target: jsonb, path: <a href="string.html">string</a>[], new_val: jsonb, insert_after: <a href="bool.html">bool</a>) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments. If <code>insert_after</code> is true (default is false), <code>new_val</code> will be inserted after path target.</p>
</span></td></tr>
<tr><td><code>jsonb_object(keys: <a href="string.html">string</a>[], values: <a href="string.html">string</a>[]) &rarr; jsonb</code></td><td><span class="funcdesc"><p>This form of json_object takes keys and values pairwise from two separate arrays. In all other respects it is identical to the one-argument form.</p>
</span></td></tr>
<tr><td><code>jsonb_object(texts: <a href="string.html">string</a>[]) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Builds a JSON or JSONB object out of a text array. The array must have exactly one dimension with an even number of members, in which case they are taken as alternating key/value pairs.</p>
</span></td></tr>
<tr><td><code>jsonb_pretty(val: jsonb) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the given JSON value as a STRING indented and with newlines.</p>
</span></td></tr>
<tr><td><code>jsonb_set(val: jsonb, path: <a href="string.html">string</a>[], to: jsonb) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments.</p>
</span></td></tr>
<tr><td><code>jsonb_set(val: jsonb, path: <a href="string.html">string</a>[], to: jsonb, create_missing: <a href="bool.html">bool</a>) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the JSON value pointed to by the variadic arguments. If <code>create_missing</code> is false, new keys will not be inserted to objects and values will not be prepended or appended to arrays.</p>
</span></td></tr>
<tr><td><code>jsonb_strip_nulls(from_json: jsonb) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns from_json with all object fields that have null values omitted. Other null values are untouched.</p>
</span></td></tr>
<tr><td><code>jsonb_typeof(val: jsonb) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the type of the outermost JSON value as a text string.</p>
</span></td></tr>
<tr><td><code>to_json(val: anyelement) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the value as JSON or JSONB.</p>
</span></td></tr>
<tr><td><code>to_jsonb(val: anyelement) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the value as JSON or JSONB.</p>
</span></td></tr></tbody>
</table>

### Math and numeric functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>abs(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the absolute value of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>abs(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the absolute value of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>abs(val: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the absolute value of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>acos(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the inverse cosine of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>asin(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the inverse sine of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>atan(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the inverse tangent of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>atan2(x: <a href="float.html">float</a>, y: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the inverse tangent of <code>x</code>/<code>y</code>.</p>
</span></td></tr>
<tr><td><code>cbrt(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the cube root (∛) of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>cbrt(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the cube root (∛) of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>ceil(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the smallest integer greater than <code>val</code>.</p>
</span></td></tr>
<tr><td><code>ceil(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the smallest integer greater than <code>val</code>.</p>
</span></td></tr>
<tr><td><code>ceiling(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the smallest integer greater than <code>val</code>.</p>
</span></td></tr>
<tr><td><code>ceiling(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the smallest integer greater than <code>val</code>.</p>
</span></td></tr>
<tr><td><code>cos(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the cosine of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>cot(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the cotangent of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>crc32c(<a href="bytes.html">bytes</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the CRC-32 hash using the Castagnoli polynomial.</p>
</span></td></tr>
<tr><td><code>crc32c(<a href="string.html">string</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the CRC-32 hash using the Castagnoli polynomial.</p>
</span></td></tr>
<tr><td><code>crc32ieee(<a href="bytes.html">bytes</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the CRC-32 hash using the IEEE polynomial.</p>
</span></td></tr>
<tr><td><code>crc32ieee(<a href="string.html">string</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the CRC-32 hash using the IEEE polynomial.</p>
</span></td></tr>
<tr><td><code>degrees(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Converts <code>val</code> as a radian value to a degree value.</p>
</span></td></tr>
<tr><td><code>div(x: <a href="decimal.html">decimal</a>, y: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the integer quotient of <code>x</code>/<code>y</code>.</p>
</span></td></tr>
<tr><td><code>div(x: <a href="float.html">float</a>, y: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the integer quotient of <code>x</code>/<code>y</code>.</p>
</span></td></tr>
<tr><td><code>div(x: <a href="int.html">int</a>, y: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the integer quotient of <code>x</code>/<code>y</code>.</p>
</span></td></tr>
<tr><td><code>exp(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates <em>e</em> ^ <code>val</code>.</p>
</span></td></tr>
<tr><td><code>exp(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates <em>e</em> ^ <code>val</code>.</p>
</span></td></tr>
<tr><td><code>floor(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the largest integer not greater than <code>val</code>.</p>
</span></td></tr>
<tr><td><code>floor(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the largest integer not greater than <code>val</code>.</p>
</span></td></tr>
<tr><td><code>fnv32(<a href="bytes.html">bytes</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 32-bit FNV-1 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>fnv32(<a href="string.html">string</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 32-bit FNV-1 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>fnv32a(<a href="bytes.html">bytes</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 32-bit FNV-1a hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>fnv32a(<a href="string.html">string</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 32-bit FNV-1a hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>fnv64(<a href="bytes.html">bytes</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 64-bit FNV-1 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>fnv64(<a href="string.html">string</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 64-bit FNV-1 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>fnv64a(<a href="bytes.html">bytes</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 64-bit FNV-1a hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>fnv64a(<a href="string.html">string</a>...) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the 64-bit FNV-1a hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>isnan(val: <a href="decimal.html">decimal</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns true if <code>val</code> is NaN, false otherwise.</p>
</span></td></tr>
<tr><td><code>isnan(val: <a href="float.html">float</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns true if <code>val</code> is NaN, false otherwise.</p>
</span></td></tr>
<tr><td><code>ln(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the natural log of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>ln(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the natural log of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>log(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the base 10 log of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>log(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the base 10 log of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>mod(x: <a href="decimal.html">decimal</a>, y: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>%<code>y</code>.</p>
</span></td></tr>
<tr><td><code>mod(x: <a href="float.html">float</a>, y: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>%<code>y</code>.</p>
</span></td></tr>
<tr><td><code>mod(x: <a href="int.html">int</a>, y: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>%<code>y</code>.</p>
</span></td></tr>
<tr><td><code>pi() &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Returns the value for pi (3.141592653589793).</p>
</span></td></tr>
<tr><td><code>pow(x: <a href="decimal.html">decimal</a>, y: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>^<code>y</code>.</p>
</span></td></tr>
<tr><td><code>pow(x: <a href="float.html">float</a>, y: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>^<code>y</code>.</p>
</span></td></tr>
<tr><td><code>pow(x: <a href="int.html">int</a>, y: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>^<code>y</code>.</p>
</span></td></tr>
<tr><td><code>power(x: <a href="decimal.html">decimal</a>, y: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>^<code>y</code>.</p>
</span></td></tr>
<tr><td><code>power(x: <a href="float.html">float</a>, y: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>^<code>y</code>.</p>
</span></td></tr>
<tr><td><code>power(x: <a href="int.html">int</a>, y: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates <code>x</code>^<code>y</code>.</p>
</span></td></tr>
<tr><td><code>radians(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Converts <code>val</code> as a degree value to a radians value.</p>
</span></td></tr>
<tr><td><code>random() &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Returns a random float between 0 and 1.</p>
</span></td></tr>
<tr><td><code>round(input: <a href="decimal.html">decimal</a>, decimal_accuracy: <a href="int.html">int</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Keeps <code>decimal_accuracy</code> number of figures to the right of the zero position in <code>input</code> using half away from zero rounding. If <code>decimal_accuracy</code> is not in the range -2^31…(2^31-1), the results are undefined.</p>
</span></td></tr>
<tr><td><code>round(input: <a href="float.html">float</a>, decimal_accuracy: <a href="int.html">int</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Keeps <code>decimal_accuracy</code> number of figures to the right of the zero position  in <code>input</code> using half to even (banker’s) rounding.</p>
</span></td></tr>
<tr><td><code>round(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Rounds <code>val</code> to the nearest integer, half away from zero: round(+/-2.4) = +/-2, round(+/-2.5) = +/-3.</p>
</span></td></tr>
<tr><td><code>round(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Rounds <code>val</code> to the nearest integer using half to even (banker’s) rounding.</p>
</span></td></tr>
<tr><td><code>sign(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Determines the sign of <code>val</code>: <strong>1</strong> for positive; <strong>0</strong> for 0 values; <strong>-1</strong> for negative.</p>
</span></td></tr>
<tr><td><code>sign(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Determines the sign of <code>val</code>: <strong>1</strong> for positive; <strong>0</strong> for 0 values; <strong>-1</strong> for negative.</p>
</span></td></tr>
<tr><td><code>sign(val: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Determines the sign of <code>val</code>: <strong>1</strong> for positive; <strong>0</strong> for 0 values; <strong>-1</strong> for negative.</p>
</span></td></tr>
<tr><td><code>sin(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the sine of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>sqrt(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Calculates the square root of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>sqrt(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the square root of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>tan(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Calculates the tangent of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>trunc(val: <a href="decimal.html">decimal</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Truncates the decimal values of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>trunc(val: <a href="float.html">float</a>) &rarr; <a href="float.html">float</a></code></td><td><span class="funcdesc"><p>Truncates the decimal values of <code>val</code>.</p>
</span></td></tr></tbody>
</table>

### Sequence functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>currval(sequence_name: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns the latest value obtained with nextval for this sequence in this session.</p>
</span></td></tr>
<tr><td><code>lastval() &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Return value most recently obtained with nextval in this session.</p>
</span></td></tr>
<tr><td><code>nextval(sequence_name: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Advances the given sequence and returns its new value.</p>
</span></td></tr>
<tr><td><code>setval(sequence_name: <a href="string.html">string</a>, value: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Set the given sequence’s current value. The next call to nextval will return <code>value + Increment</code></p>
</span></td></tr>
<tr><td><code>setval(sequence_name: <a href="string.html">string</a>, value: <a href="int.html">int</a>, is_called: <a href="bool.html">bool</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Set the given sequence’s current value. If is_called is false, the next call to nextval will return <code>value</code>; otherwise <code>value + Increment</code>.</p>
</span></td></tr></tbody>
</table>

### Set-returning functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>aclexplode(aclitems: <a href="string.html">string</a>[]) &rarr; tuple{oid AS grantor, oid AS grantee, string AS privilege_type, bool AS is_grantable}</code></td><td><span class="funcdesc"><p>Produces a virtual table containing aclitem stuff (returns no rows as this feature is unsupported in CockroachDB)</p>
</span></td></tr>
<tr><td><code>crdb_internal.unary_table() &rarr; tuple</code></td><td><span class="funcdesc"><p>Produces a virtual table containing a single row with no values.</p>
<p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>generate_series(start: <a href="int.html">int</a>, end: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Produces a virtual table containing the integer values from <code>start</code> to <code>end</code>, inclusive.</p>
</span></td></tr>
<tr><td><code>generate_series(start: <a href="int.html">int</a>, end: <a href="int.html">int</a>, step: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Produces a virtual table containing the integer values from <code>start</code> to <code>end</code>, inclusive, by increment of <code>step</code>.</p>
</span></td></tr>
<tr><td><code>generate_series(start: <a href="timestamp.html">timestamp</a>, end: <a href="timestamp.html">timestamp</a>, step: <a href="interval.html">interval</a>) &rarr; <a href="timestamp.html">timestamp</a></code></td><td><span class="funcdesc"><p>Produces a virtual table containing the timestamp values from <code>start</code> to <code>end</code>, inclusive, by increment of <code>step</code>.</p>
</span></td></tr>
<tr><td><code>generate_subscripts(array: anyelement[]) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns a series comprising the given array’s subscripts.</p>
</span></td></tr>
<tr><td><code>generate_subscripts(array: anyelement[], dim: <a href="int.html">int</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns a series comprising the given array’s subscripts.</p>
</span></td></tr>
<tr><td><code>generate_subscripts(array: anyelement[], dim: <a href="int.html">int</a>, reverse: <a href="bool.html">bool</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns a series comprising the given array’s subscripts.</p>
<p>When reverse is true, the series is returned in reverse order.</p>
</span></td></tr>
<tr><td><code>information_schema._pg_expandarray(input: anyelement[]) &rarr; anyelement</code></td><td><span class="funcdesc"><p>Returns the input array as a set of rows with an index</p>
</span></td></tr>
<tr><td><code>json_array_elements(input: jsonb) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Expands a JSON array to a set of JSON values.</p>
</span></td></tr>
<tr><td><code>json_array_elements_text(input: jsonb) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Expands a JSON array to a set of text values.</p>
</span></td></tr>
<tr><td><code>json_each(input: jsonb) &rarr; tuple{string AS key, jsonb AS value}</code></td><td><span class="funcdesc"><p>Expands the outermost JSON or JSONB object into a set of key/value pairs.</p>
</span></td></tr>
<tr><td><code>json_each_text(input: jsonb) &rarr; tuple{string AS key, string AS value}</code></td><td><span class="funcdesc"><p>Expands the outermost JSON or JSONB object into a set of key/value pairs. The returned values will be of type text.</p>
</span></td></tr>
<tr><td><code>json_object_keys(input: jsonb) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns sorted set of keys in the outermost JSON object.</p>
</span></td></tr>
<tr><td><code>jsonb_array_elements(input: jsonb) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Expands a JSON array to a set of JSON values.</p>
</span></td></tr>
<tr><td><code>jsonb_array_elements_text(input: jsonb) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Expands a JSON array to a set of text values.</p>
</span></td></tr>
<tr><td><code>jsonb_each(input: jsonb) &rarr; tuple{string AS key, jsonb AS value}</code></td><td><span class="funcdesc"><p>Expands the outermost JSON or JSONB object into a set of key/value pairs.</p>
</span></td></tr>
<tr><td><code>jsonb_each_text(input: jsonb) &rarr; tuple{string AS key, string AS value}</code></td><td><span class="funcdesc"><p>Expands the outermost JSON or JSONB object into a set of key/value pairs. The returned values will be of type text.</p>
</span></td></tr>
<tr><td><code>jsonb_object_keys(input: jsonb) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns sorted set of keys in the outermost JSON object.</p>
</span></td></tr>
<tr><td><code>pg_get_keywords() &rarr; tuple{string AS word, string AS catcode, string AS catdesc}</code></td><td><span class="funcdesc"><p>Produces a virtual table containing the keywords known to the SQL parser.</p>
</span></td></tr>
<tr><td><code>unnest(input: anyelement[]) &rarr; anyelement</code></td><td><span class="funcdesc"><p>Returns the input array as a set of rows</p>
</span></td></tr></tbody>
</table>

### String and byte functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>ascii(val: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Returns the character code of the first character in <code>val</code>. Despite the name, the function supports Unicode too.</p>
</span></td></tr>
<tr><td><code>bit_length(val: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of bits in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>bit_length(val: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of bits used to represent <code>val</code>.</p>
</span></td></tr>
<tr><td><code>btrim(input: <a href="string.html">string</a>, trim_chars: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Removes any characters included in <code>trim_chars</code> from the beginning or end of <code>input</code> (applies recursively).</p>
<p>For example, <code>btrim('doggie', 'eod')</code> returns <code>ggi</code>.</p>
</span></td></tr>
<tr><td><code>btrim(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Removes all spaces from the beginning and end of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>char_length(val: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of bytes in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>char_length(val: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of characters in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>character_length(val: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of bytes in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>character_length(val: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of characters in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>chr(val: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the character with the code given in <code>val</code>. Inverse function of <code>ascii()</code>.</p>
</span></td></tr>
<tr><td><code>concat(<a href="string.html">string</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Concatenates a comma-separated list of strings.</p>
</span></td></tr>
<tr><td><code>concat_ws(<a href="string.html">string</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Uses the first argument as a separator between the concatenation of the subsequent arguments.</p>
<p>For example <code>concat_ws('!','wow','great')</code> returns <code>wow!great</code>.</p>
</span></td></tr>
<tr><td><code>convert_from(str: <a href="bytes.html">bytes</a>, enc: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Decode the bytes in <code>str</code> into a string using encoding <code>enc</code>. Supports encodings ‘UTF8’ and ‘LATIN1’.</p>
</span></td></tr>
<tr><td><code>convert_to(str: <a href="string.html">string</a>, enc: <a href="string.html">string</a>) &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Encode the string <code>str</code> as a byte array using encoding <code>enc</code>. Supports encodings ‘UTF8’ and ‘LATIN1’.</p>
</span></td></tr>
<tr><td><code>decode(text: <a href="string.html">string</a>, format: <a href="string.html">string</a>) &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Decodes <code>data</code> using <code>format</code> (<code>hex</code> / <code>escape</code> / <code>base64</code>).</p>
</span></td></tr>
<tr><td><code>encode(data: <a href="bytes.html">bytes</a>, format: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Encodes <code>data</code> using <code>format</code> (<code>hex</code> / <code>escape</code> / <code>base64</code>).</p>
</span></td></tr>
<tr><td><code>from_ip(val: <a href="bytes.html">bytes</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts the byte string representation of an IP to its character string representation.</p>
</span></td></tr>
<tr><td><code>from_uuid(val: <a href="bytes.html">bytes</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts the byte string representation of a UUID to its character string representation.</p>
</span></td></tr>
<tr><td><code>initcap(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Capitalizes the first letter of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>left(input: <a href="bytes.html">bytes</a>, return_set: <a href="int.html">int</a>) &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Returns the first <code>return_set</code> bytes from <code>input</code>.</p>
</span></td></tr>
<tr><td><code>left(input: <a href="string.html">string</a>, return_set: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the first <code>return_set</code> characters from <code>input</code>.</p>
</span></td></tr>
<tr><td><code>length(val: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of bytes in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>length(val: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of characters in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>lower(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts all characters in <code>val</code> to their lower-case equivalents.</p>
</span></td></tr>
<tr><td><code>lpad(string: <a href="string.html">string</a>, length: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Pads <code>string</code> to <code>length</code> by adding ’ ’ to the left of <code>string</code>.If <code>string</code> is longer than <code>length</code> it is truncated.</p>
</span></td></tr>
<tr><td><code>lpad(string: <a href="string.html">string</a>, length: <a href="int.html">int</a>, fill: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Pads <code>string</code> by adding <code>fill</code> to the left of <code>string</code> to make it <code>length</code>. If <code>string</code> is longer than <code>length</code> it is truncated.</p>
</span></td></tr>
<tr><td><code>ltrim(input: <a href="string.html">string</a>, trim_chars: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Removes any characters included in <code>trim_chars</code> from the beginning (left-hand side) of <code>input</code> (applies recursively).</p>
<p>For example, <code>ltrim('doggie', 'od')</code> returns <code>ggie</code>.</p>
</span></td></tr>
<tr><td><code>ltrim(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Removes all spaces from the beginning (left-hand side) of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>md5(<a href="bytes.html">bytes</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the MD5 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>md5(<a href="string.html">string</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the MD5 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>octet_length(val: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of bytes in <code>val</code>.</p>
</span></td></tr>
<tr><td><code>octet_length(val: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the number of bytes used to represent <code>val</code>.</p>
</span></td></tr>
<tr><td><code>overlay(input: <a href="string.html">string</a>, overlay_val: <a href="string.html">string</a>, start_pos: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Replaces characters in <code>input</code> with <code>overlay_val</code> starting at <code>start_pos</code> (begins at 1).</p>
<p>For example, <code>overlay('doggie', 'CAT', 2)</code> returns <code>dCATie</code>.</p>
</span></td></tr>
<tr><td><code>overlay(input: <a href="string.html">string</a>, overlay_val: <a href="string.html">string</a>, start_pos: <a href="int.html">int</a>, end_pos: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Deletes the characters in <code>input</code> between <code>start_pos</code> and <code>end_pos</code> (count starts at 1), and then insert <code>overlay_val</code> at <code>start_pos</code>.</p>
</span></td></tr>
<tr><td><code>pg_collation_for(str: anyelement) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the collation of the argument</p>
</span></td></tr>
<tr><td><code>quote_ident(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Return <code>val</code> suitably quoted to serve as identifier in a SQL statement.</p>
</span></td></tr>
<tr><td><code>quote_literal(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Return <code>val</code> suitably quoted to serve as string literal in a SQL statement.</p>
</span></td></tr>
<tr><td><code>quote_literal(val: anyelement) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Coerce <code>val</code> to a string and then quote it as a literal.</p>
</span></td></tr>
<tr><td><code>quote_nullable(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Coerce <code>val</code> to a string and then quote it as a literal. If <code>val</code> is NULL, returns ‘NULL’.</p>
</span></td></tr>
<tr><td><code>quote_nullable(val: anyelement) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Coerce <code>val</code> to a string and then quote it as a literal. If <code>val</code> is NULL, returns ‘NULL’.</p>
</span></td></tr>
<tr><td><code>regexp_extract(input: <a href="string.html">string</a>, regex: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the first match for the Regular Expression <code>regex</code> in <code>input</code>.</p>
</span></td></tr>
<tr><td><code>regexp_replace(input: <a href="string.html">string</a>, regex: <a href="string.html">string</a>, replace: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Replaces matches for the Regular Expression <code>regex</code> in <code>input</code> with the Regular Expression <code>replace</code>.</p>
</span></td></tr>
<tr><td><code>regexp_replace(input: <a href="string.html">string</a>, regex: <a href="string.html">string</a>, replace: <a href="string.html">string</a>, flags: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Replaces matches for the regular expression <code>regex</code> in <code>input</code> with the regular expression <code>replace</code> using <code>flags</code>.</p>
<p>CockroachDB supports the following flags:</p>
<table>
<thead>
<tr>
<th>Flag</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><strong>c</strong></td>
<td>Case-sensitive matching</td>
</tr>
<tr>
<td><strong>g</strong></td>
<td>Global matching (match each substring instead of only the first)</td>
</tr>
<tr>
<td><strong>i</strong></td>
<td>Case-insensitive matching</td>
</tr>
<tr>
<td><strong>m</strong> or <strong>n</strong></td>
<td>Newline-sensitive (see below)</td>
</tr>
<tr>
<td><strong>p</strong></td>
<td>Partial newline-sensitive matching (see below)</td>
</tr>
<tr>
<td><strong>s</strong></td>
<td>Newline-insensitive (default)</td>
</tr>
<tr>
<td><strong>w</strong></td>
<td>Inverse partial newline-sensitive matching (see below)</td>
</tr>
</tbody>
</table>
<table>
<thead>
<tr>
<th>Mode</th>
<th><code>.</code> and <code>[^...]</code> match newlines</th>
<th><code>^</code> and <code>$</code> match line boundaries</th>
</tr>
</thead>
<tbody>
<tr>
<td>s</td>
<td>yes</td>
<td>no</td>
</tr>
<tr>
<td>w</td>
<td>yes</td>
<td>yes</td>
</tr>
<tr>
<td>p</td>
<td>no</td>
<td>no</td>
</tr>
<tr>
<td>m/n</td>
<td>no</td>
<td>yes</td>
</tr>
</tbody>
</table>
</span></td></tr>
<tr><td><code>repeat(input: <a href="string.html">string</a>, repeat_counter: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Concatenates <code>input</code> <code>repeat_counter</code> number of times.</p>
<p>For example, <code>repeat('dog', 2)</code> returns <code>dogdog</code>.</p>
</span></td></tr>
<tr><td><code>replace(input: <a href="string.html">string</a>, find: <a href="string.html">string</a>, replace: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Replaces all occurrences of <code>find</code> with <code>replace</code> in <code>input</code></p>
</span></td></tr>
<tr><td><code>reverse(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Reverses the order of the string’s characters.</p>
</span></td></tr>
<tr><td><code>right(input: <a href="bytes.html">bytes</a>, return_set: <a href="int.html">int</a>) &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Returns the last <code>return_set</code> bytes from <code>input</code>.</p>
</span></td></tr>
<tr><td><code>right(input: <a href="string.html">string</a>, return_set: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the last <code>return_set</code> characters from <code>input</code>.</p>
</span></td></tr>
<tr><td><code>rpad(string: <a href="string.html">string</a>, length: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Pads <code>string</code> to <code>length</code> by adding ’ ’ to the right of string. If <code>string</code> is longer than <code>length</code> it is truncated.</p>
</span></td></tr>
<tr><td><code>rpad(string: <a href="string.html">string</a>, length: <a href="int.html">int</a>, fill: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Pads <code>string</code> to <code>length</code> by adding <code>fill</code> to the right of <code>string</code>. If <code>string</code> is longer than <code>length</code> it is truncated.</p>
</span></td></tr>
<tr><td><code>rtrim(input: <a href="string.html">string</a>, trim_chars: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Removes any characters included in <code>trim_chars</code> from the end (right-hand side) of <code>input</code> (applies recursively).</p>
<p>For example, <code>rtrim('doggie', 'ei')</code> returns <code>dogg</code>.</p>
</span></td></tr>
<tr><td><code>rtrim(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Removes all spaces from the end (right-hand side) of <code>val</code>.</p>
</span></td></tr>
<tr><td><code>sha1(<a href="bytes.html">bytes</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the SHA1 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>sha1(<a href="string.html">string</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the SHA1 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>sha256(<a href="bytes.html">bytes</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the SHA256 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>sha256(<a href="string.html">string</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the SHA256 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>sha512(<a href="bytes.html">bytes</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the SHA512 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>sha512(<a href="string.html">string</a>...) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Calculates the SHA512 hash value of a set of values.</p>
</span></td></tr>
<tr><td><code>split_part(input: <a href="string.html">string</a>, delimiter: <a href="string.html">string</a>, return_index_pos: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Splits <code>input</code> on <code>delimiter</code> and return the value in the <code>return_index_pos</code>  position (starting at 1).</p>
<p>For example, <code>split_part('123.456.789.0','.',3)</code>returns <code>789</code>.</p>
</span></td></tr>
<tr><td><code>strpos(input: <a href="string.html">string</a>, find: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Calculates the position where the string <code>find</code> begins in <code>input</code>.</p>
<p>For example, <code>strpos('doggie', 'gie')</code> returns <code>4</code>.</p>
</span></td></tr>
<tr><td><code>substr(input: <a href="string.html">string</a>, regex: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> that matches the regular expression <code>regex</code>.</p>
</span></td></tr>
<tr><td><code>substr(input: <a href="string.html">string</a>, regex: <a href="string.html">string</a>, escape_char: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> that matches the regular expression <code>regex</code> using <code>escape_char</code> as your escape character instead of <code>\</code>.</p>
</span></td></tr>
<tr><td><code>substr(input: <a href="string.html">string</a>, start_pos: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> starting at <code>start_pos</code> (count starts at 1).</p>
</span></td></tr>
<tr><td><code>substr(input: <a href="string.html">string</a>, start_pos: <a href="int.html">int</a>, length: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> starting at <code>start_pos</code> (count starts at 1) and including up to <code>length</code> characters.</p>
</span></td></tr>
<tr><td><code>substring(input: <a href="string.html">string</a>, regex: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> that matches the regular expression <code>regex</code>.</p>
</span></td></tr>
<tr><td><code>substring(input: <a href="string.html">string</a>, regex: <a href="string.html">string</a>, escape_char: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> that matches the regular expression <code>regex</code> using <code>escape_char</code> as your escape character instead of <code>\</code>.</p>
</span></td></tr>
<tr><td><code>substring(input: <a href="string.html">string</a>, start_pos: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> starting at <code>start_pos</code> (count starts at 1).</p>
</span></td></tr>
<tr><td><code>substring(input: <a href="string.html">string</a>, start_pos: <a href="int.html">int</a>, length: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns a substring of <code>input</code> starting at <code>start_pos</code> (count starts at 1) and including up to <code>length</code> characters.</p>
</span></td></tr>
<tr><td><code>to_english(val: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>This function enunciates the value of its argument using English cardinals.</p>
</span></td></tr>
<tr><td><code>to_hex(val: <a href="bytes.html">bytes</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts <code>val</code> to its hexadecimal representation.</p>
</span></td></tr>
<tr><td><code>to_hex(val: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts <code>val</code> to its hexadecimal representation.</p>
</span></td></tr>
<tr><td><code>to_ip(val: <a href="string.html">string</a>) &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Converts the character string representation of an IP to its byte string representation.</p>
</span></td></tr>
<tr><td><code>to_uuid(val: <a href="string.html">string</a>) &rarr; <a href="bytes.html">bytes</a></code></td><td><span class="funcdesc"><p>Converts the character string representation of a UUID to its byte string representation.</p>
</span></td></tr>
<tr><td><code>translate(input: <a href="string.html">string</a>, find: <a href="string.html">string</a>, replace: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>In <code>input</code>, replaces the first character from <code>find</code> with the first character in <code>replace</code>; repeat for each character in <code>find</code>.</p>
<p>For example, <code>translate('doggie', 'dog', '123');</code> returns <code>1233ie</code>.</p>
</span></td></tr>
<tr><td><code>upper(val: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Converts all characters in <code>val</code> to their to their upper-case equivalents.</p>
</span></td></tr></tbody>
</table>

### System info functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>cluster_logical_timestamp() &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>Returns the logical time of the current transaction.</p>
<p>This function is reserved for testing purposes by CockroachDB
developers and its definition may change without prior notice.</p>
<p>Note that uses of this function disable server-side optimizations and
may increase either contention or retry errors, or both.</p>
</span></td></tr>
<tr><td><code>crdb_internal.check_consistency(stats_only: <a href="bool.html">bool</a>, start_key: <a href="bytes.html">bytes</a>, end_key: <a href="bytes.html">bytes</a>) &rarr; tuple{int AS range_id, bytes AS start_key, string AS start_key_pretty, string AS status, string AS detail}</code></td><td><span class="funcdesc"><p>Runs a consistency check on ranges touching the specified key range. an empty start or end key is treated as the minimum and maximum possible, respectively. stats_only should only be set to false when targeting a small number of ranges to avoid overloading the cluster. Each returned row contains the range ID, the status (a roachpb.CheckConsistencyResponse_Status), and verbose detail.</p>
<p>Example usage:
SELECT * FROM crdb_internal.check_consistency(true, ‘\x02’, ‘\x04’)</p>
</span></td></tr>
<tr><td><code>crdb_internal.cluster_id() &rarr; <a href="uuid.html">uuid</a></code></td><td><span class="funcdesc"><p>Returns the cluster ID.</p>
</span></td></tr>
<tr><td><code>crdb_internal.force_assertion_error(msg: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.force_error(errorCode: <a href="string.html">string</a>, msg: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.force_log_fatal(msg: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.force_panic(msg: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.force_retry(val: <a href="interval.html">interval</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.json_num_index_entries(val: jsonb) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.lease_holder(key: <a href="bytes.html">bytes</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>This function is used to fetch the leaseholder corresponding to a request key</p>
</span></td></tr>
<tr><td><code>crdb_internal.locality_value(key: <a href="string.html">string</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the value of the specified locality key.</p>
</span></td></tr>
<tr><td><code>crdb_internal.no_constant_folding(input: anyelement) &rarr; anyelement</code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.node_executable_version() &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the version of CockroachDB this node is running.</p>
</span></td></tr>
<tr><td><code>crdb_internal.pretty_key(raw_key: <a href="bytes.html">bytes</a>, skip_fields: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>This function is used only by CockroachDB’s developers for testing purposes.</p>
</span></td></tr>
<tr><td><code>crdb_internal.round_decimal_values(val: <a href="decimal.html">decimal</a>, scale: <a href="int.html">int</a>) &rarr; <a href="decimal.html">decimal</a></code></td><td><span class="funcdesc"><p>This function is used internally to round decimal values during mutations.</p>
</span></td></tr>
<tr><td><code>crdb_internal.round_decimal_values(val: <a href="decimal.html">decimal</a>[], scale: <a href="int.html">int</a>) &rarr; <a href="decimal.html">decimal</a>[]</code></td><td><span class="funcdesc"><p>This function is used internally to round decimal array values during mutations.</p>
</span></td></tr>
<tr><td><code>crdb_internal.set_vmodule(vmodule_string: <a href="string.html">string</a>) &rarr; <a href="int.html">int</a></code></td><td><span class="funcdesc"><p>Set the equivalent of the <code>--vmodule</code> flag on the gateway node processing this request; it affords control over the logging verbosity of different files. Example syntax: <code>crdb_internal.set_vmodule('recordio=2,file=1,gfs*=3')</code>. Reset with: <code>crdb_internal.set_vmodule('')</code>. Raising the verbosity can severely affect performance.</p>
</span></td></tr>
<tr><td><code>current_database() &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the current database.</p>
</span></td></tr>
<tr><td><code>current_schema() &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the current schema.</p>
</span></td></tr>
<tr><td><code>current_schemas(include_pg_catalog: <a href="bool.html">bool</a>) &rarr; <a href="string.html">string</a>[]</code></td><td><span class="funcdesc"><p>Returns the valid schemas in the search path.</p>
</span></td></tr>
<tr><td><code>current_user() &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the current user. This function is provided for compatibility with PostgreSQL.</p>
</span></td></tr>
<tr><td><code>version() &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the node’s version of CockroachDB.</p>
</span></td></tr></tbody>
</table>

### TUPLE functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>row_to_json(row: tuple) &rarr; jsonb</code></td><td><span class="funcdesc"><p>Returns the row as a JSON object.</p>
</span></td></tr></tbody>
</table>

### Compatibility functions

<table>
<thead><tr><th>Function &rarr; Returns</th><th>Description</th></tr></thead>
<tbody>
<tr><td><code>format_type(type_oid: oid, typemod: <a href="int.html">int</a>) &rarr; <a href="string.html">string</a></code></td><td><span class="funcdesc"><p>Returns the SQL name of a data type that is identified by its type OID and possibly a type modifier. Currently, the type modifier is ignored.</p>
</span></td></tr>
<tr><td><code>has_any_column_privilege(table: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for any column of table.</p>
</span></td></tr>
<tr><td><code>has_any_column_privilege(table: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for any column of table.</p>
</span></td></tr>
<tr><td><code>has_any_column_privilege(user: <a href="string.html">string</a>, table: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for any column of table.</p>
</span></td></tr>
<tr><td><code>has_any_column_privilege(user: <a href="string.html">string</a>, table: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for any column of table.</p>
</span></td></tr>
<tr><td><code>has_any_column_privilege(user: oid, table: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for any column of table.</p>
</span></td></tr>
<tr><td><code>has_any_column_privilege(user: oid, table: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for any column of table.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(table: <a href="string.html">string</a>, column: <a href="int.html">int</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(table: <a href="string.html">string</a>, column: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(table: oid, column: <a href="int.html">int</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(table: oid, column: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: <a href="string.html">string</a>, table: <a href="string.html">string</a>, column: <a href="int.html">int</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: <a href="string.html">string</a>, table: <a href="string.html">string</a>, column: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: <a href="string.html">string</a>, table: oid, column: <a href="int.html">int</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: <a href="string.html">string</a>, table: oid, column: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: oid, table: <a href="string.html">string</a>, column: <a href="int.html">int</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: oid, table: <a href="string.html">string</a>, column: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: oid, table: oid, column: <a href="int.html">int</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_column_privilege(user: oid, table: oid, column: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for column.</p>
</span></td></tr>
<tr><td><code>has_database_privilege(database: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for database.</p>
</span></td></tr>
<tr><td><code>has_database_privilege(database: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for database.</p>
</span></td></tr>
<tr><td><code>has_database_privilege(user: <a href="string.html">string</a>, database: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for database.</p>
</span></td></tr>
<tr><td><code>has_database_privilege(user: <a href="string.html">string</a>, database: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for database.</p>
</span></td></tr>
<tr><td><code>has_database_privilege(user: oid, database: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for database.</p>
</span></td></tr>
<tr><td><code>has_database_privilege(user: oid, database: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for database.</p>
</span></td></tr>
<tr><td><code>has_foreign_data_wrapper_privilege(fdw: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for foreign-data wrapper.</p>
</span></td></tr>
<tr><td><code>has_foreign_data_wrapper_privilege(fdw: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for foreign-data wrapper.</p>
</span></td></tr>
<tr><td><code>has_foreign_data_wrapper_privilege(user: <a href="string.html">string</a>, fdw: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign-data wrapper.</p>
</span></td></tr>
<tr><td><code>has_foreign_data_wrapper_privilege(user: <a href="string.html">string</a>, fdw: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign-data wrapper.</p>
</span></td></tr>
<tr><td><code>has_foreign_data_wrapper_privilege(user: oid, fdw: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign-data wrapper.</p>
</span></td></tr>
<tr><td><code>has_foreign_data_wrapper_privilege(user: oid, fdw: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign-data wrapper.</p>
</span></td></tr>
<tr><td><code>has_function_privilege(function: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for function.</p>
</span></td></tr>
<tr><td><code>has_function_privilege(function: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for function.</p>
</span></td></tr>
<tr><td><code>has_function_privilege(user: <a href="string.html">string</a>, function: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for function.</p>
</span></td></tr>
<tr><td><code>has_function_privilege(user: <a href="string.html">string</a>, function: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for function.</p>
</span></td></tr>
<tr><td><code>has_function_privilege(user: oid, function: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for function.</p>
</span></td></tr>
<tr><td><code>has_function_privilege(user: oid, function: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for function.</p>
</span></td></tr>
<tr><td><code>has_language_privilege(language: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for language.</p>
</span></td></tr>
<tr><td><code>has_language_privilege(language: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for language.</p>
</span></td></tr>
<tr><td><code>has_language_privilege(user: <a href="string.html">string</a>, language: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for language.</p>
</span></td></tr>
<tr><td><code>has_language_privilege(user: <a href="string.html">string</a>, language: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for language.</p>
</span></td></tr>
<tr><td><code>has_language_privilege(user: oid, language: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for language.</p>
</span></td></tr>
<tr><td><code>has_language_privilege(user: oid, language: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for language.</p>
</span></td></tr>
<tr><td><code>has_schema_privilege(schema: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for schema.</p>
</span></td></tr>
<tr><td><code>has_schema_privilege(schema: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for schema.</p>
</span></td></tr>
<tr><td><code>has_schema_privilege(user: <a href="string.html">string</a>, schema: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for schema.</p>
</span></td></tr>
<tr><td><code>has_schema_privilege(user: <a href="string.html">string</a>, schema: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for schema.</p>
</span></td></tr>
<tr><td><code>has_schema_privilege(user: oid, schema: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for schema.</p>
</span></td></tr>
<tr><td><code>has_schema_privilege(user: oid, schema: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for schema.</p>
</span></td></tr>
<tr><td><code>has_sequence_privilege(sequence: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for sequence.</p>
</span></td></tr>
<tr><td><code>has_sequence_privilege(sequence: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for sequence.</p>
</span></td></tr>
<tr><td><code>has_sequence_privilege(user: <a href="string.html">string</a>, sequence: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for sequence.</p>
</span></td></tr>
<tr><td><code>has_sequence_privilege(user: <a href="string.html">string</a>, sequence: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for sequence.</p>
</span></td></tr>
<tr><td><code>has_sequence_privilege(user: oid, sequence: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for sequence.</p>
</span></td></tr>
<tr><td><code>has_sequence_privilege(user: oid, sequence: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for sequence.</p>
</span></td></tr>
<tr><td><code>has_server_privilege(server: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for foreign server.</p>
</span></td></tr>
<tr><td><code>has_server_privilege(server: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for foreign server.</p>
</span></td></tr>
<tr><td><code>has_server_privilege(user: <a href="string.html">string</a>, server: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign server.</p>
</span></td></tr>
<tr><td><code>has_server_privilege(user: <a href="string.html">string</a>, server: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign server.</p>
</span></td></tr>
<tr><td><code>has_server_privilege(user: oid, server: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign server.</p>
</span></td></tr>
<tr><td><code>has_server_privilege(user: oid, server: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for foreign server.</p>
</span></td></tr>
<tr><td><code>has_table_privilege(table: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for table.</p>
</span></td></tr>
<tr><td><code>has_table_privilege(table: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for table.</p>
</span></td></tr>
<tr><td><code>has_table_privilege(user: <a href="string.html">string</a>, table: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for table.</p>
</span></td></tr>
<tr><td><code>has_table_privilege(user: <a href="string.html">string</a>, table: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for table.</p>
</span></td></tr>
<tr><td><code>has_table_privilege(user: oid, table: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for table.</p>
</span></td></tr>
<tr><td><code>has_table_privilege(user: oid, table: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for table.</p>
</span></td></tr>
<tr><td><code>has_tablespace_privilege(tablespace: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for tablespace.</p>
</span></td></tr>
<tr><td><code>has_tablespace_privilege(tablespace: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for tablespace.</p>
</span></td></tr>
<tr><td><code>has_tablespace_privilege(user: <a href="string.html">string</a>, tablespace: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for tablespace.</p>
</span></td></tr>
<tr><td><code>has_tablespace_privilege(user: <a href="string.html">string</a>, tablespace: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for tablespace.</p>
</span></td></tr>
<tr><td><code>has_tablespace_privilege(user: oid, tablespace: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for tablespace.</p>
</span></td></tr>
<tr><td><code>has_tablespace_privilege(user: oid, tablespace: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for tablespace.</p>
</span></td></tr>
<tr><td><code>has_type_privilege(type: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for type.</p>
</span></td></tr>
<tr><td><code>has_type_privilege(type: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the current user has privileges for type.</p>
</span></td></tr>
<tr><td><code>has_type_privilege(user: <a href="string.html">string</a>, type: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for type.</p>
</span></td></tr>
<tr><td><code>has_type_privilege(user: <a href="string.html">string</a>, type: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for type.</p>
</span></td></tr>
<tr><td><code>has_type_privilege(user: oid, type: <a href="string.html">string</a>, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for type.</p>
</span></td></tr>
<tr><td><code>has_type_privilege(user: oid, type: oid, privilege: <a href="string.html">string</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>Returns whether or not the user has privileges for type.</p>
</span></td></tr>
<tr><td><code>oid(int: <a href="int.html">int</a>) &rarr; oid</code></td><td><span class="funcdesc"><p>Converts an integer to an OID.</p>
</span></td></tr>
<tr><td><code>pg_sleep(seconds: <a href="float.html">float</a>) &rarr; <a href="bool.html">bool</a></code></td><td><span class="funcdesc"><p>pg_sleep makes the current session’s process sleep until seconds seconds have elapsed. seconds is a value of type double precision, so fractional-second delays can be specified.</p>
</span></td></tr></tbody>
</table>

