package slogger_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/m-mizutani/slogger"
	"golang.org/x/exp/slog"
)

func TestNewWithError(t *testing.T) {
	t.Run("DefaultConfig", func(t *testing.T) {
		logger, err := slogger.NewWithError()
		if err != nil {
			t.Fatal("expected no error, got:", err)
		}

		if logger == nil {
			t.Fatal("expected logger not to be nil")
		}
	})

	t.Run("InvalidFormat", func(t *testing.T) {
		_, err := slogger.NewWithError(slogger.WithFormat("invalid-format"))
		if !errors.Is(err, slogger.ErrInvalidLogFormat) {
			t.Fatal("expected invalid log format error, got:", err)
		}
	})

	t.Run("InvalidLevel", func(t *testing.T) {
		_, err := slogger.NewWithError(slogger.WithLevel("invalid-level"))
		if !errors.Is(err, slogger.ErrInvalidLogLevel) {
			t.Fatal("expected invalid log level error, got:", err)
		}
	})

	t.Run("ValidFormatAndLevel", func(t *testing.T) {
		logger, err := slogger.NewWithError(slogger.WithFormat("json"), slogger.WithLevel("debug"))
		if err != nil {
			t.Fatal("expected no error, got:", err)
		}

		if logger == nil {
			t.Fatal("expected logger not to be nil")
		}
	})

	t.Run("WithOutput", func(t *testing.T) {
		tmpfile, err := os.CreateTemp("", "slogger-test")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		logger, err := slogger.NewWithError(slogger.WithOutput(tmpfile.Name()))
		if err != nil {
			t.Fatal("expected no error, got:", err)
		}

		if logger == nil {
			t.Fatal("expected logger not to be nil")
		}
	})

	t.Run("WithReplacer", func(t *testing.T) {
		replacerFunc := func(groups []string, a slog.Attr) slog.Attr {
			return a
		}

		logger, err := slogger.NewWithError(slogger.WithReplacer(replacerFunc))
		if err != nil {
			t.Fatal("expected no error, got:", err)
		}

		if logger == nil {
			t.Fatal("expected logger not to be nil")
		}
	})
}
