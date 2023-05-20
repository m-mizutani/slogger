package slogger

import "errors"

var (
	ErrInvalidLogFormat = errors.New("invalid log format")
	ErrInvalidLogLevel  = errors.New("invalid log level")
)
