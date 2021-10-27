package gfwatch

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func (gfw *GfWatch) Decode(r io.Reader) error {

	var (
		line      string
		lineArray []string
	)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line = scanner.Text()
		lineArray = strings.Split(line, "|")
		if len(lineArray) != 3 {
			continue
		}
		gfw.gfMap[lineArray[0]] = struct{}{}
	}

	return nil
}

func FromFile(path string) (*GfWatch, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gfw := New()
	if err := gfw.Decode(f); err != nil {
		return nil, err
	}
	return gfw, nil
}