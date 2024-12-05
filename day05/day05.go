package day05

import (
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

type App struct {
	rules   map[int][]int // Page numbers to a list of pages that must follow the given page.
	updates [][]int       // Each update contains a list of page numbers to produce the update.
}

func (app *App) ReadFrom(r io.Reader) (int64, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("read all: %w", err)
	}

	x := strings.Split(string(data), "\n\n")
	rawRules := x[0]
	rawUpdates := x[1]

	app.rules = make(map[int][]int, len(rawRules))
	for _, v := range strings.Split(rawRules, "\n") {
		line := strings.Split(v, "|")
		key, _ := strconv.Atoi(line[0])
		page, _ := strconv.Atoi(line[1])
		rule, ok := app.rules[key]
		if !ok {
			app.rules[key] = []int{page}
		} else {
			app.rules[key] = append(rule, page)
		}
	}

	updateLines := strings.Split(rawUpdates, "\n")
	app.updates = make([][]int, 0, len(updateLines))
	for _, line := range updateLines {
		rawPages := strings.Split(line, ",")
		pages := make([]int, 0, len(rawPages))
		for _, p := range rawPages {
			page, _ := strconv.Atoi(p)
			pages = append(pages, page)
		}
		app.updates = append(app.updates, pages)
	}
	return int64(len(data)), nil
}

// validUpdates returns a slice of updates that are already in the correct order
// as determined by the ordering rules.
func (app *App) ValidUpdates() ([][]int, error) {
	res := make([][]int, 0, len(app.updates))
	for _, update := range app.updates {
		if app.ValidateUpdate(update) {
			res = append(res, update)
		}
	}
	return res, nil
}

// ValidateUpdate verifies the pages in the update are in the correct order
// as determined by the ordering rules
func (app *App) ValidateUpdate(update []int) bool {
	for i, page := range update {
		for _, nextPage := range update[i+1:] {
			mustFollow, ok := app.rules[nextPage]
			if !ok {
				// ignore the page if there is no rule
				continue
			}
			// If there is a rule that say the current page
			// should be before the nextPage, fail
			if slices.Contains(mustFollow, page) {
				return false
			}
		}
	}
	return true
}

func SumMiddlePages(updates [][]int) int {
	var sum int
	for _, update := range updates {
		i := math.Floor(float64(len(update)) / 2)
		sum += update[int(i)]
	}
	return sum
}

func Part1(r io.Reader) (int, error) {
	app := &App{}
	_, err := app.ReadFrom(r)
	if err != nil {
		return 0, fmt.Errorf("readFrom: %w", err)
	}

	valid, err := app.ValidUpdates()
	if err != nil {
		return 0, fmt.Errorf("validate updates: %w", err)
	}

	sum := SumMiddlePages(valid)
	return sum, nil
}
