package yalog

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
