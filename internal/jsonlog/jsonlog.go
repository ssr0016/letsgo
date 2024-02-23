package jsonlog

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

const (
	LevelInfo  Level = iota // Has the value 0
	LevelError              // Has the value 1
	LevelFatal              // Has the value 2
	LevelOff                // Has the value 3
)

// Return a human-friendly string for the severity level.
func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// Define custome logger
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

func (l *Logger) PrintInfo(msg string, properties map[string]string) {
	l.print(LevelInfo, msg, properties)
}

func (l *Logger) PrintError(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
}

func (l *Logger) PrintFatal(err error, properties map[string]string) {
	l.print(LevelFatal, err.Error(), properties)
	os.Exit(1)
}

func (l *Logger) print(level Level, msg string, properties map[string]string) (int, error) {
	// If the severity level of the log entry is below the minimum severity for the
	// logger, then return with no further action.

	if level < l.minLevel {
		return 0, nil
	}

	// Declare an anonymous struct holding the data for the log entry.
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    msg,
		Properties: properties,
	}

	// Include a stack trace for entries at the ERROR and FATAL levels
	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	// Declare a line variable for holding the actual log entry text.
	var line []byte
	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.String() + ": unable to marshal log message" + err.Error())
	}

	// Lock the mutex so that no two writes to the output destination cannot happen
	// concurrently. If we don't do this, it's possible that the text for twoor more
	// log entries will be intermingled in the output.

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}
