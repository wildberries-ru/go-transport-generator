package kitadapter

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/wildberries-ru/go-transport-generator/log/logger"
)

type adapter struct {
	log    log.Logger
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
	_ = level.Debug(a.log).Log(a.convertFields(fmt.Sprintf(format, args...))...)
}

func (a *adapter) Debug(message string) {
	_ = level.Debug(a.log).Log(a.convertFields(message)...)
}

func (a *adapter) Infof(format string, args ...interface{}) {
	_ = level.Info(a.log).Log(a.convertFields(fmt.Sprintf(format, args...))...)
}

func (a *adapter) Info(message string) {
	_ = level.Info(a.log).Log(a.convertFields(message)...)
}

func (a *adapter) Warnf(format string, args ...interface{}) {
	_ = level.Warn(a.log).Log(a.convertFields(fmt.Sprintf(format, args...))...)
}

func (a *adapter) Warn(message string) {
	_ = level.Warn(a.log).Log(a.convertFields(message)...)
}

func (a *adapter) Errorf(format string, args ...interface{}) {
	_ = level.Error(a.log).Log(a.convertFields(fmt.Sprintf(format, args...))...)
}

func (a *adapter) Error(message string) {
	_ = level.Error(a.log).Log(a.convertFields(message)...)
}

func (a *adapter) Fatalf(format string, args ...interface{}) {
	_ = level.Error(a.log).Log(a.convertFields(fmt.Sprintf(format, args...))...)
	os.Exit(1)
}

func (a *adapter) Fatal(message string) {
	_ = level.Error(a.log).Log(a.convertFields(message)...)
	os.Exit(1)
}

func (a *adapter) convertFields(msg string) []interface{} {
	fields := make([]interface{}, 0)

	fields = append(fields, "msg", msg)

	for k, v := range a.fields {
		fields = append(fields, k, v)
	}

	if a.err != nil {
		fields = append(fields, "error", a.err)
	}

	return fields
}

// New ...
func New(log log.Logger) logger.Logger {
	return &adapter{
		log: log,
	}
}
