package day03

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var mulRe = regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)`)

func parseInstructionsPart1(r io.Reader) ([]string, error) {
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
	ins, err := parseInstructionsPart1(r)
	if err != nil {
		return 0, fmt.Errorf("parse instructions: %w", err)
	}

	var product int
	for _, in := range ins {
		vals, err := parseMulExpression(in)
		if err != nil {
			return 0, fmt.Errorf("parse expression %q: %w", in, err)
		}
		product += vals[0] * vals[1]
	}
	return product, nil
}

func parseInstructionsPart2(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read all: %w", err)
	}

	re := regexp.MustCompile(`((mul\(\d{1,3},\d{1,3}\))|don't\(\)|do\(\))`)
	matches := re.FindAllString(string(data), -1)
	if matches == nil {
		return nil, fmt.Errorf("no matches found for %q", re.String())
	}
	return matches, nil
}

func RunInstructionsPart2(r io.Reader) (int, error) {
	ins, err := parseInstructionsPart2(r)
	if err != nil {
		return 0, fmt.Errorf("parse instructions: %w", err)
	}

	var product int
	var skipExpression bool
	for _, in := range ins {
		switch in {
		case "do()":
			skipExpression = false
		case "don't()":
			skipExpression = true
		default:
			if skipExpression {
				continue
			}
			operands, err := parseMulExpression(in)
			if err != nil {
				return 0, fmt.Errorf("parse expression %q: %w", in, err)
			}
			product += operands[0] * operands[1]
		}
	}

	return product, nil
}

func parseMulExpression(input string) ([]int, error) {
	res := mulRe.FindAllStringSubmatch(input, -1)
	if len(res) == 0 {
		return nil, fmt.Errorf("no matches found for %q", mulRe.String())
	}
	matches := res[0]
	if len(matches) != 4 {
		return nil, fmt.Errorf("expected 4 matches for %q, got: %+v", mulRe.String(), matches)
	}

	vals := make([]int, 0, len(matches[2:]))
	for _, v := range matches[2:] {
		x, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("parse int %q: %w", v, err)
		}
		vals = append(vals, x)
	}
	return vals, nil
}
