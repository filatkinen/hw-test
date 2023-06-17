package logger

import (
	"io"
	"log"
	"sync"
)

type LoggingLevel int

const (
	LevelError LoggingLevel = 1 << iota
	LevelWarn
	LevelInfo
	LevelDebug
	LevelNotDefined
)

func (l LoggingLevel) String() string {
	switch l {
	case LevelError:
		return "ERROR"
	case LevelWarn:
		return "WARN"
	case LevelInfo:
		return "INFO"
	case LevelDebug:
		return "DEBUG"
	case LevelNotDefined:
		return "Not defined"
	default:
		return "Not defined"
	}
}

func GetLoggingLevel(level string) LoggingLevel {
	switch level {
	case "ERROR":
		return LevelError
	case "WARN":
		return LevelWarn
	case "INFO":
		return LevelInfo
	case "DEBUG":
		return LevelDebug
	default:
		return LevelNotDefined
	}
}

type Logger struct {
	mu       sync.Mutex
	logLevel LoggingLevel
	logOut   *log.Logger
	out      io.WriteCloser
}

func New(level LoggingLevel, oint io.WriteCloser) *Logger {
	return &Logger{
		mu:       sync.Mutex{},
		logLevel: level,
		out:      oint,
		logOut:   log.New(oint, "", log.Ldate|log.Ltime|log.LUTC),
	}
}

func (l *Logger) SetLogLevel(level LoggingLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

func (l *Logger) Logging(level LoggingLevel, msg string) {
	if level == LevelError || level <= l.logLevel {
		l.logOut.Printf("%s: %s log_level=%s", level, msg, l.logLevel)
	}
}

func (l *Logger) Error(msg string) {
	l.Logging(LevelError, msg)
}

func (l *Logger) Close() {
	l.out.Close()
}
