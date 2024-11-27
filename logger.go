package cosgolog

import (
	"fmt"
	"log"
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
	logFile *os.File
)

// Levels: DEBUG, INFO, WARN, ERROR, PANIC
// By default, the color is as follows:
// DEBUG: green
// INFO: white(default text color)
// WARN: yellow
// ERROR: red
// PANIC: magenta
type Logger struct {
	level int
}

func defaultLogger() *Logger {
	if logger == nil {
		logger = &Logger{
			level: 0,
		}
	}
	if logFile == nil {
		var err error
		logFile, err = os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Error while opening log file!")
			panic(err)
		}
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return logger
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

// Wrapper for log.SetFlags()
func SetFlags(flag int) {
	log.SetFlags(flag)
}

func Debug(v ...any) {
	if defaultLogger().level <= 0 {
		log.Println(green(v...))
	}
}

func Debugf(format string, v ...any) {
	if defaultLogger().level <= 0 {
		log.Printf(greenf(format, v...))
	}
}

func Info(v ...any) {
	if defaultLogger().level <= 1 {
		log.Println(v...)
	}
}

func Infof(format string, v ...any) {
	if defaultLogger().level <= 1 {
		log.Printf(format, v...)
	}
}

func Warn(v ...any) {
	if defaultLogger().level <= 2 {
		log.Println(yellow(v...))
	}
}

func Warnf(format string, v ...any) {
	if defaultLogger().level <= 2 {
		log.Printf(yellowf(format, v...))
	}
}

func Error(v ...any) {
	if defaultLogger().level <= 3 {
		log.Println(red(v...))
	}
}

func Errorf(format string, v ...any) {
	if defaultLogger().level <= 3 {
		log.Printf(redf(format, v...))
	}
}

func Panic(v ...any) {
	if defaultLogger().level <= 4 {
		log.Println(magenta(v...))
	}
}

func Panicf(format string, v ...any) {
	if defaultLogger().level <= 4 {
		log.Printf(magentaf(format, v...))
	}
}
