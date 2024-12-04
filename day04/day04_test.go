package day04_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/harveysanders/advent_of_code_2024/day04"
	"github.com/stretchr/testify/require"
)

func TestDay04(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		test     func(t *testing.T, got int)
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`))
			},
			test: func(t *testing.T, got int) {
				require.Equal(t, 18, got)
			},
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			test: func(t *testing.T, got int) {
				require.Equal(t, 2358, got)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input := tc.getInput(t)
			defer input.Close()

			count, err := day04.CountXmas(input)
			require.NoError(t, err)
			tc.test(t, count)
		})
	}
}

func TestHelpers(t *testing.T) {
	testCases := []struct {
		desc      string
		countFunc func([]string) int
		want      int
	}{
		{
			desc:      "horizontally",
			countFunc: day04.CountHorizontally,
			want:      5,
		},
		{
			desc:      "vertically",
			countFunc: day04.CountVertically,
			want:      3,
		},
		{
			desc:      "diagonally (/)",
			countFunc: day04.CountForwardDiag,
			want:      5,
		},
		{
			desc:      "diagonally (\\)",
			countFunc: day04.CountBackwardDiag,
			want:      5,
		},
	}

	data := `
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
	`
	rows := strings.Split(string(data), "\n")
	rows = rows[1 : len(rows)-1] // Remove empty newline

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.countFunc(rows)
			require.Equal(t, tc.want, got)
		})
	}
}
