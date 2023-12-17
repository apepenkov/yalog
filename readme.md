# YALog: Yet Another Logger for Go
YALog is a versatile and customizable logging package for Go, designed to provide a robust solution for logging in Go applications. It offers a variety of features, such as customizable verbosity levels, caller information, timestamp formatting, and colored output, to enhance the logging experience.

## Features

- **Verbosity Levels**: Supports multiple levels of verbosity, including Debug, Info, Warning, Error, and Fatal.
- **Caller Information**: Option to include caller information (file and line number) in logs.
- **Timestamps**: Customizable timestamp formatting in log messages.
- **Colored Output**: Color-coded logs for better readability and differentiation between log levels.
- **Customizable Options**: Various options to tailor the logger to specific needs, such as adjusting caller padding, time formatting, and more.
- **Hierarchical Logging**: Support for hierarchical loggers with inheritance of settings and automatic name tree adjustment.

## Installation

To install YALog, use the `go get` command:
##
```bash
go get github.com/apepenkov/yalog
```



## Usage

### Basic Setup

```go
import "github.com/yourusername/yalog"

func main() {
    logger := yalog.NewLogger("myLogger")
    logger.Info("This is an info log")
}
```


### Setting Verbosity Level
The default verbosity level is `VerboseLevelInfo`. To change the verbosity level, use the `WithVerboseLevel` option when creating a new logger.

```go
logger := yalog.NewLogger("myLogger", yalog.WithVerboseLevel(yalog.VerboseLevelDebug))
logger.Debug("This is a debug log")
```

## Enabling Caller Information
This option is disabled by default. To enable caller information, use the `WithPrintCaller` option when creating a new logger. The argument to this option is the number of padding spaces to add after the caller information.

```go
logger := yalog.NewLogger("myLogger", yalog.WithPrintCaller(20))
logger.Info("Log with caller info")
```

## Customizing Timestamps
This option is disabled by default. To enable timestamps, use the `WithPrintTime` option when creating a new logger. The argument to this option is the timestamp format string, which must be compatible with the `time.Format` function.

```go
logger := yalog.NewLogger("myLogger", yalog.WithPrintTime("2006-01-02 15:04:05"))
logger.Info("Log with custom timestamp")
```

## Enabling Colored Output
This option is disabled by default. To enable colored output, use the `WithColorEnabled` option when creating a new logger.


```go
logger := yalog.NewLogger("myLogger", yalog.WithColorEnabled())
logger.Info("This is a colored info log")
```

## Changing Log Level Colors
You can change the color of each log level by using the `WithAnotherColor` option when creating a new logger. The first argument to this option is the log level, and the second argument is the color to use for that log level. The color must be one of the constants defined in the `github.com/apepenkov/yalog` package.

```go
logger := yalog.NewLogger("myLogger", yalog.WithAnotherColor(yalog.VerboseLevelInfo, yalog.ColorCyan))
logger.Info("Info log with custom color")
```

## Creating Hierarchical Loggers
You can create a hierarchical logger by using the `NewLogger` method of an existing logger. The new logger will inherit all settings from the parent logger, but will have its own name. The name of the new logger will be the name of the parent logger, followed by a dot, followed by the name of the new logger. For example, if you create a new logger with the name `child` from a logger with the name `parent`, the name of the new logger will be `parent.child`. You can create as many levels of hierarchy as you want.

For that, WithPrintTreeName option is used. The first argument to this option is the number of padding spaces to add after the logger name, and the second argument is a boolean value indicating whether this value should be adjusted automatically based on the maximum logger name length in the tree.

```go
parentLogger := yalog.NewLogger("parent", WithPrintTreeName(10, true))
childLogger := parentLogger.NewLogger("child")
childLogger.Info("This log is from the child logger")
```