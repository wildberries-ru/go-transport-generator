package logger

// Logger ...
type Logger interface {
	WithError(err error) Logger
	WithFields(fields map[string]interface{}) Logger
	Debugf(format string, args ...interface{})
	Debug(message string)
	Infof(format string, args ...interface{})
	Info(message string)
	Warnf(format string, args ...interface{})
	Warn(message string)
	Errorf(format string, args ...interface{})
	Error(message string)
	Fatalf(format string, args ...interface{})
	Fatal(message string)
}
