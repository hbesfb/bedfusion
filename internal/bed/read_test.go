package bed

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestRead(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing        string
		bed            Bedfile
		bedFileContent string
		expectedBed    Bedfile
		shouldFail     bool
	}
	testCases := []testCase{
		{
			testing: "simple bed file",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
			expectedBed: Bedfile{
				Path: "test.bed",
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
			testing: "simple bed file with header",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "browser something\n" +
				"track something\n" +
				"#something\n" +
				"1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
			expectedBed: Bedfile{
				Path: "test.bed",
				Header: []string{
					"browser something",
					"track something",
					"#something",
				},
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
				Path:      "test.bed",
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
			},
			bedFileContent: "1\t10\t100\t-1\tA\n" +
				"2\t20\t200\t-1\tB\n" +
				"3\t30\t300\t1\tC\n" +
				"4\t40\t400\t1\tD\n",
			expectedBed: Bedfile{
				Path:      "test.bed",
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
		{
			testing: "complex bed file with strand and feat and header",
			bed: Bedfile{
				Path:      "test.bed",
				StrandCol: 4 - 1,
				FeatCol:   6 - 1,
			},
			bedFileContent: "#a test header\n" +
				"1\t860259\t879955\t1\tSAMD11\tENSG00000187634\n" +
				"1\t948802\t949920\t1\tISG15\tENSG00000187608\n" +
				"10\t124768494\t124773587\t1\tACADSB\tENSG00000196177\n" +
				"10\t124782049\t124817827\t1\tACADSB\tENSG00000196177\n" +
				"10\t126085871\t126107545\t-1\tOAT\tENSG00000065154\n" +
				"X\t135067597\t135129423\t1\tSLC9A6\tENSG00000198689",
			expectedBed: Bedfile{
				Path:      "test.bed",
				StrandCol: 4 - 1,
				FeatCol:   6 - 1,
				Header:    []string{"#a test header"},
				Lines: []Line{
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
						Chr: "X", Start: 135067597, Stop: 135129423,
						Strand: "1", Feat: "ENSG00000198689",
						Full: []string{"X", "135067597", "135129423", "1", "SLC9A6", "ENSG00000198689"},
					},
				},
			},
		},
		{
			testing: "missing column",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "10\t100\n" +
				"20\t200\n" +
				"30\t300\n" +
				"40\t400\n",
			shouldFail: true,
		},
		{
			testing: "changing column numbers",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\t-1\n" +
				"4\t40\t400\n",
			shouldFail: true,
		},
		{
			testing: "start not a number",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "1\tX\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
			shouldFail: true,
		},
		{
			testing: "stop not a number",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\tCD\n",
			shouldFail: true,
		},
		{
			testing: "unknown header",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "something\n" +
				"1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
			shouldFail: true,
		},
		{
			testing: "multi track file",
			bed: Bedfile{
				Path: "test.bed",
			},
			bedFileContent: "browser something\n" +
				"track something\n" +
				"#something\n" +
				"1\t10\t100\n" +
				"2\t20\t200\n" +
				"track something\n" +
				"#something\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
			shouldFail: true,
		},
		{
			testing: "strand in incorrect format",
			bed: Bedfile{
				Path:      "test.bed",
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
			},
			bedFileContent: "1\t10\t100\t-1\tA\n" +
				"2\t20\t200\t-1\tB\n" +
				"3\t30\t300\t0\tC\n" +
				"4\t40\t400\t1\tD\n",
			shouldFail: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			err := tc.bed.read(strings.NewReader(tc.bedFileContent))
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
