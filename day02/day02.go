package day02

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func ParseReports(r io.Reader) ([][]int, error) {
	scr := bufio.NewScanner(r)
	reports := [][]int{}

	for scr.Scan() {
		line := scr.Text()
		rawReport := strings.Fields(line)
		report := make([]int, 0, len(rawReport))
		for _, v := range rawReport {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("parse int %q: %w", v, err)
			}
			report = append(report, n)
		}
		reports = append(reports, report)
	}
	if scr.Err() != nil {
		return nil, fmt.Errorf("scan line: %w", scr.Err())
	}

	return reports, nil
}

func CalcSafeReports(r io.Reader, useDampener bool) (int, error) {
	reports, err := ParseReports(r)
	if err != nil {
		return 0, fmt.Errorf("parse reports: %w", err)
	}
	var safeCount int
	for _, report := range reports {
		if IsSafeReport(report, useDampener) {
			safeCount += 1
		}
	}
	return safeCount, nil
}

type direction int

const (
	dirUnknown direction = iota
	dirAsc
	dirDesc
)

func isSafeReport(report []int) bool {
	if len(report) < 2 {
		return false
	}
	var dir direction
	lastVal := report[0]
	for _, v := range report[1:] {
		diff := lastVal - v
		lastVal = v
		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}
		newDir := dirAsc
		// If diff is positive, the direction is descending
		if diff > 0 {
			newDir = dirDesc
		}
		switch dir {
		case dirUnknown:
			// set dir
			dir = newDir
		default:
			// Direction is already set.
			// Ensure the new direction is the
			// same as before. If not, bail.
			if dir != newDir {
				return false
			}
		}
	}
	return true
}

func IsSafeReport(report []int, dampener bool) bool {
	if !dampener {
		return isSafeReport(report)
	}

	var wg sync.WaitGroup
	wg.Add(len(report))
	vals := make([]bool, len(report))

	// TODO: Fix lazy brute force report checking
	for i := range report {
		dampened := slices.Delete(slices.Clone(report), i, i+1)
		go func(s []int) {
			vals[i] = isSafeReport(s)
			wg.Done()
		}(dampened)
	}

	wg.Wait()
	return slices.Contains(vals, true)
}
