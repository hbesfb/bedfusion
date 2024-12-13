# BedFusion

Another tool for sorting and merging bed files

Usage: `bedfusion <inputs> ... [flags]`

BedFusion follows the bed file standard outlined in: [Niu J., Denisko D. & Hoffman M. M. (2022): *The Browser Extensible Data (BED)* format](https://github.com/samtools/hts-specs/blob/94500cf76f049e898dec7af23097d877fde5894e/BEDv1.pdf)

## Quick-Start Guide

BedFusion will both sort (lexicographically) and merge regions by default. 

Example bed file `merge-test.bed`:

```bed
1	1	4	1	X
1	5	8	1	X
1	6	8	1	X
1	5	8	-1	X
2	5	8	1	X
1	5	8	1	Y
1	20	30	1	X
```

``` shell
> bedfusion merge-test.bed
1       1       8       1,-1    X,Y
1       20      30      1       X
2       5       8       1       X
```

Contrary to [bedtools merge](https://bedtools.readthedocs.io/en/latest/content/tools/merge.html), BedFusion merges touching regions (like the two first lines in the example bed file). If you prefer to only merge overlapping, and not touching, regions you can use the flag `--overlap -1`:

```shell

> bedfusion merge-test.bed --overlap=-1
1       1       4       1       X
1       5       8       1,-1    X,Y
1       20      30      1       X
2       5       8       1       X
```

## Examples

- [sorting](./docs/sorting.md)
- [merging](./docs/merging.md)

TODO:
- track file support
- using config files

## Flags and arguments 

BedFusion supports the use of selecting your options in three separate ways: As flags, in a configuration file or as environmental variables. If a combination of the three is used the reading priority order is as follows: 

1. flags 
2. configuration file 
3. environmental variables


| Arguments      |                                                                                                  |
|----------------|--------------------------------------------------------------------------------------------------|
| `<inputs> ...` | Bed file path(s). If more than one is provided the files will be joined as if they were one file |


| Flags (with format and defaults)    | Environmental variables | Description                                                                                                                                                                                                                                                       |
|-------------------------------------|-------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `-h`<br>`--help`                    |                         | Show context-sensitive help                                                                                                                                                                                                                                       |
| `-c`<br>`--config-file=CONFIG-FLAG` | `CONFIG_FILE`           | The path to configuration file (must be in key-value yaml format)                                                                                                                                                                                                 |
| `-o`<br>`--output=STRING`           | `OUTPUT_FILE`           | Path to the output file. If unset the output will be written to stdout                                                                                                                                                                                            |
|                                     |                         |                                                                                                                                                                                                                                                                   |
| **input**                           |                         |                                                                                                                                                                                                                                                                   |
| `--strand-col=INT`                  | `STRAND_COL`            | The column containing the strand information (1-based column index). If this option is set regions on the same strand will not be merged                                                                                                                          |
| `--feat-col=INT`                    | `FEAT_COL`              | The column containing the feature (e.g. gene id, transcript id etc.) information (1-based column index). If this option is set regions on the same feature will not be merged                                                                                     |
|                                     |                         |                                                                                                                                                                                                                                                                   |
| **sorting**                         |                         |                                                                                                                                                                                                                                                                   |
| `-s`<br>`--sort-type="lex"`         | `SORT_TYPE`             | How the bed file should be sorted.<br>- lex = lexicographic sorting (chr: 1 < 10 < 2 < MT < X)<br>- nat = natural sorting (chr: 1 < 2 < 10 < MT < X)<br>- ccs = custom chromosome sorting (see `--chr-order flag`)                                                |
| `--chr-order=CHR-ORDER,...`         | `CHR_ORDER`             | Comma separated custom chromosome order, to be used with custom chromosome sorting (`--sort-type=ccs`). Chromosomes not on the list will be sorted naturally after the ones in the list. If none is provided human chromosome order will be used (1-21, X, Y, MT) |
| `-d`<br>`--deduplicate`             | `DEDUPLICATE`           | Remove duplicated lines                                                                                                                                                                                                                                           |
|                                     |                         |                                                                                                                                                                                                                                                                   |
| **merging**                         |                         |                                                                                                                                                                                                                                                                   |
| `--no-merge`                        | `NO_MERGE`              | Do not merge bed regions. Note that touching regions are merged (e.g. if two regions are on the same chr they will be merged if one ends at 5 and the other starts at 6)                                                                                          |
| `--overlap=0`                       | `OVERLAP`               | Overlap between regions to be merged. This can be a positive or negative number (e.g. if you don't want touching regions to be merged set overlap to -1)                                                                                                          |
