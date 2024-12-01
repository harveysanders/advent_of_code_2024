package main

import (
	"io"
	"os"
	"strings"
	"testing"

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
		got, err := part1(tc.getInput(t))
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

	lists, err := parseLists(strings.NewReader(input))
	require.NoError(t, err)

	require.Len(t, lists, 2)
	require.Len(t, lists[0], 6)
}
