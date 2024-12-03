package day03_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/harveysanders/advent_of_code_2024/day03"
	"github.com/stretchr/testify/require"
)

func TestRunInstructions(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		want     int
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`))
			},
			want: 161,
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			want: 174336360,
		},
	}

	for _, tc := range testCases {
		input := tc.getInput(t)
		defer input.Close()

		got, err := day03.RunInstructions(input)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
}
