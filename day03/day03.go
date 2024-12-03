package day03

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func parseInstructions(r io.Reader) ([]string, error) {
	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read all: %w", err)
	}

	matches := re.FindAllString(string(data), -1)
	if matches == nil {
		return nil, fmt.Errorf("no matches found for %q", re.String())
	}
	return matches, nil
}

func RunInstructions(r io.Reader) (int, error) {
	ins, err := parseInstructions(r)
	if err != nil {
		return 0, fmt.Errorf("parse instructions: %w", err)
	}

	mulRe := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)`)
	var product int
	for _, in := range ins {
		res := mulRe.FindAllStringSubmatch(in, -1)
		if len(res) == 0 {
			return 0, fmt.Errorf("no matches found for %q", mulRe.String())
		}
		matches := res[0]
		if len(matches) != 4 {
			return 0, fmt.Errorf("expected 4 matches for %q, got: %+v", mulRe.String(), matches)
		}
		vals := make([]int, 0, 2)
		for _, v := range matches[2:] {
			x, err := strconv.Atoi(v)
			if err != nil {
				return 0, fmt.Errorf("parse int %q: %w", v, err)
			}
			vals = append(vals, x)
		}

		product += vals[0] * vals[1]
	}
	return product, nil
}
