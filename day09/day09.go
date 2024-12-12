package day09

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Disk struct {
	layout []int
}

func (d *Disk) ReadFrom(r io.Reader) (int64, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	data = bytes.TrimSuffix(data, []byte("\n"))
	d.layout = make([]int, 0, len(data))
	parts := strings.Split(string(data), "")
	nRead := int64(len(data))

	for i, v := range parts {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nRead, fmt.Errorf("parse int (%s) @ index [%d]: %w", v, i, err)
		}
		d.layout = append(d.layout, n)
	}
	return nRead, nil
}

func (d *Disk) String() string {
	var out strings.Builder
	var id int
	for i, n := range d.layout {
		if i%2 == 0 {
			// Even indexes are files
			out.WriteString(strings.Repeat(fmt.Sprintf("%d", id), n))
			id += 1
		} else {
			// Odd indexes are free space blocks
			out.WriteString(strings.Repeat(".", n))
		}
	}
	return out.String()
}

func (d *Disk) Compact() string {
	layout := strings.Split(d.String(), "")
	return strings.Join(compact(layout), "")
}

func compact(layout []string) []string {
	endPos := len(layout) - 1
	var done bool
	for i := 0; !done; i++ {
		if i == endPos {
			return layout
		}
		if layout[i] != "." {
			continue
		}

		for {
			v := layout[endPos]
			if v == "." {
				endPos -= 1
				continue
			}
			layout[i] = v
			layout[endPos] = "."
			endPos -= 1
			break
		}
	}
	return layout
}

func Checksum(layout string) (int, error) {
	var result int
	for i, v := range layout {
		if v == '.' {
			continue
		}
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, err
		}
		result += (i * n)
	}
	return result, nil
}

func Part1(r io.Reader) (int, error) {
	disk := Disk{}
	_, err := disk.ReadFrom(r)
	if err != nil {
		return 0, fmt.Errorf("disk.ReadFrom(r): %w", err)
	}
	compacted := disk.Compact()
	return Checksum(compacted)
}
