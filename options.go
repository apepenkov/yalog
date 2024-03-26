package yalog

import (
	"io"
	"os"
)

type Option func(*Logger)

// WithVerboseLevel
// sets the verbose level of the logger
func WithVerboseLevel(level VerboseLevel) Option {
	return func(l *Logger) {
		l.verboseLevel = level
	}
}

// WithPrintCaller
// enables printing of the caller (file:line)
func WithPrintCaller(pad int) Option {
	return func(l *Logger) {
		l.printCaller = true
		l.padCaller = pad
	}
}

// WithPrintTime
// enables printing of the time
func WithPrintTime(format string) Option {
	return func(l *Logger) {
		l.printTime = true
		l.timeFormat = format
	}
}

// WithPrintLevel
// enables printing of the level
func WithPrintLevel() Option {
	return func(l *Logger) {
		l.printLevel = true
	}
}

// WithColorEnabled
// enables color printing
func WithColorEnabled() Option {
	return func(l *Logger) {
		l.doColor = true
	}
}

// WithAnotherColor
// level: the level to change the color of
// color: the color to change the level to
func WithAnotherColor(level VerboseLevel, color Color) Option {
	return func(l *Logger) {
		l.levelColors[level] = color
	}
}

// WithPrintTreeName
// enables printing of the full name tree (parent1.parent2.name)
// autoAdjust bool: if true, pad will be adjusted to the length of the longest name in the tree
// pad: if autoAdjust is false, pad will be the length of the name
func WithPrintTreeName(pad int, autoAdjust bool) Option {
	return func(l *Logger) {
		l.printNameTree = true
		l.padNameChars = pad
		l.autoAdjustPadNameChars = autoAdjust
	}
}

// WithDifferentOutput
// sets a different output for the logger
// w: the writer to output to
func WithDifferentOutput(w io.Writer) Option {
	return func(l *Logger) {
		l.output = w
	}
}

// WithSecondOutput
// sets a second output for the logger
// level: the minimum level to output to the second output
func WithSecondOutput(w io.Writer, level VerboseLevel) Option {
	return func(l *Logger) {
		l.secondOutput = w
		l.secondOutputMinLevel = level
	}
}

// WithWriteToFile
// writes the log to a file
// file: os.File to write to
// flush: if true, the file will be flushed after every write
func WithWriteToFile(file *os.File, flush bool) Option {
	return func(l *Logger) {
		l.file = file
		l.flushFile = flush
	}
}
