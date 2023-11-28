package log

type Logger interface {
	Print(v ...any)
	Printf(format string, v ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
}
