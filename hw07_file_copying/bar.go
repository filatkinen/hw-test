package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Bar ...
type Bar struct {
	percent int64
	cur     int64
	total   int64
	rate    string
	prefix  string
	shift   int
}

type Reader struct {
	io.Reader
	bar *Bar
}

func NewBar(size int64) *Bar {
	bar := Bar{cur: 0, total: size}
	switch {
	case size < 1<<10:
		bar.prefix = "Byte"
		bar.shift = 0
	case size < 1<<20:
		bar.prefix = "KByte"
		bar.shift = 10
	case size < 1<<30:
		bar.prefix = "MByte"
		bar.shift = 20
	case size < 1<<40:
		bar.prefix = "Gbyte"
		bar.shift = 30
	}
	return &bar
}

func (bar *Bar) NewBarProxyReader(r io.Reader) *Reader {
	return &Reader{r, bar}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.bar.cur += int64(n)
	r.bar.ShowProgress(r.bar.cur)
	if errors.Is(err, io.EOF) {
		r.bar.Finish()
	}
	return
}

func (bar *Bar) ShowProgress(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = int64((float32(bar.cur) / float32(bar.total)) * 100)
	if bar.percent != last {
		bar.rate = strings.Repeat("=", int(bar.percent)/2)
	}
	curshift := int(bar.cur >> bar.shift)
	totalshift := int(bar.total >> bar.shift)
	fmt.Printf("\r[%-50s]%3d%% %8d/%d(%s)", bar.rate, bar.percent, curshift, totalshift, bar.prefix)
}

func (bar *Bar) Finish() {
	fmt.Println()
}
