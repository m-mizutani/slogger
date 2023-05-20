package slogger

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
)

type config struct {
	format string
	level  string
	output string

	replacers []func(groups []string, a slog.Attr) slog.Attr
}

type Option func(*config)

// WithFormat sets the log format. Valid values are "text" and "json".
func WithFormat(format string) Option {
	return func(c *config) {
		c.format = format
	}
}

// WithLevel sets the log level. Valid values are "debug", "info", "warn" and "error".
func WithLevel(level string) Option {
	return func(c *config) {
		c.level = level
	}
}

// WithOutput sets the log output. Valid values are "-", "stdout", "stderr" and a file path.
func WithOutput(output string) Option {
	return func(c *config) {
		c.output = output
	}
}

// WithReplacer sets the log replacer.
func WithReplacer(replacer func(groups []string, a slog.Attr) slog.Attr) Option {
	return func(c *config) {
		c.replacers = append(c.replacers, replacer)
	}
}

// New creates a new logger of slog with options.
func New(options ...Option) *slog.Logger {
	logger, err := NewWithError(options...)
	if err != nil {
		panic(err)
	}
	return logger
}

var logLevelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func NewWithError(options ...Option) (*slog.Logger, error) {
	cfg := &config{
		format: "text",
		level:  "info",
		output: "stdout",
	}

	for _, opt := range options {
		opt(cfg)
	}

	logLevel, ok := logLevelMap[cfg.level]
	if !ok {
		return nil, ErrInvalidLogLevel
	}

	var w io.Writer
	switch cfg.output {
	case "-", "stdout":
		w = os.Stdout
	case "stderr":
		w = os.Stderr
	default:
		fd, err := os.Create(filepath.Clean(cfg.output))
		if err != nil {
			return nil, err
		}
		w = fd
	}

	opt := &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			for _, f := range cfg.replacers {
				attr = f(groups, attr)
			}
			return attr
		},
	}

	var newLogger *slog.Logger
	switch cfg.format {
	case "text":
		newLogger = slog.New(slog.NewTextHandler(w, opt))
	case "json":
		newLogger = slog.New(slog.NewJSONHandler(w, opt))
	default:
		return nil, ErrInvalidLogFormat
	}

	return newLogger, nil
}
