package day08_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/harveysanders/advent_of_code_2024/day08"
	"github.com/stretchr/testify/require"
)

func TestDay08(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		toTest   func(r io.Reader) (int, error)
		test     func(t *testing.T, got int)
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`))
			},
			toTest: day08.Part1,
			test: func(t *testing.T, got int) {
				require.Equal(t, 14, got)
			},
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			toTest: day08.Part1,
			test: func(t *testing.T, got int) {
				require.Greater(t, got, 0)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input := tc.getInput(t)
			defer input.Close()

			got, err := tc.toTest(input)
			require.NoError(t, err)

			tc.test(t, got)
		})
	}
}
