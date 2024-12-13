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

func TestVerifyAndHandle(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
		shouldFail  bool
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
			testing: "correct input with strand col",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				Output:    "/some/output/path/output.bed",
				StrandCol: 4,
			},
			expectedBed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				Output:    "/some/output/path/output.bed",
				StrandCol: 3,
			},
		},
		{
			testing: "correct input with feat col",
			bed: Bedfile{
				Inputs:  []string{"/some/path/test.bed"},
				Output:  "/some/output/path/output.bed",
				FeatCol: 3,
			},
			expectedBed: Bedfile{
				Inputs:  []string{"/some/path/test.bed"},
				Output:  "/some/output/path/output.bed",
				FeatCol: 2,
			},
		},
		{
			testing: "correct input with both cols",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				Output:    "/some/output/path/output.bed",
				StrandCol: 4,
				FeatCol:   3,
			},
			expectedBed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				Output:    "/some/output/path/output.bed",
				StrandCol: 3,
				FeatCol:   2,
			},
		},
		{
			testing: "unclean paths",
			bed: Bedfile{
				Inputs:    []string{"/some/../path/test1.bed", "./some/../path/./test2.bed"},
				Output:    "/some/./output/./path/./output.bed",
				StrandCol: 4,
				FeatCol:   3,
			},
			expectedBed: Bedfile{
				Inputs:    []string{"/path/test1.bed", "path/test2.bed"},
				Output:    "/some/output/path/output.bed",
				StrandCol: 3,
				FeatCol:   2,
			},
		},
		{
			testing: "sortType is ccs, and chrOrder is empty",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				SortType: "ccs",
			},
			expectedBed: Bedfile{
				Inputs:      []string{"/some/path/test.bed"},
				SortType:    "ccs",
				ChrOrder:    humanChrOrder,
				chrOrderMap: chrOrderToMap(humanChrOrder),
			},
		},
		{
			testing: "sortType is ccs, and is set",
			bed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				SortType: "ccs",
				ChrOrder: []string{"4", "3", "2", "1"},
			},
			expectedBed: Bedfile{
				Inputs:   []string{"/some/path/test.bed"},
				SortType: "ccs",
				ChrOrder: []string{"4", "3", "2", "1"},
				chrOrderMap: map[string]int{
					"4": 1,
					"3": 2,
					"2": 3,
					"1": 4,
				},
			},
		},
		{
			testing: "strand col less than 3",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				Output:    "/some/output/path/output.bed",
				StrandCol: 2,
				FeatCol:   3,
			},
			shouldFail: true,
		},
		{
			testing: "feat col less than 3",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				Output:    "/some/output/path/output.bed",
				StrandCol: 4,
				FeatCol:   2,
			},
			shouldFail: true,
		},
		{
			testing: "overlapping strand and feat cols",
			bed: Bedfile{
				Inputs:    []string{"/some/path/test.bed"},
				Output:    "/some/output/path/output.bed",
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
			err := tc.bed.VerifyAndHandle()
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
