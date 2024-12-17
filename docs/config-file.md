# Configuration file

BedFusion supports the possibility to set options in a configuration file. This can be a very good option for documentation purposes and if one for example always work with bed files of the same format. 

Note that the options in the yaml file will match the flags.

Example configuration file `examples/config-test.yml`:

``` yaml
strand-col: 4
feat-col: 5
sort-type: ccs
chr-order: x,y,mt
overlap: 10
no-merge: false
```

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

Used together:

``` shell
> bedfusion examples/sort-test.bed -c examples/config-test.yml
X       10      11      1       A
Y       10      11      1       A
MT      10      11      1       A
1       8       11      -1      B
1       10      11      -1      A
1       10      13      1       A
2       12      13      1       C
10      12      13      1       D
GL000209.1      10      11      1       A
```

## Using a configuration file together with flags and environmental variables

The read priority order is

1. flags 
1. configuration file
3. environmental variables

Since flags have a higher priority over than the configuration file we can overwrite the configuration file settings by for example removing the strand and feature columns:

``` shell
> bedfusion examples/sort-test.bed -c examples/config-test.yml --strand-col=0 --feat-col=0
X       10      11      1       A
Y       10      11      1       A
MT      10      11      1       A
1       8       13      -1,1    B,A
2       12      13      1       C
10      12      13      1       D
GL000209.1      10      11      1       A
```

Using environmental variables on the other hand, one can not overwrite the configuration file, as environmental variables have a lower read priority:

```shell
> STRAND_COL=0 FEAT_COL=0 bedfusion examples/sort-test.bed -c examples/config-test.yml
X       10      11      1       A
Y       10      11      1       A
MT      10      11      1       A
1       8       11      -1      B
1       10      11      -1      A
1       10      13      1       A
2       12      13      1       C
10      12      13      1       D
GL000209.1      10      11      1       A
```
