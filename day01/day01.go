package day01

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

// CalcDifferenceScore sorts each vertical list of ints and calculates the difference
// between the values on each line
func CalcDifferenceScore(r io.Reader) (int, error) {
	lists, err := ParseLists(r)
	if err != nil {
		return 0, fmt.Errorf("parse lists: %w", err)
	}
	for _, list := range lists {
		slices.Sort(list)
	}

	sum := float64(0)
	for _, list := range lists {
		for i, v := range list {
			diff := math.Abs(float64(v) - float64(lists[1][i]))
			sum += diff
		}
	}

	return int(sum), nil
}

func CalcSimilarityScore(r io.Reader) (int, error) {
	lists, err := ParseLists(r)
	if err != nil {
		return 0, fmt.Errorf("parse lists: %w", err)
	}
	similarityScore := 0

	for _, item := range lists[0] {
		appearances := 0
		for _, v := range lists[1] {
			if item == v {
				appearances += 1
			}
		}
		similarityScore += item * appearances
	}
	return similarityScore, nil
}

func ParseLists(r io.Reader) ([][]int, error) {
	scr := bufio.NewScanner(r)
	lists := make([][]int, 2)
	for i := range lists {
		lists[i] = []int{}
	}

	for scr.Scan() {
		line := scr.Text()
		nums := strings.Fields(line)
		for i, rawN := range nums {
			n, err := strconv.Atoi(rawN)
			if err != nil {
				return nil, fmt.Errorf("parse int %q: %w", rawN, err)
			}
			lists[i] = append(lists[i], n)
		}
	}
	if scr.Err() != nil {
		return nil, scr.Err()
	}
	return lists, nil
}
