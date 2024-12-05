package day05_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/harveysanders/advent_of_code_2024/day05"
	"github.com/stretchr/testify/require"
)

func TestDay05(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func(t *testing.T) io.ReadCloser
		toTest   func(r io.Reader) (int, error)
		test     func(t *testing.T, got int)
	}{
		{
			desc: "part1 - example input",
			getInput: func(t *testing.T) io.ReadCloser {
				return io.NopCloser(strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`))
			},
			toTest: day05.Part1,
			test: func(t *testing.T, got int) {
				require.Equal(t, 143, got)
			},
		},
		{
			desc: "part1 - actual input",
			getInput: func(t *testing.T) io.ReadCloser {
				f, err := os.Open("./input.txt")
				require.NoError(t, err)
				return f
			},
			toTest: day05.Part1,
			test: func(t *testing.T, got int) {
				require.Equal(t, 5964, got)
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

func TestValidateUpdate(t *testing.T) {
	r := strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	app := &day05.App{}
	_, err := app.ReadFrom(r)
	require.NoError(t, err)
	testCases := []struct {
		desc   string
		update []int
		want   bool
	}{
		{
			desc:   "75,47,61,53,29",
			update: []int{75, 47, 61, 53, 29},
			want:   true,
		},
		{
			desc:   "97,61,53,29,13",
			update: []int{97, 61, 53, 29, 13},
			want:   true,
		},
		{
			desc:   "75,29,13",
			update: []int{75, 29, 13},
			want:   true,
		},
		{
			desc:   "75,97,47,61,53",
			update: []int{75, 97, 47, 61, 53},
			want:   false,
		},
		{
			desc:   "61,13,29",
			update: []int{61, 13, 29},
			want:   false,
		},
		{
			desc:   "97,13,75,29,47",
			update: []int{97, 13, 75, 29, 47},
			want:   false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := app.ValidateUpdate(tc.update)
			require.Equal(t, tc.want, got)
		})
	}
}
