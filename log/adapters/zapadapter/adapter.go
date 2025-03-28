package zapadapter

import (
	"fmt"

	"github.com/wildberries-ru/go-transport-generator/log/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type adapter struct {
	log    *zap.Logger
	err    error
	fields map[string]interface{}
}

func (a *adapter) WithError(err error) logger.Logger {
	return &adapter{log: a.log, err: err, fields: a.fields}
}

func (a *adapter) WithFields(fields map[string]interface{}) logger.Logger {
	return &adapter{log: a.log, err: a.err, fields: fields}
}

func (a *adapter) Debugf(format string, args ...interface{}) {
	a.log.Debug(fmt.Sprintf(format, args...), a.convertFields()...)
}

func (a *adapter) Debug(message string) {
	a.log.Debug(message, a.convertFields()...)
}

func (a *adapter) Infof(format string, args ...interface{}) {
	a.log.Info(fmt.Sprintf(format, args...), a.convertFields()...)
}

func (a *adapter) Info(message string) {
	a.log.Info(message, a.convertFields()...)
}

func (a *adapter) Warnf(format string, args ...interface{}) {
	a.log.Warn(fmt.Sprintf(format, args...), a.convertFields()...)
}

func (a *adapter) Warn(message string) {
	a.log.Warn(message, a.convertFields()...)
}

func (a *adapter) Errorf(format string, args ...interface{}) {
	a.log.Error(fmt.Sprintf(format, args...), a.convertFields()...)
}

func (a *adapter) Error(message string) {
	a.log.Error(message, a.convertFields()...)
}

func (a *adapter) Fatalf(format string, args ...interface{}) {
	a.log.Fatal(fmt.Sprintf(format, args...), a.convertFields()...)
}

func (a *adapter) Fatal(message string) {
	a.log.Fatal(message, a.convertFields()...)
}

func (a *adapter) convertFields() []zapcore.Field {
	fields := make([]zapcore.Field, 0)

	for k, v := range a.fields {
		fields = append(fields, zap.Any(k, v))
	}

	if a.err != nil {
		fields = append(fields, zap.Error(a.err))
	}

	return fields
}

// New ...
func New(logger *zap.Logger) logger.Logger {
	return &adapter{
		log: logger,
	}
}
