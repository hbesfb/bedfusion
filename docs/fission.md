# Fission

The `-f/--fission` options split regions, instead of merging regions which is the default behaviour of BedFusion. Note that bed files will still be [sorted](./sorting.md).

Example bed file `examples/fission-test.bed`:

``` text
1	1	1000	1	A
1	1	1000	-1	B
2	1	1000	1	C
```

Example: 

``` shell 
> bedfusion examples/fission-test.bed -f
1       1       101     1       A
1       1       101     -1      B
1       102     202     1       A
1       102     202     -1      B
1       203     303     1       A
1       203     303     -1      B
1       304     404     1       A
1       304     404     -1      B
1       405     505     1       A
1       405     505     -1      B
1       506     606     1       A
1       506     606     -1      B
1       607     707     1       A
1       607     707     -1      B
1       708     808     1       A
1       708     808     -1      B
1       809     909     1       A
1       809     909     -1      B
1       910     1000    1       A
1       910     1000    -1      B
2       1       101     1       C
2       102     202     1       C
2       203     303     1       C
2       304     404     1       C
2       405     505     1       C
2       506     606     1       C
2       607     707     1       C
2       708     808     1       C
2       809     909     1       C
2       910     1000    1       C
```

## Using split size

BedFusion `-f/--fission` splits regions into new regions of 100bp as default (`--split-size=100`), but one can choose any split size using the `--split-size` flag. 

For example if one would like bigger regions of 499bp:

``` shell
1       1       500     1       A
1       1       500     -1      B
1       501     1000    1       A
1       501     1000    -1      B
2       1       500     1       C
2       501     1000    1       C
```
