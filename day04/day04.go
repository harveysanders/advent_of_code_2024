package day04

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

const (
	xmas  = "XMAS"
	samx  = "SAMX"
	x_mas = "MAS"
	x_sam = "SAM"
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

func CountX_mas(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("readAll: %w", err)
	}

	var count int
	rows := strings.Split(string(data), "\n")
	rows = rows[:len(rows)-1] // Remove empty newline

	// Find a diagonal (↗) "MAS"
	bottomRightX := len(rows[0]) - 1
	bottomRightY := len(rows) - 1

	masRe := regexp.MustCompile(x_mas)
	samRe := regexp.MustCompile(x_sam)

out:
	for i := 0; ; i++ {
		var diag strings.Builder
		for x := 0; x < len(rows[0]); x++ {
			y := i - x
			if x == bottomRightX && y == bottomRightY {
				break out
			}
			if y < 0 || y >= len(rows) {
				continue
			}
			diag.WriteByte(rows[y][x])
		}
		matches := masRe.FindAllStringIndex(diag.String(), -1)
		matchesSam := samRe.FindAllStringIndex(diag.String(), -1)
		if matchesSam != nil {
			matches = append(matches, matchesSam...)
		}
		if len(matches) == 0 {
			continue
		}

		for _, match := range matches {
			startOriginX := match[0]

			// Adjust x coordinate when "scan line" goes beyond the grid
			if i > len(rows)-1 {
				startOriginX = i - (len(rows) - 1) + match[0]
			}
			startOriginY := i - startOriginX

			// fmt.Printf("[%d] -> %q @ {%d, %d}\n",
			// 	i,
			// 	string(rows[startOriginY][startOriginX]),
			// 	startOriginX,
			// 	startOriginY,
			// )
			// fmt.Println("______")

			// Check if other diag ↘ has "MAS"
			// Grab the 3-char string from using the startOrigin {x,y}
			diagDown := make([]byte, 0, len(x_mas))
			for i := 0; i < len(x_mas); i++ {
				x := startOriginX + i
				y := startOriginY - (len(x_mas) - 1) + i
				diagDown = append(diagDown, rows[y][x])
			}
			if countStrings(string(diagDown), x_mas, x_sam) > 0 {
				count += 1
			}
		}
	}
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

// CountForwardDiag starts from the top and scans the rows diagonally (↗).
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

// CountBackwardDiag starts from bottom and scans the rows diagonally (↘).
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
