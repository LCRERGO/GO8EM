// A package for logs on the sdl version
package sdl

import "log"

type Logger interface {
	Print(v ...any)
	Printf(format string, v ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
}

type logger struct {
	l *log.Logger
}

func New() Logger {
	return &logger{
		l: log.Default(),
	}
}

func (logger *logger) Print(v ...any) {
	logger.l.Print(v...)
}

func (logger *logger) Printf(format string, v ...any) {
	logger.l.Printf(format, v...)
}

func (logger *logger) Fatal(v ...any) {
	logger.l.Fatal(v...)
	panic(v)
}

func (logger *logger) Fatalf(format string, v ...any) {
	logger.l.Fatalf(format, v...)
	panic(v)
}
