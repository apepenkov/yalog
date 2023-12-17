package yalog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type VerboseLevel int

const (
	VerboseLevelDebug VerboseLevel = iota
	VerboseLevelInfo
	VerboseLevelWarning
	VerboseLevelError
	VerboseLevelFatal
)

type Color string

const (
	ColorReset             Color = "\033[0m"
	ColorBlack             Color = "\033[30m"
	ColorRed               Color = "\033[31m"
	ColorGreen             Color = "\033[32m"
	ColorYellow            Color = "\033[33m"
	ColorBlue              Color = "\033[34m"
	ColorPurple            Color = "\033[35m"
	ColorCyan              Color = "\033[36m"
	ColorWhite             Color = "\033[37m"
	ColorBrightBlack       Color = "\033[30;1m"
	ColorBrightRed         Color = "\033[31;1m"
	ColorBrightGreen       Color = "\033[32;1m"
	ColorBrightYellow      Color = "\033[33;1m"
	ColorBrightBlue        Color = "\033[34;1m"
	ColorBrightPurple      Color = "\033[35;1m"
	ColorBrightCyan        Color = "\033[36;1m"
	ColorBrightWhite       Color = "\033[37;1m"
	BackgroundBlack        Color = "\033[40m"
	BackgroundRed          Color = "\033[41m"
	BackgroundGreen        Color = "\033[42m"
	BackgroundYellow       Color = "\033[43m"
	BackgroundBlue         Color = "\033[44m"
	BackgroundPurple       Color = "\033[45m"
	BackgroundCyan         Color = "\033[46m"
	BackgroundWhite        Color = "\033[47m"
	BackgroundBrightBlack  Color = "\033[40;1m"
	BackgroundBrightRed    Color = "\033[41;1m"
	BackgroundBrightGreen  Color = "\033[42;1m"
	BackgroundBrightYellow Color = "\033[43;1m"
	BackgroundBrightBlue   Color = "\033[44;1m"
	BackgroundBrightPurple Color = "\033[45;1m"
	BackgroundBrightCyan   Color = "\033[46;1m"
	BackgroundBrightWhite  Color = "\033[47;1m"
)

var defaultLevelNames = []string{
	VerboseLevelDebug:   "DEBUG",
	VerboseLevelInfo:    "INFO",
	VerboseLevelWarning: "WARNING",
	VerboseLevelError:   "ERROR",
	VerboseLevelFatal:   "FATAL",
}

var defaultColorCodes = []Color{
	VerboseLevelDebug:   ColorBrightBlack,
	VerboseLevelInfo:    ColorCyan,
	VerboseLevelWarning: ColorYellow,
	VerboseLevelError:   ColorRed,
	VerboseLevelFatal:   ColorBrightRed,
}

var levelCount = len(defaultLevelNames)

type Logger struct {
	verboseLevel VerboseLevel
	parent       *Logger
	allChildren  []*Logger

	name                   string
	printCaller            bool
	padCaller              int
	printTime              bool
	printLevel             bool
	autoAdjustPadNameChars bool
	padNameChars           int
	timeFormat             string
	doColor                bool
	printNameTree          bool

	levelColors []Color
}

func NewLogger(name string, options ...Option) *Logger {
	l := &Logger{
		verboseLevel: VerboseLevelInfo,
		name:         name,

		printCaller:   false,
		padCaller:     0,
		printTime:     false,
		printLevel:    false,
		padNameChars:  len(name),
		timeFormat:    time.RFC3339,
		doColor:       false,
		printNameTree: false,
		levelColors:   make([]Color, levelCount),
	}
	copy(l.levelColors, defaultColorCodes)
	for _, option := range options {
		option(l)
	}
	return l
}

func (l *Logger) RecursiveAddChild(another *Logger) {
	l.allChildren = append(l.allChildren, another)
	if l.parent == nil {
		if l.autoAdjustPadNameChars {
			// we reached the root, now we need to propagate the `padNameChars` value accordingly to all children
			// first, we need to find the maximum `padNameChars` value among all children
			maxPad := len(l.name)
			for _, child := range l.allChildren {
				if len(child.getTreeNameInner()) > maxPad {
					maxPad = len(child.getTreeNameInner())
				}
			}

			for _, child := range l.allChildren {
				if child.autoAdjustPadNameChars {
					child.padNameChars = maxPad
				}
			}
			l.padNameChars = maxPad
		}
		return
	}
	l.parent.RecursiveAddChild(another)
}

// NewLogger
// creates a new logger and inherits the options from the parent
func (l *Logger) NewLogger(name string, options ...Option) *Logger {
	newl := &Logger{
		verboseLevel: l.verboseLevel,
		parent:       l,

		name: name,

		printCaller:            l.printCaller,
		padCaller:              l.padCaller,
		printTime:              l.printTime,
		printLevel:             l.printLevel,
		autoAdjustPadNameChars: l.autoAdjustPadNameChars,
		padNameChars:           l.padNameChars,
		timeFormat:             l.timeFormat,
		doColor:                l.doColor,
		printNameTree:          l.printNameTree,
		levelColors:            make([]Color, levelCount),
	}
	copy(newl.levelColors, l.levelColors)
	for _, option := range options {
		option(newl)
	}
	l.RecursiveAddChild(newl)
	return newl
}

func (l *Logger) getTreeNameInner() string {
	if !l.printNameTree {
		return l.name
	}
	fullName := l.name
	ptr := l.parent
	for ptr != nil {
		fullName = ptr.name + "." + fullName
		ptr = ptr.parent
	}
	return fullName
}

func (l *Logger) getTreeName() string {
	if l.padNameChars > 0 {
		return fmt.Sprintf(" [%-*s] ", l.padNameChars, l.getTreeNameInner())
	}
	return " [" + l.getTreeNameInner() + "] "
}

func (l *Logger) getDateTime() string {
	if !l.printTime {
		return ""
	}
	return time.Now().Format(l.timeFormat)
}

func (l *Logger) getCaller() string {
	if !l.printCaller {
		return ""
	}
	_, file, line, _ := runtime.Caller(3)
	file = filepath.Base(file)
	callerStr := fmt.Sprintf("%s:%d", file, line)
	if l.padCaller > 0 {
		callerStr = fmt.Sprintf("%-*s", l.padCaller, callerStr)
	}
	return " " + callerStr
}

func (l *Logger) getLevel(level VerboseLevel) string {
	if !l.printLevel {
		return ""
	}
	return fmt.Sprintf(" |%-*s| ", 7, defaultLevelNames[level])
}

func (l *Logger) print(level VerboseLevel, doNewLine bool, args ...any) {
	if level < l.verboseLevel {
		return
	}
	name := l.getTreeName()
	timeStr := l.getDateTime()
	levelStr := l.getLevel(level)
	callerStr := l.getCaller()

	var postfix string

	if doNewLine {
		postfix = "\n"
	} else {
		postfix = ""
	}

	res := fmt.Sprintf("%s%s%s%s > %s%s", timeStr, levelStr, name, callerStr, fmt.Sprint(args...), postfix)

	if l.doColor {
		res = string(l.levelColors[level]) + res + string(ColorReset)
	}

	print(res)
}

func (l *Logger) Debug(args ...any) {
	l.print(VerboseLevelDebug, false, args...)
}

func (l *Logger) Debugln(args ...any) {
	l.print(VerboseLevelDebug, true, args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.print(VerboseLevelDebug, false, fmt.Sprintf(format, args...))
}

func (l *Logger) Info(args ...any) {
	l.print(VerboseLevelInfo, false, args...)
}

func (l *Logger) Infoln(args ...any) {
	l.print(VerboseLevelInfo, true, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.print(VerboseLevelInfo, false, fmt.Sprintf(format, args...))
}

func (l *Logger) Warning(args ...any) {
	l.print(VerboseLevelWarning, false, args...)
}

func (l *Logger) Warningln(args ...any) {
	l.print(VerboseLevelWarning, true, args...)
}

func (l *Logger) Warningf(format string, args ...any) {
	l.print(VerboseLevelWarning, false, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(args ...any) {
	l.print(VerboseLevelError, false, args...)
}

func (l *Logger) Errorln(args ...any) {
	l.print(VerboseLevelError, true, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.print(VerboseLevelError, false, fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(args ...any) {
	l.print(VerboseLevelFatal, false, args...)
	os.Exit(1)
}

func (l *Logger) Fatalln(args ...any) {
	l.print(VerboseLevelFatal, true, args...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.print(VerboseLevelFatal, false, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (l *Logger) SetVerboseLevel(level VerboseLevel) {
	l.verboseLevel = level
}

func (l *Logger) V(level int) bool {
	return level <= int(l.verboseLevel)
}
