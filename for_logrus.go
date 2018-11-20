package reqid

import (
	"context"

	"github.com/sirupsen/logrus"
)

const (
	// LogFieldKeyReqID is the logfield key for RequestID
	LogFieldKeyReqID = "requestId"
)

var (
	loglvl = logrus.DebugLevel
)

// SetReqIDGlobalLogLevel can set log level for newly created FieldLogger instance
func SetReqIDGlobalLogLevel(lvl logrus.Level) {
	loglvl = lvl
}

// NewLoggerFromReqIDStr creates  a *logrus.Entry that has requestID as a field. A new LogField inst will be created if log is nil
func NewLoggerFromReqIDStr(reqID string, log logrus.FieldLogger) logrus.FieldLogger {

	if log == nil {
		nlog := logrus.New()
		nlog.SetLevel(loglvl)
		log = nlog.WithField(LogFieldKeyReqID, reqID)
		log.Debugln("New logrus FieldLogger initiated")
		return log
	}

	return log.WithField(LogFieldKeyReqID, reqID)
}

// NewLoggerFromReqIDCtx creates a *logrus.Entry that has requestID as a field.  A new LogField inst will be created if log is ni
func NewLoggerFromReqIDCtx(ctx context.Context, log logrus.FieldLogger) logrus.FieldLogger {
	reqID := GetReqID(ctx)

	return NewLoggerFromReqIDStr(reqID, log)
}
