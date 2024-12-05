package bed

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

var testChrSort = []Line{
	{
		Chr: "chr10", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "chrX", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "HG987_PATCH", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "chr10", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "chrMT", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "GL000209.1", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "HG385_PATCH", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "chr2", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "GL000226.1", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
	{
		Chr: "chr1", Start: 8, Stop: 9,
		Full: []string{"1", "8", "9"},
	},
}

var testFullSort = []Line{
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
	{
		Chr: "1", Start: 8, Stop: 9,
		Strand: "-1", Feat: "B",
		Full: []string{"1", "8", "9", "-1", "B"},
	},
	{
		Chr: "MT", Start: 10, Stop: 11,
		Strand: "1", Feat: "A",
		Full: []string{"MT", "10", "11", "1", "A"},
	},
	{
		Chr: "10", Start: 12, Stop: 13,
		Strand: "1", Feat: "D",
		Full: []string{"10", "12", "13", "1", "D"},
	},
	{
		Chr: "GL000209.1", Start: 10, Stop: 11,
		Strand: "1", Feat: "A",
		Full: []string{"GL000209.1", "10", "11", "1", "A"},
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
		Chr: "HG385_PATCH", Start: 10, Stop: 11,
		Strand: "1", Feat: "A",
		Full: []string{"X", "10", "11", "1", "A"},
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
		lines         []Line
		expectedLines []Line
	}
	testCases := []testCase{
		{
			testing: "chr sort",
			lines:   deepCopyLines(testChrSort),
			expectedLines: []Line{
				{
					Chr: "chr1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr2", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrMT", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrX", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "GL000209.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "GL000226.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG385_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG987_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
			},
		},
		{
			testing: "full sort",
			lines:   deepCopyLines(testFullSort),
			expectedLines: []Line{
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
					Chr: "GL000209.1", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"GL000209.1", "10", "11", "1", "A"},
				},
				{
					Chr: "HG385_PATCH", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"X", "10", "11", "1", "A"},
				},
				{
					Chr: "MT", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"MT", "10", "11", "1", "A"},
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

func TestNaturalSort(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing       string
		lines         []Line
		expectedLines []Line
	}
	testCases := []testCase{
		{
			testing: "chr sort",
			lines:   deepCopyLines(testChrSort),
			expectedLines: []Line{
				{
					Chr: "chr1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr2", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrMT", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrX", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "GL000209.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "GL000226.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG385_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG987_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
			},
		},
		{
			testing: "full sort",
			lines:   deepCopyLines(testFullSort),
			expectedLines: []Line{
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
					Chr: "2", Start: 12, Stop: 13,
					Strand: "1", Feat: "C",
					Full: []string{"2", "12", "13", "1", "C"},
				},
				{
					Chr: "10", Start: 12, Stop: 13,
					Strand: "1", Feat: "D",
					Full: []string{"10", "12", "13", "1", "D"},
				},
				{
					Chr: "GL000209.1", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"GL000209.1", "10", "11", "1", "A"},
				},
				{
					Chr: "HG385_PATCH", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"X", "10", "11", "1", "A"},
				},
				{
					Chr: "MT", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"MT", "10", "11", "1", "A"},
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
			receivedLines := naturalSort(tc.lines)
			if diff := deep.Equal(tc.expectedLines, receivedLines); diff != nil {
				t.Error("expected VS received lines", diff)
			}
		})
	}
}

func TestHumanSort(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing       string
		lines         []Line
		expectedLines []Line
	}
	testCases := []testCase{
		{
			testing: "chr sort",
			lines:   deepCopyLines(testChrSort),
			expectedLines: []Line{
				{
					Chr: "chr1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr2", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrX", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrMT", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "GL000209.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "GL000226.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG385_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG987_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
			},
		},
		{
			testing: "full sort",
			lines:   deepCopyLines(testFullSort),
			expectedLines: []Line{
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
					Chr: "2", Start: 12, Stop: 13,
					Strand: "1", Feat: "C",
					Full: []string{"2", "12", "13", "1", "C"},
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
					Chr: "MT", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"MT", "10", "11", "1", "A"},
				},
				{
					Chr: "GL000209.1", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"GL000209.1", "10", "11", "1", "A"},
				},
				{
					Chr: "HG385_PATCH", Start: 10, Stop: 11,
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
			receivedLines := humanSort(tc.lines)
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
		lines         []Line
		expectedLines []Line
	}
	testCases := []testCase{
		{
			testing: "chr sort",
			lines:   deepCopyLines(testChrSort),
			expectedLines: []Line{
				{
					Chr: "GL000209.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "GL000226.1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG385_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "HG987_PATCH", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr1", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr10", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chr2", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrMT", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
				{
					Chr: "chrX", Start: 8, Stop: 9,
					Full: []string{"1", "8", "9"},
				},
			},
		},
		{
			testing: "full sort",
			lines:   deepCopyLines(testFullSort),
			expectedLines: []Line{
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
					Chr: "GL000209.1", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"GL000209.1", "10", "11", "1", "A"},
				},
				{
					Chr: "HG385_PATCH", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"X", "10", "11", "1", "A"},
				},
				{
					Chr: "MT", Start: 10, Stop: 11,
					Strand: "1", Feat: "A",
					Full: []string{"MT", "10", "11", "1", "A"},
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

func TestNaturalStringCompare(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a              string
		b              string
		expectedResult int
	}
	testCases := []testCase{
		{a: "a", b: "b", expectedResult: -1},
		{a: "chr1", b: "chr2", expectedResult: -1},
		{a: "chr1", b: "chrX", expectedResult: -1},
		{a: "chr10", b: "chr2", expectedResult: 1},
		{a: "chr10", b: "chrX", expectedResult: -1},
		{a: "chrY", b: "chr1", expectedResult: 1},
		{a: "chrMT", b: "chrX", expectedResult: -1},
		{a: "chrMT", b: "GL000209.1", expectedResult: -1},
		{a: "1", b: "2", expectedResult: -1},
		{a: "1", b: "X", expectedResult: -1},
		{a: "10", b: "2", expectedResult: 1},
		{a: "10", b: "X", expectedResult: -1},
		{a: "Y", b: "1", expectedResult: 1},
		{a: "MT", b: "X", expectedResult: -1},
		{a: "MT", b: "GL000209.1", expectedResult: 1},
		{a: "HGNC:10", b: "HGNC:2", expectedResult: 1},
	}
	for _, tc := range testCases {
		tc := tc
		description := fmt.Sprintf("%s vs %s", tc.a, tc.b)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			result := naturalStringCompare(tc.a, tc.b)
			if tc.expectedResult != result {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
			}
		})
	}
}

func TestHumanChrCompare(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a              string
		b              string
		expectedResult int
	}
	testCases := []testCase{
		{a: "a", b: "b", expectedResult: -1},
		{a: "chr1", b: "chr2", expectedResult: -1},
		{a: "chr1", b: "chrX", expectedResult: -1},
		{a: "chr10", b: "chr2", expectedResult: 1},
		{a: "chr10", b: "chrX", expectedResult: -1},
		{a: "chrY", b: "chr1", expectedResult: 1},
		{a: "chrMT", b: "chrX", expectedResult: 1},
		{a: "chrMT", b: "GL000209.1", expectedResult: -1},
		{a: "1", b: "2", expectedResult: -1},
		{a: "1", b: "X", expectedResult: -1},
		{a: "10", b: "2", expectedResult: 1},
		{a: "10", b: "X", expectedResult: -1},
		{a: "Y", b: "1", expectedResult: 1},
		{a: "MT", b: "X", expectedResult: 1},
		{a: "MT", b: "GL000209.1", expectedResult: -1},
		{a: "HGNC:10", b: "HGNC:2", expectedResult: 1},
	}
	for _, tc := range testCases {
		tc := tc
		description := fmt.Sprintf("%s vs %s", tc.a, tc.b)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			result := stringSliceCompare(tc.a, tc.b, humanChrOrder)
			if tc.expectedResult != result {
				t.Errorf("expected %d got %d", tc.expectedResult, result)
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
