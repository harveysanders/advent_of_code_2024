package day09_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/harveysanders/advent_of_code_2024/day09"
	"github.com/stretchr/testify/require"
)

func TestDay09(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		toTest   func(r io.Reader) (int, error)
		test     func(t *testing.T, got int)
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`12345`))
			},
			toTest: day09.Part1,
			test: func(t *testing.T, got int) {
				require.Equal(t, 60, got)
			},
		},
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`2333133121414131402`))
			},
			toTest: day09.Part1,
			test: func(t *testing.T, got int) {
				require.Equal(t, 1928, got)
			},
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			toTest: day09.Part1,
			test: func(t *testing.T, got int) {
				require.Greater(t, got, 89813169309)
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

func TestDiskString(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "2333133121414131402",
			want:  "00...111...2...333.44.5555.6666.777.888899",
		},
		{
			input: "12345",
			want:  "0..111....22222",
		},
		{
			input: "90909",
			want:  "000000000111111111222222222",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			disk := day09.Disk{}
			_, err := disk.ReadFrom(strings.NewReader(tc.input))
			require.NoError(t, err)
			got := disk.String()
			require.Equal(t, tc.want, got)
		})
	}
}
func TestDiskCompact(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "2333133121414131402",
			want:  "0099811188827773336446555566..............",
		},
		{
			input: "12345",
			want:  "022111222......",
		},
		{
			input: "90909",
			want:  "000000000111111111222222222",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			disk := day09.Disk{}
			_, err := disk.ReadFrom(strings.NewReader(tc.input))
			require.NoError(t, err)
			got := disk.Compact()
			require.Equal(t, tc.want, got)
		})
	}
}
