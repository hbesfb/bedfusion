package bed

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

func TestMain(m *testing.M) {
	// Compare unexported fields in structs
	deep.CompareUnexportedFields = true
	os.Exit(m.Run())
}

func TestVerifyAndHandleColumns(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
		shouldFail  bool
	}
	testCases := []testCase{
		{
			testing: "correct input with strand col",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				StrandCol: 4,
			},
			expectedBed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				StrandCol: 3,
			},
		},
		{
			testing: "correct input with feat col",
			bed: Bedfile{
				Inputs:  []string{"/some/path/test.bed"},
				FeatCol: 3,
			},
			expectedBed: Bedfile{
				Inputs:  []string{"/some/path/test.bed"},
				FeatCol: 2,
			},
		},
		{
			testing: "correct input with both cols",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				StrandCol: 4,
				FeatCol:   3,
			},
			expectedBed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				StrandCol: 3,
				FeatCol:   2,
			},
		},
		{
			testing: "strand col less than 3",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				StrandCol: 2,
				FeatCol:   3,
			},
			shouldFail: true,
		},
		{
			testing: "feat col less than 3",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				StrandCol: 4,
				FeatCol:   2,
			},
			shouldFail: true,
		},
		{
			testing: "overlapping strand and feat cols",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				StrandCol: 4,
				FeatCol:   4,
			},
			shouldFail: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			err := tc.bed.verifyAndHandleColumns()
			if (!tc.shouldFail && err != nil) || (tc.shouldFail && err == nil) {
				t.Fatalf("shouldFail is %t, but err is %q", tc.shouldFail, err)
			}
			if !tc.shouldFail {
				if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
					t.Error("expected VS received bed", diff)
				}
			}
		})
	}
}

func TestVerifyFastaIdxCombinations(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing    string
		bed        Bedfile
		shouldFail bool
	}
	testCases := []testCase{
		{
			testing: "correct input, with both padding and fasta-idx selected",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				FastaIdx: "/some/fasta/idx/file.fasta.fai",
				Padding:  2,
			},
		},
		{
			testing: "correct input, with both sorting type == fidx and fasta-idx selected",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				FastaIdx: "/some/fasta/idx/file.fasta.fai",
				SortType: FidxST,
			},
		},
		{
			testing: "correct input, with padding, sorting type == fidx and fasta-idx selected",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				FastaIdx: "/some/fasta/idx/file.fasta.fai",
				Padding:  2,
				SortType: FidxST,
			},
		},
		{
			testing: "padding selected, but missing fasta index file",
			bed: Bedfile{
				Inputs:  []string{"/some/path/test.bed"},
				Padding: 2,
			},
			shouldFail: true,
		},
		{
			testing: "padding of type safe selected, but missing fasta index file",
			bed: Bedfile{
				Inputs:      []string{"/some/path/test.bed"},
				Padding:     2,
				PaddingType: SafePT,
			},
			shouldFail: true,
		},
		{
			testing: "padding of type lax selected, but missing fasta index file",
			bed: Bedfile{
				Inputs:      []string{"/some/path/test.bed"},
				Padding:     2,
				PaddingType: LaxPT,
			},
			shouldFail: true,
		},
		{
			testing: "padding of type force selected, but missing fasta index file",
			bed: Bedfile{
				Inputs:      []string{"/some/path/test.bed"},
				Padding:     2,
				PaddingType: ForcePT,
			},
			shouldFail: false,
		},
		{
			testing: "sorting type == fidx and fasta-idx selected, but missing fasta index file",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				FastaIdx: "/some/fasta/idx/file.fasta.fai",
				SortType: FidxST,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			err := tc.bed.verifyFastaIdxCombinations()
			if (!tc.shouldFail && err != nil) || (tc.shouldFail && err == nil) {
				t.Fatalf("shouldFail is %t, but err is %q", tc.shouldFail, err)
			}
		})
	}
}

func TestVerifyFirstBase(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing    string
		bed        Bedfile
		shouldFail bool
	}
	testCases := []testCase{
		{
			testing: "correct first base is 0",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				FirstBase: 0,
			},
		},
		{
			testing: "correct first base is 1",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				FirstBase: 1,
			},
		},
		{
			testing: "wrong first base is -1",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				FirstBase: -1,
			},
			shouldFail: true,
		},
		{
			testing: "wrong first base is 2",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				FirstBase: 2,
			},
			shouldFail: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			err := tc.bed.verifyFirstBase()
			if (!tc.shouldFail && err != nil) || (tc.shouldFail && err == nil) {
				t.Fatalf("shouldFail is %t, but err is %q", tc.shouldFail, err)
			}
		})
	}
}

func TestHandleCCSSorting(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
	}
	testCases := []testCase{
		{
			testing: "sortType is ccs, and chrOrder is empty",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				SortType: CcsST,
			},
			expectedBed: Bedfile{
				Inputs:      []string{"/some/path/test.bed"},
				SortType:    CcsST,
				ChrOrder:    humanChrOrder,
				chrOrderMap: chrOrderToMap(humanChrOrder),
			},
		},
		{
			testing: "sortType is ccs, and is set",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				SortType: CcsST,
				ChrOrder: []string{"4", "3", "2", "1"},
			},
			expectedBed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				SortType: CcsST,
				ChrOrder: []string{"4", "3", "2", "1"},
				chrOrderMap: map[string]int{
					"4": 1,
					"3": 2,
					"2": 3,
					"1": 4,
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			tc.bed.handleCCSSorting()
			if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
				t.Error("expected VS received bed", diff)
			}
		})
	}
}

func TestCleanPaths(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
	}
	testCases := []testCase{
		{
			testing: "correct input, only input path",
			bed: Bedfile{
				Inputs: []string{"/some/path/test.bed"},
			},
			expectedBed: Bedfile{
				Inputs: []string{"/some/path/test.bed"},
			},
		},
		{
			testing: "unclean paths",
			bed: Bedfile{
				Inputs:   []string{"/some/../path/test1.bed", "./some/../path/./test2.bed"},
				Output:   "/some/./output/./path/./output.bed",
				FastaIdx: "some/../some/./fasta/idx/path/./file.fasta.fai",
			},
			expectedBed: Bedfile{
				Inputs:   []string{"/path/test1.bed", "path/test2.bed"},
				Output:   "/some/output/path/output.bed",
				FastaIdx: "some/fasta/idx/path/file.fasta.fai",
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			tc.bed.cleanPaths()
			if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
				t.Error("expected VS received bed", diff)
			}
		})
	}
}

func TestChrOrderMap(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing             string
		slice               []string
		expectedChrOrderMap map[string]int
	}
	testCases := []testCase{
		{
			testing: "all lowercase",
			slice:   []string{"chr1", "chr2", "chr3", "chr4"},
			expectedChrOrderMap: map[string]int{
				"chr1": 1,
				"chr2": 2,
				"chr3": 3,
				"chr4": 4,
			},
		},
		{
			testing: "mixture of lowercase and uppercase",
			slice:   []string{"chr1", "chrX", "chrY", "chrMT"},
			expectedChrOrderMap: map[string]int{
				"chr1":  1,
				"chrx":  2,
				"chry":  3,
				"chrmt": 4,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			chrOrderMap := chrOrderToMap(tc.slice)
			if diff := deep.Equal(tc.expectedChrOrderMap, chrOrderMap); diff != nil {
				t.Error("expected VS received chrOrderMap", diff)
			}
		})
	}
}

// --- Test Helper Functions ---

// To make deep copies of Lines
func deepCopyLines(lines []Line) []Line {
	var copiedLines []Line
	for _, l := range lines {
		fullLineCopy := make([]string, len(l.Full))
		_ = copy(fullLineCopy, l.Full)
		copiedLine := Line{
			Chr:    l.Chr,
			Start:  l.Start,
			Stop:   l.Stop,
			Strand: l.Strand,
			Feat:   l.Feat,
			Full:   fullLineCopy,
		}
		copiedLines = append(copiedLines, copiedLine)
	}
	return copiedLines
}

// To make deep copy line
func deepCopyLine(line Line) Line {
	fullLineCopy := make([]string, len(line.Full))
	_ = copy(fullLineCopy, line.Full)
	copiedLine := Line{
		Chr:    line.Chr,
		Start:  line.Start,
		Stop:   line.Stop,
		Strand: line.Strand,
		Feat:   line.Feat,
		Full:   fullLineCopy,
	}
	return copiedLine
}
