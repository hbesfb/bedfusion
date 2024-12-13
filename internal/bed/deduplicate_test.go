package bed

import (
	"testing"

	"github.com/go-test/deep"
)

func TestDeduplicateLines(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
		shouldFail  bool
	}
	testCases := []testCase{
		{
			testing: "simple bed file, with duplicates",
			bed: Bedfile{
				Inputs: []string{"test.bed"},
				Lines: []Line{
					{
						Chr: "1", Start: 10, Stop: 100,
						Full: []string{"1", "10", "100"},
					},
					{
						Chr: "2", Start: 20, Stop: 200,
						Full: []string{"2", "20", "200"},
					},
					{
						Chr: "1", Start: 10, Stop: 100,
						Full: []string{"1", "10", "100"},
					},
					{
						Chr: "3", Start: 30, Stop: 300,
						Full: []string{"3", "30", "300"},
					},
					{
						Chr: "4", Start: 40, Stop: 400,
						Full: []string{"4", "40", "400"},
					},
					{
						Chr: "3", Start: 30, Stop: 300,
						Full: []string{"3", "30", "300"},
					},
				},
			},
			expectedBed: Bedfile{
				Inputs: []string{"test.bed"},
				Lines: []Line{
					{
						Chr: "1", Start: 10, Stop: 100,
						Full: []string{"1", "10", "100"},
					},
					{
						Chr: "2", Start: 20, Stop: 200,
						Full: []string{"2", "20", "200"},
					},
					{
						Chr: "3", Start: 30, Stop: 300,
						Full: []string{"3", "30", "300"},
					},
					{
						Chr: "4", Start: 40, Stop: 400,
						Full: []string{"4", "40", "400"},
					},
				},
			},
		},
		{
			testing: "complex bed file with strand and feat",
			bed: Bedfile{
				Inputs:    []string{"test.bed"},
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 10, Stop: 100,
						Strand: "-1", Feat: "A",
						Full: []string{"1", "10", "100", "-1", "A"},
					},
					{
						Chr: "2", Start: 20, Stop: 200,
						Strand: "-1", Feat: "B",
						Full: []string{"2", "20", "200", "-1", "B"},
					},
					{
						Chr: "3", Start: 30, Stop: 300,
						Strand: "1", Feat: "C",
						Full: []string{"3", "30", "300", "1", "C"},
					},
					{
						Chr: "4", Start: 40, Stop: 400,
						Strand: "1", Feat: "D",
						Full: []string{"4", "40", "400", "1", "D"},
					},
				},
			},
			expectedBed: Bedfile{
				Inputs:    []string{"test.bed"},
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 10, Stop: 100,
						Strand: "-1", Feat: "A",
						Full: []string{"1", "10", "100", "-1", "A"},
					},
					{
						Chr: "2", Start: 20, Stop: 200,
						Strand: "-1", Feat: "B",
						Full: []string{"2", "20", "200", "-1", "B"},
					},
					{
						Chr: "3", Start: 30, Stop: 300,
						Strand: "1", Feat: "C",
						Full: []string{"3", "30", "300", "1", "C"},
					},
					{
						Chr: "4", Start: 40, Stop: 400,
						Strand: "1", Feat: "D",
						Full: []string{"4", "40", "400", "1", "D"},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			tc.bed.DeduplicateLines()
			if !tc.shouldFail {
				if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
					t.Error("expected VS received bed", diff)
				}
			}
		})
	}
}
