package day06_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/harveysanders/advent_of_code_2024/day06"
	"github.com/stretchr/testify/require"
)

func TestDay06(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		toTest   func(r io.Reader) (int, error)
		test     func(t *testing.T, got int)
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`))
			},
			toTest: day06.Part1,
			test: func(t *testing.T, got int) {
				require.Equal(t, 41, got)
			},
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			toTest: day06.Part1,
			test: func(t *testing.T, got int) {
				require.Equal(t, 4982, got)
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
