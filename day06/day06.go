package day06

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	markSpace    = "."
	markObstacle = "#"
	markVisited  = "X"
)

var errOutOfBounds = errors.New("out of bounds")

type direction rune

const (
	dirUp    direction = '^'
	dirRight direction = '>'
	dirLeft  direction = '<'
	dirDown  direction = 'v'
)

type Coord struct {
	X int
	Y int
}

type guard struct {
	pos Coord
	dir direction
	app *App
}

type App struct {
	grid  [][]byte
	guard guard
	debug bool
}

func (a *App) ReadFrom(r io.Reader) (int64, error) {
	rows, err := io.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("readAll: %w", err)
	}

	a.grid = bytes.Split(rows, []byte("\n"))
	for y, row := range a.grid {
		if x := bytes.IndexAny(row, "^<>v"); x > -1 {
			a.guard.pos.X = x
			a.guard.pos.Y = y
			a.guard.dir = dirUp
			a.guard.app = a
			return 0, nil
		}
	}
	return 0, nil
}

func (m *App) NVisited() int {
	var count int
	for _, row := range m.grid {
		count += bytes.Count(row, []byte(markVisited))
	}
	return count
}

func (m *App) MoveGuard() {
	for {
		switch m.guard.dir {
		case dirUp:
			if err := m.guard.MoveUp(); err != nil {
				return
			}

		case dirRight:
			if err := m.guard.MoveRight(); err != nil {
				return
			}
		case dirLeft:
			if err := m.guard.MoveLeft(); err != nil {
				return
			}
		case dirDown:
			if err := m.guard.MoveDown(); err != nil {
				return
			}
		}
	}
}

func (g *guard) MoveUp() error {
	for y := g.pos.Y; y >= -1; y-- {
		if y == -1 {
			return errOutOfBounds
		}

		mark := g.app.grid[y][g.pos.X]
		if string(mark) == markObstacle {
			g.dir = dirRight
			return nil
		}
		g.app.grid[y][g.pos.X] = markVisited[0]
		g.pos.Y = y
	}

	// Should never here here
	panic("expected obstacle or out-of-bounds check")
}

func (g *guard) MoveDown() error {
	for y := g.pos.Y; y <= len(g.app.grid); y++ {
		if y == len(g.app.grid) {
			return errOutOfBounds
		}

		mark := g.app.grid[y][g.pos.X]
		if string(mark) == markObstacle {
			g.dir = dirLeft
			return nil
		}
		g.app.grid[y][g.pos.X] = markVisited[0]
		g.pos.Y = y
	}
	panic("expected obstacle or out-of-bounds check")
}

func (g *guard) MoveLeft() error {
	for x := g.pos.X; x >= -1; x-- {
		if x == -1 {
			return errOutOfBounds
		}

		mark := g.app.grid[g.pos.Y][x]
		if string(mark) == markObstacle {
			g.dir = dirUp
			return nil
		}
		g.app.grid[g.pos.Y][x] = markVisited[0]
		g.pos.X = x
	}
	panic("expected obstacle or out-of-bounds check")
}

func (g *guard) MoveRight() error {
	for x := g.pos.X; x <= len(g.app.grid[0]); x++ {
		if x == len(g.app.grid[0]) {
			return errOutOfBounds
		}

		mark := g.app.grid[g.pos.Y][x]
		if string(mark) == markObstacle {
			g.dir = dirDown
			return nil
		}
		g.app.grid[g.pos.Y][x] = markVisited[0]
		g.pos.X = x
	}
	panic("expected obstacle or out-of-bounds check")
}

func (a *App) String() string {
	var s strings.Builder
	for _, row := range a.grid {
		s.Write(row)
		s.WriteString("\n")
	}
	return s.String()
}

func Part1(r io.Reader) (int, error) {
	app := &App{}
	_, err := app.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	if app.debug {
		fmt.Println(app.String())
		fmt.Println("_____________")
		fmt.Println("")
	}

	app.MoveGuard()
	count := app.NVisited()

	if app.debug {
		fmt.Println(app.String())
	}

	return count, nil
}
