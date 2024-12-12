package day07_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay07(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		toTest   func(r io.Reader) (int, error)
		test     func(t *testing.T, got int)
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`))
			},
			toTest: func(r io.Reader) (int, error) {
				return 0, nil
			},
			test: func(t *testing.T, got int) {
				require.Equal(t, 3749, got)
			},
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
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
