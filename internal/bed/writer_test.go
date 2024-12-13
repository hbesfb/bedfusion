package bed

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/spf13/afero"
)

func TestWrite(t *testing.T) {
	type testCase struct {
		testing             string
		bed                 Bedfile
		expectedFileContent string
		shouldFail          bool
	}
	testCases := []testCase{
		{
			testing: "bed file with headers",
			bed: Bedfile{
				Output: "/a/test/folder/output.bed",
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
			expectedFileContent: "browser something\n" +
				"track something\n" +
				"#something\n" +
				"1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
		},
	}
	// Setting up virtual file system
	appFS := afero.NewMemMapFs()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			// Create writer and write to file
			writer, err := appFS.Create(tc.bed.Output)
			if err != nil {
				t.Fatal(err)
			}
			err = tc.bed.write(writer)
			if (!tc.shouldFail && err != nil) || (tc.shouldFail && err == nil) {
				t.Fatalf("shouldFail is %t, but err is %q", tc.shouldFail, err)
			}
			if !tc.shouldFail {
				// Read the file we wrote to
				fileContent, err := afero.ReadFile(appFS, tc.bed.Output)
				if err != nil {
					t.Fatal(err)
				}
				// Check if the content is as expected
				if tc.expectedFileContent != string(fileContent) {
					t.Error("expectedFileContent vs fileContent:\n",
						tc.expectedFileContent, "\n!=\n", string(fileContent))
					t.Log("expectedFileContent: ", []byte(tc.expectedFileContent))
					t.Log("fileContent:     ", fileContent)
				}
			}
		})
	}
}

func TestToString(t *testing.T) {
	type testCase struct {
		testing        string
		bed            Bedfile
		expectedString string
	}
	testCases := []testCase{
		{
			testing: "simple bed file",
			bed: Bedfile{
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
			expectedString: "1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
		},
		{
			testing: "bed file with headers",
			bed: Bedfile{
				Inputs: []string{"test.bed"},
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
			expectedString: "browser something\n" +
				"track something\n" +
				"#something\n" +
				"1\t10\t100\n" +
				"2\t20\t200\n" +
				"3\t30\t300\n" +
				"4\t40\t400\n",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			receivedString := tc.bed.toString()
			if diff := deep.Equal(tc.expectedString, receivedString); diff != nil {
				t.Error("expected VS received bed", diff)
			}
		})
	}
}
