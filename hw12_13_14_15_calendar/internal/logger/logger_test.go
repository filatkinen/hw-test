package logger

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	msg   string
	level LoggingLevel
	want  string
}{
	{msg: "test 1", level: LevelInfo, want: "INFO: test 1 log_level=DEBUG"},
	{msg: "test 2", level: LevelDebug, want: "DEBUG: test 2 log_level=DEBUG"},
	{msg: "test 3", level: LevelWarn, want: "WARN: test 3 log_level=DEBUG"},
}

func TestLogger(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	defer w.Close()
	os.Stdout = w

	logger := New(LevelDebug, os.Stdout)
	defer logger.Close()
	rb := bufio.NewScanner(r)
	for _, test := range tests {
		switch test.level { //nolint
		case LevelInfo:
			logger.Info(test.msg)
		case LevelDebug:
			logger.Debug(test.msg)
		case LevelWarn:
			logger.Warn(test.msg)
		}
		rb.Scan()
		require.Equal(t, strings.SplitN(rb.Text(), " ", 3)[2], test.want)
	}

	os.Stdout = rescueStdout
}

func TestLoggerFile(t *testing.T) {
	f, err := os.CreateTemp("", "testlogger")
	if err != nil {
		log.Fatalf("error creating file %v", err)
	}
	logger := New(LevelDebug, f)

	for _, test := range tests {
		switch test.level { //nolint
		case LevelInfo:
			logger.Info(test.msg)
		case LevelDebug:
			logger.Debug(test.msg)
		case LevelWarn:
			logger.Warn(test.msg)
		}
	}
	filename := f.Name()
	logger.Close()

	fCheck, err1 := os.OpenFile(filename, os.O_RDONLY, 0o644)
	if err1 != nil {
		log.Fatalf("error opening file %v", err)
	}
	defer fCheck.Close()
	scanner := bufio.NewScanner(fCheck)
	counter := 0
	for scanner.Scan() {
		require.Equal(t, strings.SplitN(scanner.Text(), " ", 3)[2], tests[counter].want)
		counter++
	}
}
