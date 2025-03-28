package logrusadapter

import (
	"github.com/sirupsen/logrus"
	"github.com/wildberries-ru/go-transport-generator/log/logger"
)

type adapter struct {
	log    logrus.FieldLogger
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
	a.log.WithError(a.err).WithFields(a.convertFields()).Debugf(format, args...)
}

func (a *adapter) Debug(message string) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Debug(message)
}

func (a *adapter) Infof(format string, args ...interface{}) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Infof(format, args...)
}

func (a *adapter) Info(message string) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Info(message)
}

func (a *adapter) Warnf(format string, args ...interface{}) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Warnf(format, args...)
}

func (a *adapter) Warn(message string) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Warn(message)
}

func (a *adapter) Errorf(format string, args ...interface{}) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Errorf(format, args...)
}

func (a *adapter) Error(message string) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Error(message)
}

func (a *adapter) Fatalf(format string, args ...interface{}) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Fatalf(format, args...)
}

func (a *adapter) Fatal(message string) {
	a.log.WithError(a.err).WithFields(a.convertFields()).Fatal(message)
}

func (a *adapter) convertFields() logrus.Fields {
	return logrus.Fields(a.fields)
}

// New ...
func New(log logrus.FieldLogger) logger.Logger {
	return &adapter{
		log: log,
	}
}
