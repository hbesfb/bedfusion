package sort

import (
	"fmt"
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

func TestLexicographicSort(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing       string
		lines         []bed.Line
		expectedLines []bed.Line
	}
	testCases := []testCase{
		{
			testing: "complex bed",
			lines: []bed.Line{
				{
					Chr: "2", Start: 31747549, Stop: 31763836,
					Strand: "-1", Feat: "ENSG00000049319",
					Full: []string{"2", "31747549", "31763836", "-1", "SRD5A2", "ENSG00000049319"},
				},
				{
					Chr: "10", Start: 126085871, Stop: 126107545,
					Strand: "-1", Feat: "ENSG00000065154",
					Full: []string{"10", "126085871", "126107545", "-1", "OAT", "ENSG00000065154"},
				},
				{
					Chr: "10", Start: 124768494, Stop: 124773587,
					Strand: "1", Feat: "ENSG00000196177",
					Full: []string{"10", "124768494", "124773587", "1", "ACADSB", "ENSG00000196177"},
				},
				{
					Chr: "1", Start: 948802, Stop: 949920,
					Strand: "1", Feat: "ENSG00000187608",
					Full: []string{"1", "948802", "949920", "1", "ISG15", "ENSG00000187608"},
				},
				{
					Chr: "X", Start: 135067597, Stop: 135129423,
					Strand: "1", Feat: "ENSG00000198689",
					Full: []string{"X", "135067597", "135129423", "1", "SLC9A6", "ENSG00000198689"},
				},
				{
					Chr: "2", Start: 1800689, Stop: 31806136,
					Strand: "-1", Feat: "ENSG00000049319",
					Full: []string{"2", "31800689", "31806136", "-1", "SRD5A2", "ENSG00000049319"},
				},
				{
					Chr: "10", Start: 124782049, Stop: 124817827,
					Strand: "1", Feat: "ENSG00000196177",
					Full: []string{"10", "124782049", "124817827", "1", "ACADSB", "ENSG00000196177"},
				},
				{
					Chr: "1", Start: 860259, Stop: 879955,
					Strand: "1", Feat: "ENSG00000187634",
					Full: []string{"1", "860259", "879955", "1", "SAMD11", "ENSG00000187634"},
				},
			},
			expectedLines: []bed.Line{
				{
					Chr: "1", Start: 860259, Stop: 879955,
					Strand: "1", Feat: "ENSG00000187634",
					Full: []string{"1", "860259", "879955", "1", "SAMD11", "ENSG00000187634"},
				},
				{
					Chr: "1", Start: 948802, Stop: 949920,
					Strand: "1", Feat: "ENSG00000187608",
					Full: []string{"1", "948802", "949920", "1", "ISG15", "ENSG00000187608"},
				},
				{
					Chr: "10", Start: 124768494, Stop: 124773587,
					Strand: "1", Feat: "ENSG00000196177",
					Full: []string{"10", "124768494", "124773587", "1", "ACADSB", "ENSG00000196177"},
				},
				{
					Chr: "10", Start: 124782049, Stop: 124817827,
					Strand: "1", Feat: "ENSG00000196177",
					Full: []string{"10", "124782049", "124817827", "1", "ACADSB", "ENSG00000196177"},
				},
				{
					Chr: "10", Start: 126085871, Stop: 126107545,
					Strand: "-1", Feat: "ENSG00000065154",
					Full: []string{"10", "126085871", "126107545", "-1", "OAT", "ENSG00000065154"},
				},
				{
					Chr: "2", Start: 1800689, Stop: 31806136,
					Strand: "-1", Feat: "ENSG00000049319",
					Full: []string{"2", "31800689", "31806136", "-1", "SRD5A2", "ENSG00000049319"},
				},
				{
					Chr: "2", Start: 31747549, Stop: 31763836,
					Strand: "-1", Feat: "ENSG00000049319",
					Full: []string{"2", "31747549", "31763836", "-1", "SRD5A2", "ENSG00000049319"},
				},
				{
					Chr: "X", Start: 135067597, Stop: 135129423,
					Strand: "1", Feat: "ENSG00000198689",
					Full: []string{"X", "135067597", "135129423", "1", "SLC9A6", "ENSG00000198689"},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			receivedLines := LexicographicSort(tc.lines)
			fmt.Printf("%+v", receivedLines)
			if diff := deep.Equal(tc.expectedLines, receivedLines); diff != nil {
				t.Error("expected VS received lines", diff)
			}
		})
	}
}
