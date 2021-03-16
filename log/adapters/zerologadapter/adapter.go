package zerologadapter

import (
	"github.com/rs/zerolog"
	"github.com/wildberries-ru/go-transport-generator/log/logger"
)

type adapter struct {
	log    zerolog.Logger
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
	a.log.Debug().Fields(a.fields).Err(a.err).Msgf(format, args...)
}

func (a *adapter) Debug(message string) {
	a.log.Debug().Fields(a.fields).Err(a.err).Msg(message)
}

func (a *adapter) Infof(format string, args ...interface{}) {
	a.log.Info().Fields(a.fields).Err(a.err).Msgf(format, args...)
}

func (a *adapter) Info(message string) {
	a.log.Info().Fields(a.fields).Err(a.err).Msg(message)
}

func (a *adapter) Warnf(format string, args ...interface{}) {
	a.log.Warn().Fields(a.fields).Err(a.err).Msgf(format, args...)
}

func (a *adapter) Warn(message string) {
	a.log.Warn().Fields(a.fields).Err(a.err).Msg(message)
}

func (a *adapter) Errorf(format string, args ...interface{}) {
	a.log.Error().Fields(a.fields).Err(a.err).Msgf(format, args...)
}

func (a *adapter) Error(message string) {
	a.log.Error().Fields(a.fields).Err(a.err).Msg(message)
}

func (a *adapter) Fatalf(format string, args ...interface{}) {
	a.log.Fatal().Fields(a.fields).Err(a.err).Msgf(format, args...)
}

func (a *adapter) Fatal(message string) {
	a.log.Fatal().Fields(a.fields).Err(a.err).Msg(message)
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
func New(log zerolog.Logger) logger.Logger {
	return &adapter{
		log: log,
	}
}
