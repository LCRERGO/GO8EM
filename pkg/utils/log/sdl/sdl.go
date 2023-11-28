// A package for logs on the sdl version
package sdl

import "log"

type Logger struct {
	l *log.Logger
}

func New() *Logger {
	return &Logger{
		l: log.Default(),
	}
}

func Print(logger *Logger, v ...any) {
	logger.l.Print(v...)
}

func Printf(logger *Logger, format string, v ...any) {
	logger.l.Printf(format, v...)
}

func Fatal(logger *Logger, v ...any) {
	logger.l.Fatal(v...)
	panic(v)
}

func Fatalf(logger *Logger, format string, v ...any) {
	logger.l.Fatalf(format, v...)
	panic(v)
}
