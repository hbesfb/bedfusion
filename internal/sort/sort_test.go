package sort

import (
	"os"
	"testing"

	"github.com/go-test/deep"
	"github.com/hbesfb/bedfusion/internal/bed"
)

func TestMain(m *testing.M) {
	// Compare unexported fields in structs
	deep.CompareUnexportedFields = true
	os.Exit(m.Run())
}

var linesToSort = []bed.Line{
	{
		Chr: "2", Start: 12, Stop: 13,
		Strand: "1", Feat: "C",
		Full: []string{"2", "12", "13", "1", "C"},
	},
	{
		Chr: "1", Start: 8, Stop: 9,
		Strand: "-1", Feat: "B",
		Full: []string{"1", "8", "9", "-1", "B"},
	},
	{
		Chr: "10", Start: 12, Stop: 13,
		Strand: "1", Feat: "D",
		Full: []string{"10", "12", "13", "1", "D"},
	},
	{
		Chr: "X", Start: 10, Stop: 11,
		Strand: "1", Feat: "A",
		Full: []string{"X", "10", "11", "1", "A"},
	},
	{
		Chr: "1", Start: 10, Stop: 11,
		Strand: "-1", Feat: "A",
		Full: []string{"1", "10", "11", "-1", "A"},
	},
	{
		Chr: "1", Start: 12, Stop: 13,
		Strand: "1", Feat: "A",
		Full: []string{"1", "12", "13", "1", "A"},
	},
	{
		Chr: "1", Start: 10, Stop: 11,
		Strand: "1", Feat: "A",
		Full: []string{"1", "10", "11", "1", "A"},
	},
	{
		Chr: "1", Start: 10, Stop: 11,
		Strand: "-1", Feat: "B",
		Full: []string{"1", "10", "11", "-1", "B"},
	},
}

func TestLexicographicSort(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing       string
		lines         []bed.Line
		expectedLines []bed.Line
	}
	testCases := []testCase{
		{
			testing: "linesToSort",
			lines:   linesToSort,
			expectedLines: []bed.Line{
				{
					Chr: "1", Start: 8, Stop: 9,
					Strand: "-1", Feat: "B",
					Full: []string{"1", "8", "9", "-1", "B"},
				},
				{
					Chr: "1", Start: 10, Stop: 11,
					Strand: "-1", Feat: "A",
					Full: []string{"1", "10", "11", "-1", "A"},
				},
				{
					Chr: "1", Start: 10, Stop: 11,
					Strand: "-1", Feat: "B",
					Full: []string{"1", "10", "11", "-1", "B"},
				},
				{
					Chr: "1", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"1", "10", "11", "1", "A"},
				},
				{
					Chr: "1", Start: 12, Stop: 13,
					Strand: "1", Feat: "A",
					Full: []string{"1", "12", "13", "1", "A"},
				},
				{
					Chr: "10", Start: 12, Stop: 13,
					Strand: "1", Feat: "D",
					Full: []string{"10", "12", "13", "1", "D"},
				},
				{
					Chr: "2", Start: 12, Stop: 13,
					Strand: "1", Feat: "C",
					Full: []string{"2", "12", "13", "1", "C"},
				},
				{
					Chr: "X", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"X", "10", "11", "1", "A"},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			receivedLines := lexicographicSort(tc.lines)
			if diff := deep.Equal(tc.expectedLines, receivedLines); diff != nil {
				t.Error("expected VS received lines", diff)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing       string
		lines         []bed.Line
		expectedLines []bed.Line
	}
	testCases := []testCase{
		{
			testing: "linesToSort",
			lines:   linesToSort,
			expectedLines: []bed.Line{
				{
					Chr: "1", Start: 10, Stop: 11,
					Strand: "-1", Feat: "A",
					Full: []string{"1", "10", "11", "-1", "A"},
				},
				{
					Chr: "1", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"1", "10", "11", "1", "A"},
				},
				{
					Chr: "1", Start: 12, Stop: 13,
					Strand: "1", Feat: "A",
					Full: []string{"1", "12", "13", "1", "A"},
				},
				{
					Chr: "X", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"X", "10", "11", "1", "A"},
				},
				{
					Chr: "1", Start: 8, Stop: 9,
					Strand: "-1", Feat: "B",
					Full: []string{"1", "8", "9", "-1", "B"},
				},
				{
					Chr: "1", Start: 10, Stop: 11,
					Strand: "-1", Feat: "B",
					Full: []string{"1", "10", "11", "-1", "B"},
				},
				{
					Chr: "2", Start: 12, Stop: 13,
					Strand: "1", Feat: "C",
					Full: []string{"2", "12", "13", "1", "C"},
				},
				{
					Chr: "10", Start: 12, Stop: 13,
					Strand: "1", Feat: "D",
					Full: []string{"10", "12", "13", "1", "D"},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			receivedLines := mergeSort(tc.lines)
			if diff := deep.Equal(tc.expectedLines, receivedLines); diff != nil {
				t.Error("expected VS received lines", diff)
			}
		})
	}
}
