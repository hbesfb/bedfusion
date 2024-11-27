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
			testing: "correct input, only path",
			bed: Bedfile{
				Path: "/some/path/test.bed",
			},
			expectedBed: Bedfile{
				Path: "/some/path/test.bed",
			},
		},
		{
			testing: "correct input with strand col",
			bed: Bedfile{
				Path:      "/some/path/test.bed",
				StrandCol: 4,
			},
			expectedBed: Bedfile{
				Path:      "/some/path/test.bed",
				StrandCol: 3,
			},
		},
		{
			testing: "correct input with feat col",
			bed: Bedfile{
				Path:    "/some/path/test.bed",
				FeatCol: 3,
			},
			expectedBed: Bedfile{
				Path:    "/some/path/test.bed",
				FeatCol: 2,
			},
		},
		{
			testing: "correct input with both cols",
			bed: Bedfile{
				Path:      "/some/path/test.bed",
				StrandCol: 4,
				FeatCol:   3,
			},
			expectedBed: Bedfile{
				Path:      "/some/path/test.bed",
				StrandCol: 3,
				FeatCol:   2,
			},
		},
		{
			testing: "unclean path",
			bed: Bedfile{
				Path:      "/some/../path/test.bed",
				StrandCol: 4,
				FeatCol:   3,
			},
			expectedBed: Bedfile{
				Path:      "/path/test.bed",
				StrandCol: 3,
				FeatCol:   2,
			},
		},
		{
			testing: "strand col less than 3",
			bed: Bedfile{
				Path:      "/some/path/test.bed",
				StrandCol: 2,
				FeatCol:   3,
			},
			shouldFail: true,
		},
		{
			testing: "feat col less than 3",
			bed: Bedfile{
				Path:      "/some/path/test.bed",
				StrandCol: 4,
				FeatCol:   2,
			},
			shouldFail: true,
		},
		{
			testing: "overlapping strand and feat cols",
			bed: Bedfile{
				Path:      "/some/path/test.bed",
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
