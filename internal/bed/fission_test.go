package bed

import (
	"testing"

	"github.com/go-test/deep"
)

func TestSplitLines(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
	}
	testCases := []testCase{
		{
			testing: "split bed",
			bed: Bedfile{
				SplitSize: 24,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 100,
						Strand: "1", Feat: "A",
						Full: []string{"1", "1", "100", "1", "A"},
					},
					{
						Chr: "2", Start: 1, Stop: 25,
						Strand: "-1", Feat: "A",
						Full: []string{"2", "1", "25", "-1", "A"},
					},
					{
						Chr: "X", Start: 1, Stop: 30,
						Strand: "1", Feat: "C",
						Full: []string{"X", "1", "30", "1", "C"},
					},
				},
			},
			expectedBed: Bedfile{
				SplitSize: 24,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 25,
						Strand: "1", Feat: "A",
						Full: []string{"1", "1", "25", "1", "A"},
					},
					{
						Chr: "1", Start: 26, Stop: 50,
						Strand: "1", Feat: "A",
						Full: []string{"1", "26", "50", "1", "A"},
					},
					{
						Chr: "1", Start: 51, Stop: 75,
						Strand: "1", Feat: "A",
						Full: []string{"1", "51", "75", "1", "A"},
					},
					{
						Chr: "1", Start: 76, Stop: 100,
						Strand: "1", Feat: "A",
						Full: []string{"1", "76", "100", "1", "A"},
					},
					{
						Chr: "2", Start: 1, Stop: 25,
						Strand: "-1", Feat: "A",
						Full: []string{"2", "1", "25", "-1", "A"},
					},
					{
						Chr: "X", Start: 1, Stop: 25,
						Strand: "1", Feat: "C",
						Full: []string{"X", "1", "25", "1", "C"},
					},
					{
						Chr: "X", Start: 26, Stop: 30,
						Strand: "1", Feat: "C",
						Full: []string{"X", "26", "30", "1", "C"},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			tc.bed.SplitLines()
			if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
				t.Error("expected VS received bed", diff)
			}
		})
	}
}

func TestReplaceStartAndStop(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing      string
		start        int
		stop         int
		line         Line
		expectedLine Line
	}
	testCases := []testCase{
		{
			testing: "replace start and stop",
			start:   26,
			stop:    50,
			line: Line{
				Chr: "1", Start: 1, Stop: 100,
				Strand: "1", Feat: "A",
				Full: []string{"1", "1", "100", "1", "A"},
			},
			expectedLine: Line{
				Chr: "1", Start: 26, Stop: 50,
				Strand: "1", Feat: "A",
				Full: []string{"1", "26", "50", "1", "A"},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			receivedLine := tc.line.replaceStartAndStop(tc.start, tc.stop)
			if diff := deep.Equal(tc.expectedLine, receivedLine); diff != nil {
				t.Error("expected VS received line", diff)
			}
		})
	}
}
