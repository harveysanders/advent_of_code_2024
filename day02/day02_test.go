package day02_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/harveysanders/advent_of_code_2024/day02"
	"github.com/stretchr/testify/require"
)

func TestDay02(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		want     int
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`))
			},
			want: 2,
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			want: 224,
		},
	}

	for _, tc := range testCases {
		input := tc.getInput(t)
		defer input.Close()

		got, err := day02.CalcSafeReports(input)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
}

func TestIsSafeReport(t *testing.T) {
	testCases := []struct {
		desc            string
		report          []int
		dampenerEnabled bool
		want            bool
	}{
		{
			desc:   "7 6 4 2 1",
			report: []int{7, 6, 4, 2, 1},
			want:   true,
		},
		{
			desc:   "1 2 7 8 9",
			report: []int{1, 2, 7, 8, 9},
			want:   false,
		},
		{
			desc:   "9 7 6 2 1",
			report: []int{9, 7, 6, 2, 1},
			want:   false,
		},
		{
			desc:   "1 3 2 4 5",
			report: []int{1, 3, 2, 4, 5},
			want:   false,
		},
		{
			desc:   "8 6 4 4 1",
			report: []int{8, 6, 4, 4, 1},
			want:   false,
		},
		{
			desc:   "1 3 6 7 9",
			report: []int{1, 3, 6, 7, 9},
			want:   true,
		},
		// With dampener
		{
			desc:            "with dampener - 7 6 4 2 1",
			report:          []int{7, 6, 4, 2, 1},
			dampenerEnabled: true,
			want:            true,
		},
		{
			desc:            "with dampener - 1 2 7 8 9",
			report:          []int{1, 2, 7, 8, 9},
			dampenerEnabled: true,
			want:            false,
		},
		{
			desc:            "with dampener - 9 7 6 2 1",
			report:          []int{9, 7, 6, 2, 1},
			dampenerEnabled: true,
			want:            false,
		},
		{
			desc:            "with dampener - 1 3 2 4 5",
			report:          []int{1, 3, 2, 4, 5},
			dampenerEnabled: true,
			want:            true,
		},
		{
			desc:            "with dampener - 8 6 4 4 1",
			report:          []int{8, 6, 4, 4, 1},
			dampenerEnabled: true,
			want:            true,
		},
		{
			desc:            "with dampener - 1 3 6 7 9",
			report:          []int{1, 3, 6, 7, 9},
			dampenerEnabled: true,
			want:            true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := day02.IsSafeReport(tc.report)
			require.Equal(t, tc.want, got)
		})
	}
}
