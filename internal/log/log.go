package log

import "github.com/hashicorp/go-hclog"

var (
	defaultLogger = hclog.Default()
)

// Level represents a log level.
type Level = hclog.Level

// Logger levels used by paper robot.
const (
	LevelDebug Level = 2
	LevelInfo  Level = 3
	LevelWarn  Level = 4
	LevelError Level = 5
)

// Debug print information for programmer lowlevel analysis.
func Debug(msg string, args ...interface{}) {
	defaultLogger.Debug(msg, args...)
}

// Info print information about steady state operations.
func Info(msg string, args ...interface{}) {
	defaultLogger.Info(msg, args...)
}

// Warn print information about rare but handled events.
func Warn(msg string, args ...interface{}) {
	defaultLogger.Warn(msg, args...)
}

// Error print information about unrecoverable events.
func Error(msg string, args ...interface{}) {
	defaultLogger.Error(msg, args...)
}

// SetLevel will set level for default logger.
func SetLevel(lvl Level) {
	defaultLogger.SetLevel(lvl)
}
