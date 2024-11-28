package cosgolog

import (
	"io"
	"time"
	"runtime"
	"strconv"
)

var (
	useDate = true
	dateFormat = "2006-01-02 15:04:05"
	useCaller = true
	callerFullPath = false
	useBracket = true
	bracketLeft = "["
	bracketRight = "]"
)

// Whether the log should have a date prefix
func SetDate(b bool) {
	useDate = b
}

// Takes the same format as `time.Format()`, default is "2006-01-02 15:04:05"
func DateFormat(format string) {
	dateFormat = format
}

// Whether the log should have a caller prefix
func SetCaller(b bool) {
	useCaller = b
}

// Full path for caller or just the file name
func SetCallerFullPath(b bool) {
	callerFullPath = b
}

// Set where the log should be written to
func SetOutput(w io.Writer) {
	defaultLogger().writer = w
}

// Whether the log should have brackets
func SetBracket(b bool) {
	useBracket = b
}

// Set the bracket for the log, left and right respectively
func SetBrackets(s string, e string) {
	bracketLeft = s
	bracketRight = e
}

// Levels: DEBUG, INFO, WARN, ERROR, PANIC
// 4: PANIC
// 3: ERROR + PANIC
// 2: WARN + ERROR + PANIC
// 1: INFO + WARN + ERROR + PANIC
// 0: DEBUG + INFO + WARN + ERROR + PANIC
func SetLogLevel(level int) {
	defaultLogger().level = level
}

// metadata: time, file, line
func date() string {
	return time.Now().Format(dateFormat)
}

func caller() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return "???"
	}
	if !callerFullPath {
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}
		}
	}
	return file + ":" + strconv.Itoa(line)
}

func Prefix() string {
	prefix := ""
	if useDate{
		prefix += date() + " "
	}
	if useCaller {
		prefix += caller() + " "
	}
	if len(prefix) > 0 && prefix[len(prefix)-1] == ' ' {
		prefix = prefix[:len(prefix)-1]
	}
	if useBracket {
		prefix = bracketLeft + prefix + bracketRight
	}
	if prefix != "" {
		prefix += " "
	}
	return prefix
}
