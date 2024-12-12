package day06

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	markSpace        = "."
	markObstacle     = "#"
	markVisited      = "X"
	markLoopObstacle = "O"
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
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("readAll: %w", err)
	}

	a.grid = bytes.Split(bytes.TrimSuffix(data, []byte("\n")), []byte("\n"))
	for y, row := range a.grid {
		if x := bytes.IndexAny(row, "^<>v"); x > -1 {
			a.guard.pos.X = x
			a.guard.pos.Y = y
			a.guard.dir = dirUp
			a.guard.app = a
			return int64(len(data)), nil
		}
	}
	return int64(len(data)), nil
}

func (a *App) countSquares(char string) int {
	var count int
	for _, row := range a.grid {
		count += bytes.Count(row, []byte(char))
	}
	return count
}

func (m *App) MoveGuard(ctx context.Context) {
	for {
		switch m.guard.dir {
		case dirUp:
			if err := m.guard.MoveUp(ctx); err != nil {
				return
			}

		case dirRight:
			if err := m.guard.MoveRight(ctx); err != nil {
				return
			}
		case dirLeft:
			if err := m.guard.MoveLeft(ctx); err != nil {
				return
			}
		case dirDown:
			if err := m.guard.MoveDown(ctx); err != nil {
				return
			}
		}
	}
}

// moveGuardChan returns a channel that sends a signal if/when the guard movement
// has completed. It's possible that the guard will move in an infinite loop and never complete.
// Use the kill argument to send a signal to stop the potential infinite loop.
func (m *App) moveGuardChan(ctx context.Context) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func(done chan<- struct{}) {
		for {
			switch m.guard.dir {
			case dirUp:
				if err := m.guard.MoveUp(ctx); err != nil {
					done <- struct{}{}
					return
				}

			case dirRight:
				if err := m.guard.MoveRight(ctx); err != nil {
					done <- struct{}{}
					return
				}
			case dirLeft:
				if err := m.guard.MoveLeft(ctx); err != nil {
					done <- struct{}{}
					return
				}
			case dirDown:
				if err := m.guard.MoveDown(ctx); err != nil {
					done <- struct{}{}
					return
				}
			}
		}
	}(c)
	return c
}

// The following "moveX" functions will move the guard in a direction until it
// reaches the end of the grid, or an obstacle.

// MoveUp moves the guard up until an obstacle is hit. If the guard attempts
// to move off the grid, and errOutOfBounds is returned.
func (g *guard) MoveUp(ctx context.Context) error {
	for y := g.pos.Y; y >= -1; y-- {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
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

	}
	// Should never here here
	panic("expected obstacle or out-of-bounds check")
}

// MoveDown moves the guard down until an obstacle is hit. If the guard attempts
// to move off the grid, and errOutOfBounds is returned.
func (g *guard) MoveDown(ctx context.Context) error {
	for y := g.pos.Y; y <= len(g.app.grid); y++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
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
	}
	panic("expected obstacle or out-of-bounds check")
}

// MoveLeft moves the guard left until an obstacle is hit. If the guard attempts
// to move off the grid, and errOutOfBounds is returned.
func (g *guard) MoveLeft(ctx context.Context) error {
	for x := g.pos.X; x >= -1; x-- {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
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
	}
	panic("expected obstacle or out-of-bounds check")
}

// MoveRight moves the guard right until an obstacle is hit. If the guard attempts
// to move off the grid, and errOutOfBounds is returned.
func (g *guard) MoveRight(ctx context.Context) error {
	for x := g.pos.X; x <= len(g.app.grid[0]); x++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
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
	}
	panic("expected obstacle or out-of-bounds check")
}

// placeLoopsObstacles is a very brute-force effort to place an obstacle on every
// possible square. The function then moves the guard to either completion or a
// possible infinite loop. If an infinite loop is detected (via timeout), it's
// assumed we've successfully added a loop-causing obstacle, and we move on to the next square.
func (a *App) placeLoopObstacles() {
	var wg sync.WaitGroup
	// brute force implementation
	var count atomic.Int32
	n := len(a.grid) * len(a.grid[0])
	count.Store(int32(n))
	wg.Add(n)
	fmt.Printf("wait group count: %d\n", n)

	for y, row := range a.grid {
		for x, v := range row {
			go func() {
				defer func() {
					wg.Done()
					count.Add(-1)
					nRemaining := count.Load()
					if nRemaining < 300 {
						fmt.Printf("wait group done. %d remaining\n", nRemaining)
					}
				}()

				if strings.ContainsRune(markObstacle+string(dirUp), rune(v)) {
					return
				}
				clone := a.clone()
				clone.grid[y][x] = markObstacle[0]
				ctx, cancel := context.WithDeadlineCause(context.Background(), time.Now().Add(time.Millisecond*500), fmt.Errorf("infinite loop detected"))
				defer func() {
					cancel()
				}()
				select {
				case <-ctx.Done():
					// probably infinite loop
					// place obstacle
					a.grid[y][x] = markLoopObstacle[0]
					// fmt.Printf("probable infinite loop\n")
					cancel()
					return
				case <-clone.moveGuardChan(ctx):
					// temp obstacle did not create a loop
					return
				}
			}()
		}
	}
	wg.Wait()
	fmt.Println(a.String())
}

func (a *App) clone() *App {
	grid := make([][]byte, 0, len(a.grid))
	for _, row := range a.grid {
		buf := make([]byte, len(row))
		copy(buf, row)
		grid = append(grid, buf)
	}
	app := &App{
		grid: grid,
		guard: guard{
			pos: Coord{X: a.guard.pos.X, Y: a.guard.pos.Y},
			dir: a.guard.dir,
		},
	}
	app.guard.app = app
	return app
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

	app.MoveGuard(context.Background())
	count := app.countSquares(markVisited)

	if app.debug {
		fmt.Println(app.String())
	}

	return count, nil
}

func Part2(r io.Reader) (int, error) {
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

	app.placeLoopObstacles()
	if app.debug {
		fmt.Println(app.String())
	}

	count := app.countSquares(markLoopObstacle)
	fmt.Printf("Obstacles: %d\n", count)
	return count, nil
}
