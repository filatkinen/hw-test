package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const tofile = "testdata/out.txt"

func TestCopy(t *testing.T) {
	from := "testdata/input.txt"
	src, err := os.Open(from)
	if err != nil {
		t.Error("Failed to read testing file")
	}
	fi, _ := src.Stat()
	ifsize := fi.Size()
	bufin := make([]byte, ifsize)
	bufout := make([]byte, ifsize)
	src.Close()

	for _, v := range []int64{1, 10, 100, 1000} {
		t.Run(fmt.Sprintf("App call test with offset=%d and limit=%d", v, v), func(t *testing.T) {
			err := Copy(from, tofile, v, v)
			require.Nil(t, err)
			src, err := os.Open(from)
			if err != nil {
				t.Errorf("Failed to open in file %s in test with offset=%d and limit=%d", tofile, v, v)
			}
			defer src.Close()
			dst, err := os.Open(tofile)
			if err != nil {
				t.Errorf("Failed to open out file %s in test with offset=%d and limit=%d", tofile, v, v)
			}
			defer dst.Close()
			src.ReadAt(bufin, v)
			n, _ := dst.Read(bufout)
			for i := 0; i < n; i++ {
				require.Equal(t, bufin[i], bufout[i], "in position ", i)
			}
		})
	}
	_ = os.Remove(tofile)
}

func TestParams(t *testing.T) {
	from := "testdata/input.txt"
	src, err := os.Open(from)
	if err != nil {
		t.Error("Failed to read testing file")
	}
	fi, _ := src.Stat()
	ifsize := fi.Size()
	t.Run("offset больше, чем размер файла - невалидная ситуация", func(t *testing.T) {
		ofs := ifsize + 100
		err := Copy(from, tofile, ofs, 0)
		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})
	_ = os.Remove(tofile)
}

func TestUnsupportedFiles(t *testing.T) {
	t.Run("Входной файл - папка: ожидаем ошибку ErrUnsupportedFile", func(t *testing.T) {
		from := "/tmp"
		err := Copy(from, tofile, 0, 0)
		require.Equal(t, err, ErrUnsupportedFile)
	})
	t.Run("Входной файл - /dev/urandom(размер файла=0): ожидаем ошибку ErrUnsupportedFile", func(t *testing.T) {
		from := "/dev/urandom"
		err := Copy(from, tofile, 0, 0)
		require.Equal(t, err, ErrUnsupportedFile)
	})
	_ = os.Remove(tofile)
}
