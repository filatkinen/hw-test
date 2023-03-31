package main

import (
	"errors"
	"io"
	"os"
	"time"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
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

	buf := make([]byte, 1<<15)

	var barsize int64
	switch {
	case offset+limit > srcsize:
		barsize = srcsize - offset
	case limit > 0:
		barsize = limit
	default:
		barsize = srcsize - offset
	}
	bar := NewBar(barsize)

	ticker := time.NewTicker(30 * time.Millisecond)

	rbytes := int64(0)
	wbytes := int64(0)
	for {
		n, err := src.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		rbytes += int64(n)
		if limit != 0 && rbytes >= limit {
			n, err = dst.Write(buf[:n-int(rbytes-limit)])
		} else {
			n, err = dst.Write(buf[:n])
		}
		if err != nil {
			return err
		}
		wbytes += int64(n)
		select {
		case <-ticker.C:
			bar.ShowProgress(wbytes)
		default:
		}
	}
	bar.ShowProgress(wbytes)
	bar.Finish()

	if err := src.Close(); err != nil {
		return err
	}
	return dst.Close()
}
