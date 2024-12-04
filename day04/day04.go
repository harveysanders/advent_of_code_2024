package day04

import (
	"fmt"
	"io"
	"strings"
)

const (
	xmas = "XMAS"
	samx = "SAMX"
)

func CountXmas(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("readAll: %w", err)
	}

	rows := strings.Split(string(data), "\n")
	rows = rows[:len(rows)-1] // Remove empty newline

	count := CountHorizontally(rows)
	count += CountVertically(rows)
	count += CountForwardDiag(rows)
	count += CountBackwardDiag(rows)

	return count, nil
}

func countStrings(line string, str ...string) int {
	var count int
	for _, v := range str {
		n := strings.Count(line, v)
		count += n
	}
	return count
}

func CountHorizontally(rows []string) int {
	var count int
	for _, row := range rows {
		count += countStrings(row, xmas, samx)
	}
	return count
}

func CountVertically(rows []string) int {
	var count int
	for x := range rows[0] {
		var vertLine strings.Builder
		for y := range rows {
			vertLine.WriteByte(rows[y][x])
		}
		count += countStrings(vertLine.String(), xmas, samx)
	}
	return count
}

// CountForwardDiag scans the rows diagonally (/).
func CountForwardDiag(rows []string) int {
	bottomRightX := len(rows[0]) - 1
	bottomRightY := len(rows) - 1
	var count int

	for i := 0; ; i++ {
		var diag strings.Builder
		for x := 0; x < len(rows[0]); x++ {
			y := i - x
			if x == bottomRightX && y == bottomRightY {
				return count
			}
			if y < 0 || y >= len(rows) {
				continue
			}
			diag.WriteByte(rows[y][x])
		}
		count += countStrings(diag.String(), xmas, samx)
	}
}

// CountBackwardDiag scans the rows diagonally (\).
func CountBackwardDiag(rows []string) int {
	topRightX := len(rows[0]) - 1
	topRightY := 0
	var count int

	for i := len(rows); ; i-- {
		var diag strings.Builder
		for x := range rows[0] {
			y := i + x
			if x == topRightX && y == topRightY {
				return count
			}
			if y < 0 || y >= len(rows) {
				continue
			}
			diag.WriteByte(rows[y][x])
		}
		count += countStrings(diag.String(), xmas, samx)
	}
}
