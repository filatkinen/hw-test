package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	sstat, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	dstat, err := os.Stat(toPath)
	if err == nil && os.SameFile(sstat, dstat) {
		return fmt.Errorf("files in parameters fromPath=%s and toPath=%s are equal", fromPath, toPath)
	}

	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer src.Close()

	fi, err := src.Stat()
	if err != nil {
		return err
	}
	srcsize := fi.Size()
	if srcsize == 0 || fi.IsDir() {
		return ErrUnsupportedFile
	}
	if offset >= srcsize {
		return ErrOffsetExceedsFileSize
	}

	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if offset > 0 {
		_, err = src.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	var rd io.Reader
	if limit > 0 {
		rd = io.LimitReader(src, limit)
	} else {
		rd = io.Reader(src)
	}
	barsize := srcsize - offset
	if limit > 0 && srcsize-offset > limit {
		barsize = limit
	}
	bar := NewBar(barsize)
	rdbar := bar.NewBarProxyReader(rd)

	wt := bufio.NewWriter(dst)

	_, err = io.Copy(wt, rdbar)
	if err != nil {
		return err
	}

	if err := src.Close(); err != nil {
		return err
	}
	return dst.Close()
}
