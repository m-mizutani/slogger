# slogger [![test](https://github.com/m-mizutani/slogger/actions/workflows/test.yml/badge.svg)](https://github.com/m-mizutani/slogger/actions/workflows/test.yml)

The `slogger` package provides a simple builder of [slog](golang.org/x/exp/slog). This package allows you to configure the log format, log level, log output, and log replacers.

## Features

- Supports log formats: `text` and `json`
- Configurable log levels: `debug`, `info`, `warn`, and `error`
- Configurable log output: `stdout`, `stderr`, or a file path
- Customizable log replacers for modifying log attributes

## Usage

To use the `slogger` package, first import it:

```go
import "path/to/slogger"
```

Then, create a new logger with your desired options:

```go
logger := slogger.New(
    slogger.WithFormat("json"),
    slogger.WithLevel("debug"),
    slogger.WithOutput("logs/output.log"),
)
```

Now you can use the logger in your application:

```go
logger.Debug("Debug log message")
logger.Info("Info log message")
logger.Warn("Warn log message")
logger.Error("Error log message")
```

## Options

The following options are available for configuring the logger:

- `WithFormat(format string)`: Sets the log format. Valid values are `"text"` and `"json"`.
- `WithLevel(level string)`: Sets the log level. Valid values are `"debug"`, `"info"`, `"warn"`, and `"error"`.
- `WithOutput(output string)`: Sets the log output. Valid values are `"-"`, `"stdout"`, `"stderr"`, and a file path.
- `WithReplacer(replacer func(groups []string, a slog.Attr) slog.Attr)`: Sets the log replacer. This function takes the log attribute groups and the current log attribute, and returns the modified log attribute.

## Error Handling

If you want to handle errors when creating the logger, you can use the `NewWithError` function:

```go
logger, err := slogger.NewWithError(
    slogger.WithFormat("json"),
    slogger.WithLevel("debug"),
    slogger.WithOutput("logs/output.log"),
)
if err != nil {
    // handle error
}
```

This function returns the logger and an error if there is any issue with the provided options.

## License

Apache License 2.0