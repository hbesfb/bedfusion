# Padding

BedFusion offers the option to pad and merge at the same time, or only pad your files without merging. To make these examples easier to follow we will work with simple bed files only containing chromosomes, start and stop positions, we will also focus on padding without merging by using the `--no-merge` flag. But ``--padding` can be of course be combined with other options like `--strand-col` and `--feat-col`.

When merging it is important to note that the regions as padded before they are merged. So if `--padding` is used together with `--overlap` padding is added first, and the overlap will between the merged regions will then be considered. 

Example bed files 

- `examples/padding-test.bed`:

```bed
1	1	4
1	5	8
10	5	8
1	20	30
```
- `examples/padding-test.bed`:

```bed
1	1	4
1	5	8
2	5	8
1	20	30
```

Example FASTA index file `examples/test.fasta.fai`:

```text
1	249250621	52	60	61
10	135534747	1708379889	60	61
```

## Padding Type Error (Default)

Error is the safest padding option and will only pad regions present in the provided FASTA index file, and will immediately fail if it encounters a chromosome that is not present in the FASTA index file.

Example of padding when all chromosomes are present in the FASTA index file:

```shell
> bedfusion examples/padding-test.bed --fasta-idx=examples/test.fasta.fai --padding=10 --no-merge
1       1       14
1       1       18
1       10      40
10      1       18
```

```shell
> bedfusion examples/padding-test.bed --fasta-idx=examples/test.fasta.fai --padding=10 --padding-type="err" --no-merge
1       1       14
1       1       18
1       10      40
10      1       18
```


Example with merging where one or more chromosomes are missing in the FASTA index file:

```shell
> bedfusion examples/padding-test2.bed --fasta-idx=examples/test.fasta.fai --padding=10 --no-merge
bedfusion: error: while padding: chromosome 2 is not in fasta index file examples/test.fasta.fai
exit status 1
```

## Padding Type Warning

Warning is the next safest padding option and will only pad regions present in the provided FASTA index file. It will NOT add padding for chromosomes not present in the FASTA index file, but warn about the user about them.

Example of padding when all chromosomes are present in the FASTA index file:

```shell
> bedfusion examples/padding-test.bed --fasta-idx=examples/test.fasta.fai --padding=10 --padding-type="warn" --no-merge
1       1       14
1       1       18
1       10      40
10      1       18
```

Example with merging where one or more chromosomes are missing in the FASTA index file:

```shell
> bedfusion examples/padding-test2.bed --fasta-idx=examples/test.fasta.fai --padding=10 --padding-type="warn" --no-merge
warning: chromosomes [2] not in fasta index file examples/test.fasta.fai, no padding was added to regions on these chromosomes
1       1       14
1       1       18
1       10      40
2       5       8
```

## Padding Type Force

Force is the most unsafe padding option and will pad all regions regardless if the chromosome is present in the FASTA index file of not. For this padding type providing a FASTA index file is optional, but if a FASTA index file is provided BedFusion will warn about chromosomes missing in the FASTA index file. Please note that using this padding option might result in your bed files not being compatible with other tools. 

Example of padding when all chromosomes are present in the FASTA index file:

```shell
> bedfusion examples/padding-test.bed --fasta-idx=examples/test.fasta.fai --padding=10 --padding-type="force" --no-merge
1       1       14
1       1       18
1       10      40
10      1       18
```

Example with merging where one or more chromosomes are missing in the FASTA index file:

```shell
> bedfusion examples/padding-test2.bed --fasta-idx=examples/test.fasta.fai --padding=10 --padding-type="force" --no-merge
warning: chromosomes [2] not in fasta index file examples/test.fasta.fai, regions on these chromosomes were still padded
1       1       14
1       1       18
1       10      40
2       1       18

```

Example when not using a FASTA index file:

```shell
> bedfusion examples/padding-test2.bed --padding=10 --padding-type="force" --no-merge
1       1       14
1       1       18
1       10      40
2       1       18
```

## TODO: Combining padding and merging at the same time
