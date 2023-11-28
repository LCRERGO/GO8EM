package log

import (
	"github.com/LCRERGO/GO8EM/pkg/utils/log/sdl"
)

const (
	SDLLoggerType = "sdl"
)

var logger *Logger

type LoggerType string

type Logger struct {
	lType string
	sl    *sdl.Logger
}

func New(lType LoggerType) *Logger {
	switch lType {
	case SDLLoggerType:
		return &Logger{
			lType: SDLLoggerType,
			sl:    sdl.New(),
		}
		// invalid type
	default:
		return nil
	}
}

func Default() *Logger {
	if logger == nil {
		logger = New(SDLLoggerType)
	}

	return logger
}

func Print(v ...any) {
	defaultLogger := Default()

	switch defaultLogger.lType {
	case SDLLoggerType:
		sdl.Print(defaultLogger.sl, v...)
		// invalid
	default:
	}
}

func Printf(format string, v ...any) {
	defaultLogger := Default()

	switch defaultLogger.lType {
	case SDLLoggerType:
		sdl.Printf(defaultLogger.sl, format, v...)
		// invalid
	default:
	}
}

func Fatal(v ...any) {
	defaultLogger := Default()

	switch defaultLogger.lType {
	case SDLLoggerType:
		sdl.Fatal(defaultLogger.sl, v...)
		// invalid
	default:
	}
}

func Fatalf(format string, v ...any) {
	defaultLogger := Default()

	switch defaultLogger.lType {
	case SDLLoggerType:
		sdl.Fatalf(defaultLogger.sl, format, v...)
		// invalid
	default:
	}
}
