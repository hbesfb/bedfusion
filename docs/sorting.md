# Sorting

Note that bed fusion merges by default. So for the sorting examples we will not merge the bed regions, by using the flag `--no-merge`, to highlight the differences between the sorting methods.

Sorting is always done after merging, so the resulting output will be sorted in the preferred format regardless if the bed file is merged or not.

Example bed file `examples/sort-test.bed`:

``` text
2	12	13	1	C
Y	10	11	1	A
1	8	9	-1	B
10	12	13	1	D
GL000209.1	10	11	1	A
1	10	11	-1	A
1	12	13	1	A
X	10	11	1	A
1	10	11	1	A
1	10	11	-1	B
MT	10	11	1	A
```

## Lexicographic Sorting (default)

Lexicographic sorting is the default sorting method.

Example:

``` shell
> bedfusion examples/sort-test.bed --no-merge
1       8       9       -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       12      13      1       A
10      12      13      1       D
2       12      13      1       C
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```

``` shell
> bedfusion examples/sort-test.bed --no-merge -s lex
1       8       9       -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       12      13      1       A
10      12      13      1       D
2       12      13      1       C
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```

To also sort using strand and/or gene information we can set the flags `--strand-col` and/or `--feat-col`. If `--feat-col` is set it will also be sorted lexicographically In this example we will use gene as the `--feat-col`, but any optional column can be chosen:

``` shell
> bedfusion examples/sort-test.bed --no-merge --strand-col=4 --feat-col=5
1       8       9       -1      B
1       10      11      -1      A
1       10      11      -1      B
1       10      11      1       A
1       12      13      1       A
10      12      13      1       D
2       12      13      1       C
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```

## Natural Sorting

Example: 

``` shell 
> bedfusion examples/sort-test.bed --no-merge -s nat
1       8       9       -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       12      13      1       A
2       12      13      1       C
10      12      13      1       D
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```

Using the `--strand-col` and/or `--feat-col` options these column will also be sorted. If `--feat-col` is set the feature column (in this example gene) will be sorted using natural sorting:

``` shell 
> bedfusion examples/sort-test.bed --no-merge --strand-col=4 --feat-col=5 -s nat
1       8       9       -1      B
1       10      11      -1      A
1       10      11      -1      B
1       10      11      1       A
1       12      13      1       A
2       12      13      1       C
10      12      13      1       D
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```

## Custom Chromosome Sorting

Custom chromosome sorting lets you sort the chromosomes in any order that you would prefer.

By default it will use human chromosome sorting (1-21, X, Y, MT), for example:

``` shell
> bedfusion examples/sort-test.bed --no-merge -s ccs
1       8       9       -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       12      13      1       A
2       12      13      1       C
10      12      13      1       D
X       10      11      1       A
Y       10      11      1       A
MT      10      11      1       A
GL000209.1      10      11      1       A
```

Using the `--strand-col` and/or `--feat-col` options these column will also be sorted. If `--feat-col` is set the feature column (in this example gene) will be sorted using natural sorting:

``` shell 
> bedfusion examples/sort-test.bed --no-merge --strand-col=4 --feat-col=5 -s ccs
1       8       9       -1      B
1       10      11      -1      A
1       10      11      -1      B
1       10      11      1       A
1       12      13      1       A
2       12      13      1       C
10      12      13      1       D
X       10      11      1       A
Y       10      11      1       A
MT      10      11      1       A
GL000209.1      10      11      1       A
```

### Custom Chromosome Order

Adding the option custom chromosome sorting, one can sort the chromosomes in any order that one would want with the flag `--chr-order`. Chromosomes not on the list will be sorted naturally after the ones in the list.

``` shell
> bedfusion examples/sort-test.bed --no-merge --strand-col=4 --feat-col=5 -s ccs --chr-order=X,Y,10
X       10      11      1       A
Y       10      11      1       A
10      12      13      1       D
1       8       9       -1      B
1       10      11      -1      A
1       10      11      -1      B
1       10      11      1       A
1       12      13      1       A
2       12      13      1       C
GL000209.1      10      11      1       A
MT      10      11      1       A
```

Please note that the provided chromosome order is case insensitive.

## FASTA Index File Sorting 

FASTA index file sorting lets you sort the chromosomes according to the two first columns of a fasta-index file. Chromosomes not in the file will be sorted naturally after the ones in the file.

Example FASTA index file `examples/test.fasta.fai`: 

``` text
1	249250621	52	60	61
2	243199373	253404903	60	61
10	135534747	1708379889	60	61
```

Example: 

``` shell
> bedfusion examples/sort-test.bed --no-merge -s fidx --fasta-idx=examples/test.fasta.fai 
1       8       9       -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       12      13      1       A
10      12      13      1       D
2       12      13      1       C
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```

Using the `--strand-col` and/or `--feat-col` options these column will also be sorted. If `--feat-col` is set the feature column (in this example gene) will be sorted using natural sorting:

``` shell 
> bedfusion examples/sort-test.bed --no-merge -s fidx --fasta-idx=examples/test.fasta.fai --strand-col=4 --feat-col=5
1       8       9       -1      B
1       10      11      -1      A
1       10      11      -1      B
1       10      11      1       A
1       12      13      1       A
10      12      13      1       D
2       12      13      1       C
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```

## Deduplication 

When choosing not to merge the bed regions (by using the flag `--no-merge`) one might still want to remove duplicates.

Lets pretend that we copy the `examples/sort-test.bed` file and call this copy `examples/sort-test-copy.bed`, and provide both as input to BedFusion:

Without the flag `--deduplicate/-d`:

``` shell
> bedfusion examples/sort-test.bed examples/sort-test-copy.bed --no-merge
1       8       9       -1      B
1       8       9       -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       12      13      1       A
1       12      13      1       A
10      12      13      1       D
10      12      13      1       D
2       12      13      1       C
2       12      13      1       C
GL000209.1      10      11      1       A
GL000209.1      10      11      1       A
MT      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
X       10      11      1       A
Y       10      11      1       A
Y       10      11      1       A
```

With the flag `--deduplicate/-d`:

``` shell
> bedfusion examples/sort-test.bed sort-test-copy.bed --no-merge -d 
1       8       9       -1      B
1       10      11      -1      A
1       10      11      1       A
1       10      11      -1      B
1       12      13      1       A
10      12      13      1       D
2       12      13      1       C
GL000209.1      10      11      1       A
MT      10      11      1       A
X       10      11      1       A
Y       10      11      1       A
```
