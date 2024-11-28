package cosgolog

import (
	"fmt"
	"os"
	"io"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	PANIC
)

var (
	logger *Logger
)

// Levels: DEBUG, INFO, WARN, ERROR, PANIC
// By default, the color is as follows:
// DEBUG: green
// INFO: white(default text color)
// WARN: yellow
// ERROR: red
// PANIC: magenta
// If `SetOutput()` is not called, the log will be written to the file `logs.log` and the console.
type Logger struct {
	level int
	logFile *os.File
	writer io.Writer
}

func defaultLogger() *Logger {
	if logger == nil {
		f, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Error while opening log file!")
			panic(err)
		}
		logger = &Logger{
			level: 0,
			logFile: f,
		}
		logger.writer = io.MultiWriter(f, os.Stdout)
	}
	return logger
}


func(l *Logger) Write(str string) {
	l.writer.Write([]byte(str))
}

func(l *Logger) Writeln(str string) {
	l.writer.Write([]byte(str + "\n"))
}

func Debug(v ...any) {
	l := defaultLogger()
	if l.level <= 0 {
		l.Writeln(green(Prefix(), v...))
	}
}

func Debugf(format string, v ...any) {
	l := defaultLogger()
	if l.level <= 0 {
		l.Write(greenf(Prefix(), format, v...))
	}
}

func Info(v ...any) {
	l := defaultLogger()
	if l.level <= 1 {
		l.Writeln(Prefix() + fmt.Sprint(v...))
	}
}

func Infof(format string, v ...any) {
	l := defaultLogger()
	if l.level <= 1 {
		l.Write(fmt.Sprintf(format, v...))
	}
}

func Warn(v ...any) {
	l := defaultLogger()
	if l.level <= 2 {
		l.Writeln(yellow(Prefix(), v...))
	}
}

func Warnf(format string, v ...any) {
	l := defaultLogger()
	if defaultLogger().level <= 2 {
		l.Write(yellowf(Prefix(), format, v...))
	}
}

func Error(v ...any) {
	l := defaultLogger()
	if l.level <= 3 {
		l.Writeln(red(Prefix(), v...))
	}
}

func Errorf(format string, v ...any) {
	l := defaultLogger()
	if defaultLogger().level <= 3 {
		l.Write(redf(Prefix(), format, v...))
	}
}

func Panic(v ...any) {
	l := defaultLogger()
	if l.level <= 4 {
		l.Writeln(magenta(Prefix(), v...))
	}
}

func Panicf(format string, v ...any) {
	l := defaultLogger()
	if defaultLogger().level <= 4 {
		l.Write(magentaf(Prefix(), format, v...))
	}
}
