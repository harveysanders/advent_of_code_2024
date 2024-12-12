package day08

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

//	{
//	   freqs: {
//	     "A": {
//	         antennas: []Position{{ X: 4, Y: 7 }},
//	         antinodes: []Position{}
//	      }
//	   }
//
// }
type App struct {
	freqs map[string]frequency
	grid  []string
}

type Position struct {
	X int
	Y int
}

type frequency struct {
	antinodes map[string]Position
	antennas  []Position
}

func (a *App) ReadFrom(r io.Reader) (int64, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(data), "\n")
	a.grid = lines[:len(lines)-1]
	a.freqs = make(map[string]frequency)

	freqRe := regexp.MustCompile(`\w{1}`)
	for y, line := range a.grid {
		matches := freqRe.FindAllStringIndex(line, -1)
		for _, match := range matches {
			if len(match) == 0 {
				continue
			}

			x := match[0]
			fLabel := string(line[x])
			pos := Position{X: x, Y: y}
			f, ok := a.freqs[fLabel]
			if !ok {
				a.freqs[fLabel] = frequency{
					antinodes: make(map[string]Position),
					antennas:  []Position{pos},
				}
				continue
			}
			f.antennas = append(f.antennas, pos)
			a.freqs[fLabel] = f
		}
	}
	return int64(len(data)), nil
}

func Part1(r io.Reader) (int, error) {
	app := App{}
	_, err := app.ReadFrom(r)
	if err != nil {
		return 0, fmt.Errorf("readFrom: %w", err)
	}
	app.findAntinodes()
	return 0, nil
}

func (a *App) findAntinodes() {
	for label, freq := range a.freqs {
		if len(freq.antennas) < 2 {
			continue
		}
		for _, antAPos := range freq.antennas {
			for _, antBPos := range freq.antennas {
				if antAPos.Equal(antBPos) {
					continue
				}
				antinode1 := antAPos.oppositePosition(antBPos)
				if a.inBounds(antinode1) {
					freq.antinodes[antinode1.String()] = antinode1
				}
				antinode2 := antBPos.oppositePosition(antAPos)
				if a.inBounds(antinode2) {
					freq.antinodes[antinode2.String()] = antinode2
				}

				fmt.Printf("%q - A,        %+v\n", label, antAPos)
				fmt.Printf("%q - B,        %+v\n", label, antBPos)
				fmt.Printf("%q - antinode, %+v\n", label, antinode1)
				fmt.Printf("%q - antinode, %+v\n", label, antinode2)
			}
		}
		a.freqs[label] = freq
	}
}

func (app *App) inBounds(p Position) bool {
	if p.X < 0 || p.X > len(app.grid[0])-1 {
		return false
	}
	if p.Y < 0 || p.Y > len(app.grid)-1 {
		return false
	}
	return true
}

func (p Position) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (pivot Position) oppositePosition(b Position) Position {
	return Position{
		X: (2 * pivot.X) - b.X,
		Y: (2 * pivot.Y) - b.Y,
	}
}

func (p Position) Equal(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}
