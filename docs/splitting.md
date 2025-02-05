# Splitting

The `--split-size` flag splits regions into new regions of a chosen bp size after [padding](./padding.md) and [merging](./merging.md).

Example bed file `examples/split-test.bed`:

``` bed
1	25	100	1	A
2	1	40	1	B
1	1	75	1	A
```

Example when merging:

``` shell
> bedfusion examples/split-test.bed --split-size=24
1       1       25      1       A
1       26      50      1       A
1       51      75      1       A
1       76      100     1       A
2       1       25      1       B
2       26      40      1       B
```

Example when not merging:

``` shell
> bedfusion examples/split-test.bed --split-size=24 --no-merge
1       1       25      1       A
1       25      49      1       A
1       26      50      1       A
1       50      74      1       A
1       51      75      1       A
1       75      99      1       A
1       100     100     1       A
2       1       25      1       B
2       26      40      1       B
```

Note the overlapping regions when `--split-size` is used without merging. 
