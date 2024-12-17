# Merging

BedFusion merges and sorts by default. To highlight the differences between the different merging methods, only the default [lexicographic sorting](./sorting.md#lexicographic-sorting) will be used in these examples.

Example bed file `examples/merge-test.bed`:
```bed
1       1       4       1       A
1       5       8       1       A
1       6       8       1       A
1       5       8       -1      A
2       5       8       1       A
1       5       8       1       B
1       20      30      1       A
```

## Default merging

When merging by default all touching and overlapping regions within the same chromosome will be merged. Unique values in optional columns will be concatenated and comma-separated if merged.

Example:

``` text
> bedfusion examples/merge-test.bed 
1       1       8       1,-1    A,B
1       20      30      1       A
2       5       8       1       A
```

## Merging with strand column set

When `--strand-col` is set regions on different strands and chromosomes will not be merged.

Example:

``` shell
> bedfusion examples/merge-test.bed --strand-col=4
1       1       8       1       A,B
1       5       8       -1      A
1       20      30      1       A
2       5       8       1       A
```

## Merging with feature column set

When `--feat-col` is set regions on different features (here we use the gene column as feature, but any other optional column can be chosen) and chromosomes will not be merged.

Example:

``` shell
> bedfusion examples/merge-test.bed --feat-col=5
1       1       8       1,-1    A
1       5       8       1       B
1       20      30      1       A
2       5       8       1       A
```

## Merging with both strand and feature columns set

When both `--feat -col` and `--feat-col` is set regions on different features, strand and chromosomes will not be merged.

Example:

``` shell
> bedfusion examples/merge-test.bed --strand-col=4 --feat-col=5
1       1       8       1       A
1       5       8       -1      A
1       5       8       1       B
1       20      30      1       A
2       5       8       1       A
```

## Using overlap

BedFusion merges touching regions by default (`--overlap=0`), but one can choose a custom overlap for regions one wants to be merged.

For example if one would only want overlapping, but not touching regions to merge one can set `--overlap=-1`:

``` shell 
> bedfusion examples/merge-test.bed --overlap=-1
1       1       4       1       A
1       5       8       1,-1    A,B
1       20      30      1       A
2       5       8       1       A
```

If one on the other hand would like regions further apart to be merged one can set the overlap to a higher number. For example, by setting `--overlap=11` we get this result:

``` shell 
> bedfusion examples/merge-test.bed --overlap=11
1       1       30      1,-1    A,B
2       5       8       1       A
```

Used together with `--strand-col` and `--feat-col`:

``` shell 
> bedfusion examples/merge-test.bed --strand-col=4 --feat-col=5 --overlap=11
1       1       30      1       A
1       5       8       -1      A
1       5       8       1       B
2       5       8       1       A
```

## No Merge

If one would prefer not to merge the `--no-merge` flag can be used.

Example:

``` shell
> bedfusion examples/merge-test.bed --no-merge
1       1       4       1       A
1       5       8       1       A
1       5       8       -1      A
1       5       8       1       B
1       6       8       1       A
1       20      30      1       A
2       5       8       1       A
```
