package bed

import (
	"testing"

	"github.com/go-test/deep"
)

func TestPadLines(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
		shouldFail  bool
	}
	testCases := []testCase{
		{
			testing: "padding within chromosome",
			bed: Bedfile{
				Padding: 10,
				Lines: []Line{
					{
						Chr: "1", Start: 50, Stop: 51,
						Full: []string{"1", "50", "51"},
					},
					{
						Chr: "2", Start: 150, Stop: 151,
						Full: []string{"2", "150", "151"},
					},
					{
						Chr: "3", Start: 250, Stop: 251,
						Full: []string{"3", "250", "251"},
					},
					{
						Chr: "4", Start: 350, Stop: 351,
						Full: []string{"4", "350", "351"},
					},
				},
				chrLengthMap: map[string]int{
					"1": 100,
					"2": 200,
					"3": 300,
					"4": 400,
				},
			},
			expectedBed: Bedfile{
				Padding: 10,
				Lines: []Line{
					{
						Chr: "1", Start: 40, Stop: 61,
						Full: []string{"1", "40", "61"},
					},
					{
						Chr: "2", Start: 140, Stop: 161,
						Full: []string{"2", "140", "161"},
					},
					{
						Chr: "3", Start: 240, Stop: 261,
						Full: []string{"3", "240", "261"},
					},
					{
						Chr: "4", Start: 340, Stop: 361,
						Full: []string{"4", "340", "361"},
					},
				},
				chrLengthMap: map[string]int{
					"1": 100,
					"2": 200,
					"3": 300,
					"4": 400,
				},
			},
		},
		{
			testing: "padding beyond chromosome",
			bed: Bedfile{
				Padding: 1000,
				Lines: []Line{
					{
						Chr: "1", Start: 50, Stop: 51,
						Full: []string{"1", "50", "51"},
					},
					{
						Chr: "2", Start: 150, Stop: 151,
						Full: []string{"2", "150", "151"},
					},
					{
						Chr: "3", Start: 250, Stop: 251,
						Full: []string{"3", "250", "251"},
					},
					{
						Chr: "4", Start: 350, Stop: 351,
						Full: []string{"4", "350", "351"},
					},
				},
				chrLengthMap: map[string]int{
					"1": 100,
					"2": 200,
					"3": 300,
					"4": 400,
				},
			},
			expectedBed: Bedfile{
				Padding: 1000,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 100,
						Full: []string{"1", "1", "100"},
					},
					{
						Chr: "2", Start: 1, Stop: 200,
						Full: []string{"2", "1", "200"},
					},
					{
						Chr: "3", Start: 1, Stop: 300,
						Full: []string{"3", "1", "300"},
					},
					{
						Chr: "4", Start: 1, Stop: 400,
						Full: []string{"4", "1", "400"},
					},
				},
				chrLengthMap: map[string]int{
					"1": 100,
					"2": 200,
					"3": 300,
					"4": 400,
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			tc.bed.PadLines()
			if !tc.shouldFail {
				if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
					t.Error("expected VS received bed", diff)
				}
			}
		})
	}
}

func TestPadLine(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing          string
		bed              Bedfile
		line             Line
		expectedLine     Line
		expectedChrInMap bool
	}
	testCases := []testCase{
		{
			testing: "padding within chromosome",
			bed: Bedfile{
				Padding: 10,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 40, Stop: 61,
				Full: []string{"1", "40", "61"},
			},
			expectedChrInMap: true,
		},
		{
			testing: "padding beyond chromosome",
			bed: Bedfile{
				Padding: 1000,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 1, Stop: 100,
				Full: []string{"1", "1", "100"},
			},
			expectedChrInMap: true,
		},
		{
			testing: "chromosome not part of chrLengthMap",
			bed: Bedfile{
				Padding: 10,
				chrLengthMap: map[string]int{
					"2": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 40, Stop: 61,
				Full: []string{"1", "40", "61"},
			},
			expectedChrInMap: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			paddedLine, chrInMap := tc.bed.padLine(tc.line)
			if diff := deep.Equal(tc.expectedLine, paddedLine); diff != nil {
				t.Error("expected VS received line", diff)
			}
			if (!tc.expectedChrInMap && chrInMap) || (tc.expectedChrInMap && !chrInMap) {
				t.Fatalf("expectedChrInMap is %t, but chrInMap is %t", tc.expectedChrInMap, chrInMap)
			}
		})
	}
}
