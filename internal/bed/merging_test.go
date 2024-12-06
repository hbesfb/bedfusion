package bed

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

var testMergeChrOnly = []Line{
	{
		Chr: "1", Start: 1, Stop: 4,
		Full: []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Full: []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Full: []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Full: []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Full: []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Full: []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Full: []string{"1", "20", "30", "1", "A"},
	},
}

var testMergeChrStrand = []Line{
	{
		Chr: "1", Start: 1, Stop: 4,
		Strand: "1",
		Full:   []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1",
		Full:   []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Strand: "1",
		Full:   []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "-1",
		Full:   []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Strand: "1",
		Full:   []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1",
		Full:   []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Strand: "1",
		Full:   []string{"1", "20", "30", "1", "A"},
	},
}

var testMergeChrFeat = []Line{

	{
		Chr: "1", Start: 1, Stop: 4,
		Feat: "A",
		Full: []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Feat: "A",
		Full: []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Feat: "A",
		Full: []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Feat: "A",
		Full: []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Feat: "A",
		Full: []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Feat: "B",
		Full: []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Feat: "A",
		Full: []string{"1", "20", "30", "1", "A"},
	},
}

var testMergeFull = []Line{
	{
		Chr: "1", Start: 1, Stop: 4,
		Strand: "1", Feat: "A",
		Full: []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1", Feat: "A",
		Full: []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Strand: "1", Feat: "A",
		Full: []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "-1", Feat: "A",
		Full: []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Strand: "1", Feat: "A",
		Full: []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1", Feat: "B",
		Full: []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Strand: "1", Feat: "A",
		Full: []string{"1", "20", "30", "1", "A"},
	},
}

func TestMergeLines(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
	}
	testCases := []testCase{
		{
			testing: "testMergeChrOnly",
			bed: Bedfile{
				Lines: deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 8,
						Full: []string{"1", "1", "8", "1,-1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, overlap -1",
			bed: Bedfile{
				Overlap: -1,
				Lines:   deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Overlap: -1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 4,
						Full: []string{"1", "1", "4", "1", "A"},
					},
					{
						Chr: "1", Start: 5, Stop: 8,
						Full: []string{"1", "5", "8", "1,-1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, overlap 10",
			bed: Bedfile{
				Overlap: 10,
				Lines:   deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Overlap: 10,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 8,
						Full: []string{"1", "1", "8", "1,-1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, overlap 11",
			bed: Bedfile{
				Overlap: 11,
				Lines:   deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Overlap: 11,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 30,
						Full: []string{"1", "1", "30", "1,-1", "A,B"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrStrand",
			bed: Bedfile{
				StrandCol: 4 - 1,
				Lines:     deepCopyLines(testMergeChrStrand),
			},
			expectedBed: Bedfile{
				StrandCol: 4 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 5, Stop: 8,
						Strand: "-1",
						Full:   []string{"1", "5", "8", "-1", "A"},
					},
					{
						Chr: "1", Start: 1, Stop: 8,
						Strand: "1",
						Full:   []string{"1", "1", "8", "1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Strand: "1",
						Full:   []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Strand: "1",
						Full:   []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrFeat",
			bed: Bedfile{
				FeatCol: 5 - 1,
				Lines:   deepCopyLines(testMergeChrFeat),
			},
			expectedBed: Bedfile{
				FeatCol: 5 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 8,
						Feat: "A",
						Full: []string{"1", "1", "8", "1,-1", "A"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Feat: "A",
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Feat: "A",
						Full: []string{"2", "6", "8", "1", "A"},
					},
					{
						Chr: "1", Start: 5, Stop: 8,
						Feat: "B",
						Full: []string{"1", "5", "8", "1", "B"},
					},
				},
			},
		},
		{
			testing: "testMergeFull",
			bed: Bedfile{
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
				Lines:     deepCopyLines(testMergeFull),
			},
			expectedBed: Bedfile{
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 5, Stop: 8,
						Strand: "-1", Feat: "A",
						Full: []string{"1", "5", "8", "-1", "A"},
					},
					{
						Chr: "1", Start: 1, Stop: 8,
						Strand: "1", Feat: "A",
						Full: []string{"1", "1", "8", "1", "A"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Strand: "1", Feat: "A",
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Strand: "1", Feat: "A",
						Full: []string{"2", "6", "8", "1", "A"},
					},
					{
						Chr: "1", Start: 5, Stop: 8,
						Strand: "1", Feat: "B",
						Full: []string{"1", "5", "8", "1", "B"},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			tc.bed.MergeLines()
			if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
				t.Error("expected VS received bed", diff)
			}
		})
	}
}

func TestIdxInSlice(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing        string
		slice          []string
		item           string
		expectedResult int
	}
	testCases := []testCase{
		{
			testing:        "not in slice",
			slice:          []string{"10", "11", "1000"},
			item:           "1",
			expectedResult: -1,
		},
		{
			testing:        "in slice",
			slice:          []string{"10", "11", "1000"},
			item:           "11",
			expectedResult: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		description := fmt.Sprintf("%s in %v", tc.item, tc.slice)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			result := idxInSlice(tc.slice, tc.item)
			if tc.expectedResult != result {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}
