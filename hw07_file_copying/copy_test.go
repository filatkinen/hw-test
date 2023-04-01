package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const tofile = "testdata/out.txt"

func TestCopy(t *testing.T) {
	from := "testdata/input.txt"
	src, err := os.Open(from)
	require.NoError(t, err)
	fi, _ := src.Stat()
	ifsize := fi.Size()
	bufin := make([]byte, ifsize)
	bufout := make([]byte, ifsize)
	src.Close()

	for _, v := range []int64{0, 1, 10, 100, 1000} {
		t.Run(fmt.Sprintf("App call test with offset=%d and limit=%d", v, v), func(t *testing.T) {
			err := Copy(from, tofile, v, v)
			require.NoError(t, err)

			src, err := os.Open(from)
			require.NoError(t, err)
			defer src.Close()

			dst, err := os.Open(tofile)
			require.NoError(t, err)
			defer dst.Close()

			if v > 0 {
				_, err = src.Seek(v, io.SeekStart)
				if err != nil {
					t.Errorf(err.Error())
				}
			}

			var rd io.Reader
			if v > 0 {
				rd = io.LimitReader(src, v)
			} else {
				rd = io.Reader(src)
			}
			nread, _ := rd.Read(bufin)
			nwrite, _ := dst.Read(bufout)
			require.Equal(t, nread, nwrite)
			for i := 0; i < nwrite; i++ {
				require.Equal(t, bufin[i], bufout[i], "in position ", i)
			}
		})
	}
	_ = os.Remove(tofile)
}

func TestParams(t *testing.T) {
	from := "testdata/input.txt"
	src, err := os.Open(from)
	require.NoError(t, err)
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
