package day07

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type equation struct {
	result   int
	operands []int
}

func ParseEquations(r io.Reader) ([]equation, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("readAll: %w", err)
	}
	data = data[:len(data)-1]
	lines := strings.Split(string(data), "\n")
	eqs := make([]equation, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("parse int %q: %w", parts[0], err)
		}

		rawOps := strings.Fields(parts[1])
		operands := make([]int, len(rawOps))
		for i, v := range rawOps {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("parse int %q: %w", v, err)
			}
			operands[i] = n
		}

		eq := equation{
			result:   result,
			operands: operands,
		}
		eqs = append(eqs, eq)
	}
	return eqs, nil
}

func IsValid(eq equation) bool {
	return false
}
