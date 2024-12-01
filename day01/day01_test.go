package day01_test

import (
	"io"
	"os"
	"strings"
	"testing"

	day01 "github.com/harveysanders/advent_of_code_2024/day01"
	"github.com/stretchr/testify/require"
)

func TestDay01_part01(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		want     int
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3
`))
			},
			want: 11,
		},
		{
			desc: "part1 - real input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			want: 3246517,
		},
	}

	for _, tc := range testCases {
		input := tc.getInput(t)
		defer input.Close()

		got, err := day01.CalcDifferenceScore(input)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
}

func TestDay01_part02(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		want     int
	}{
		{
			desc: "part2 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3
`))
			},
			want: 31,
		},
		{
			desc: "part2 - real input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			want: 29379307,
		},
	}

	for _, tc := range testCases {
		input := tc.getInput(t)
		defer input.Close()

		got, err := day01.CalcSimilarityScore(input)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
}

func TestParseLists(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3
`

	lists, err := day01.ParseLists(strings.NewReader(input))
	require.NoError(t, err)

	require.Len(t, lists, 2)
	require.Len(t, lists[0], 6)
}
