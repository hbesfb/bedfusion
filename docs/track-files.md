# Track file support

BedFusion has limited track file support and will only handle track under the following conditions:

1. Headers lines are only in the top of the file
1. A header lines start with one of the following words or symbols:
    1. "browser"
    1. "track"
    1. "#"
1. If more than one file is used as input only the first track file can contain header lines

Example track file `track-test.bed` (example taken from [Genome Browser: Data File Formats - BED format](https://genome.ucsc.edu/FAQ/FAQformat.html#format1)):

``` text
browser position chr7:127471196-127495720
browser hide all
track name="ItemRGBDemo" description="Item RGB demonstration" visibility=2 itemRgb="On"
chr7	127471196	127472363	Pos1	0	+	127471196	127472363	255,0,0
chr7	127472363	127473530	Pos2	0	+	127472363	127473530	255,0,0
chr7	127473530	127474697	Pos3	0	+	127473530	127474697	255,0,0
chr7	127474697	127475864	Pos4	0	+	127474697	127475864	255,0,0
chr7	127475864	127477031	Neg1	0	-	127475864	127477031	0,0,255
chr7	127477031	127478198	Neg2	0	-	127477031	127478198	0,0,255
chr7	127478198	127479365	Neg3	0	-	127478198	127479365	0,0,255
chr7	127479365	127480532	Pos5	0	+	127479365	127480532	255,0,0
chr7	127480532	127481699	Neg4	0	-	127480532	127481699	0,0,255
```

Example:

``` shell
> bedfusion track-test.bed --strand-col=6
browser position chr7:127471196-127495720
browser hide all
track name="ItemRGBDemo" description="Item RGB demonstration" visibility=2 itemRgb="On"
chr7    127471196       127475864       Pos1,Pos2,Pos3,Pos4     0       +       127471196,127472363,127473530,127474697 127472363,127473530,127474697,127475864 255,0,0,255,0,0,255,0,0,255,0,0
chr7    127475864       127479365       Neg1,Neg2,Neg3  0       -       127475864,127477031,127478198   127477031,127478198,127479365   0,0,255,0,0,255,0,0,255
chr7    127479365       127480532       Pos5    0       +       127479365       127480532       255,0,0
chr7    127480532       127481699       Neg4    0       -       127480532       127481699       0,0,255
```
