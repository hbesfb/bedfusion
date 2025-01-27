# Padding

BedFusion offers the option to pad and merge at the same time, or only pad your files without merging. To make these examples easier to follow we will work with simple bed files only containing chromosomes, start and stop positions, we will also focus on padding without merging by using the `--no-merge` flag. But `--padding` can be of course be combined with other options like `--strand-col` and `--feat-col`.

When merging it is important to note that the regions as padded before they are merged. So if `--padding` is used together with `--overlap`, padding is added first, and then the overlap will be considered when merging the regions.

Example bed files 

- `examples/padding-test.bed`:

``` bed
1	1	4
1	5	9
10	5	8
1	20	30
```

- `examples/padding-test2.bed`:

``` bed
1	1	4
1	5	8
2	5	9
1	20	30
```

Example FASTA index file `examples/test.fasta.fai`:

``` txt
1	249250621	52	60	61
10	135534747	1708379889	60	61
```

## Padding Type Error (Default)

Error is the safest padding option and will only pad regions present in the provided FASTA index file, and will immediately fail if it encounters a chromosome that is not present in the FASTA index file.

Example of padding when all chromosomes are present in the FASTA index file:

``` shell
> bedfusion examples/padding-test.bed --no-merge --fasta-idx=examples/test.fasta.fai --padding=10
1       1       14
1       1       19
1       10      40
10      1       18
```

``` shell
> bedfusion examples/padding-test.bed --no-merge --fasta-idx=examples/test.fasta.fai --padding-type=err --padding=10
1       1       14
1       1       19
1       10      40
10      1       18
```

Example where one or more chromosomes are missing in the FASTA index file:

``` shell
> bedfusion examples/padding-test2.bed --no-merge --fasta-idx=examples/test.fasta.fai --padding=10
bedfusion: error: while padding: chromosome 2 is not in fasta index file examples/test.fasta.fai
```

## Padding Type Warning

Warning is the next safest padding option and will only pad regions present in the provided FASTA index file. It will NOT add padding for chromosomes not present in the FASTA index file, but warn about the user about them.

Example of padding when all chromosomes are present in the FASTA index file:

``` shell
> bedfusion examples/padding-test.bed --no-merge --fasta-idx=examples/test.fasta.fai --padding-type=warn --padding=10
1       1       14
1       1       19
1       10      40
10      1       18
```

Example with merging where one or more chromosomes are missing in the FASTA index file:

``` shell
> bedfusion examples/padding-test2.bed --no-merge --fasta-idx=examples/test.fasta.fai --padding-type=warn --padding=10
warning: chromosomes [2] not in fasta index file examples/test.fasta.fai, no padding was added to regions on these chromosomes
1       1       14
1       1       18
1       10      40
2       5       9
```

Note that the region on chromosome 2 is not padded, as it is missing from the FASTA index file.

## Padding Type Force

Force is the most unsafe padding option and will pad all regions regardless if the chromosome is present in the FASTA index file of not. For this padding type providing a FASTA index file is optional, but if a FASTA index file is provided BedFusion will warn about chromosomes missing in the FASTA index file. Please note that using this padding option might result in your bed files not being compatible with other tools. 

Example of padding when all chromosomes are present in the FASTA index file:

``` shell
> bedfusion examples/padding-test.bed --no-merge --fasta-idx=examples/test.fasta.fai --padding-type=force --padding=10
1       1       14
1       1       19
1       10      40
10      1       18
```

Example with merging where one or more chromosomes are missing in the FASTA index file:

``` shell
> bedfusion examples/padding-test2.bed --no-merge --fasta-idx=examples/test.fasta.fai --padding-type=force --padding=10
warning: chromosomes [2] not in fasta index file examples/test.fasta.fai, regions on these chromosomes were still padded
1       1       14
1       1       18
1       10      40
2       1       19
```

Note that the region on chromosome 2 is still padded even if it is missing from the FASTA index file.

Example when not using a FASTA index file:

``` shell
> bedfusion examples/padding-test2.bed --no-merge --padding-type=force --padding=10
1       1       14
1       1       18
1       10      40
2       1       19
```

## Combined use of `--overlap` and `--padding` when merging bed files

As mentioned above `--padding` and `--overlap` can be used together when merging. If so the padding is added first and then the overlap is considered after.

For example, to get the regions in `examples/padding-test.bed` to be merged together because regions are touching we ned to pad with 6 bp:

``` shell
> bedfusion examples/padding-test.bed --fasta-idx=examples/test.fasta.fai --padding=5 
1       1       35
10      1       13
```

However if we don't touching regions to be merged, only overlapping ones we can adjust this with `--overlap=-1`:

``` shell
> bedfusion examples/padding-test.bed --fasta-idx=examples/test.fasta.fai --padding=5 --overlap=-1
1       1       14
1       15      35
10      1       13
```



