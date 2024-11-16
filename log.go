package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
)

var Level = struct {
	Trace string
	Debug string
	Info  string
	Warn  string
	Error string
	Fatal string
}{
	Trace: "trace",
	Debug: "debug",
	Info:  "info",
	Warn:  "warn",
	Error: "error",
	Fatal: "fatal",
}

var logger zerolog.Logger

func debugMode(level zerolog.Level) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		col := color.New(color.FgWhite, color.Bold)
		switch i {
		case "trace":
			col = color.New(color.FgCyan)
		case "debug":
			col = color.New(color.FgBlue)
		case "info":
			col = color.New(color.FgGreen)
		case "warn":
			col = color.New(color.FgYellow)
		case "error":
			col = color.New(color.FgRed)
		case "fatal":
			col = color.New(color.FgMagenta)
		}
		return col.Sprintf(" %-6s", i)
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf(" %s ", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s ", i)
	}

	logger = zerolog.New(output).With().Timestamp().Logger()
	fmt.Println(level)
	logger.Level(level)
	zerolog.SetGlobalLevel(level)
}

func defaultMode(level zerolog.Level) {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Level(level)
}

func SetLevel(level string) {
	switch strings.ToLower(level) {
	case Level.Trace:
		defaultMode(zerolog.InfoLevel)
	case Level.Debug:
		debugMode(zerolog.DebugLevel)
	case Level.Info:
		defaultMode(zerolog.InfoLevel)
	case Level.Warn:
		defaultMode(zerolog.InfoLevel)
	case Level.Error:
		defaultMode(zerolog.InfoLevel)
	case Level.Fatal:
		defaultMode(zerolog.InfoLevel)
	default:
		defaultMode(zerolog.InfoLevel)
	}

}

func getLogLevelByLib(l zerolog.Level) string {
	switch l {
	case zerolog.TraceLevel:
		return Level.Trace
	case zerolog.DebugLevel:
		return Level.Debug
	case zerolog.InfoLevel:
		return Level.Info
	case zerolog.WarnLevel:
		return Level.Warn
	case zerolog.ErrorLevel:
		return Level.Error
	case zerolog.FatalLevel:
		return Level.Fatal
	default:
		return Level.Info
	}
}

func Trace(message interface{}, args ...interface{}) {
	msg(Level.Trace, message, args...)
}

func Debug(message interface{}, args ...interface{}) {
	msg(Level.Debug, message, args...)
}

func Info(message string, args ...interface{}) {
	msg(Level.Info, message, args...)
}

func Warn(message string, args ...interface{}) {
	msg(Level.Warn, message, args...)
}

func Error(message interface{}, args ...interface{}) {
	if getLevel() == Level.Debug {
		msg(Level.Debug, message, args...)
	}

	msg(Level.Error, message, args...)
}

func getLevel() string {
	l := logger.GetLevel()
	return l.String()
}

func Fatal(message interface{}, args ...interface{}) {
	msg(Level.Fatal, message, args...)
	os.Exit(1)
}

func msg(level string, message interface{}, args ...interface{}) {
	var event *zerolog.Event

	switch level {
	case Level.Trace:
		event = logger.Trace()
	case Level.Debug:
		event = logger.Debug()
	case Level.Info:
		event = logger.Info()
	case Level.Warn:
		event = logger.Warn()
	case Level.Error:
		event = logger.Error()
	case Level.Fatal:
		event = logger.Fatal()
	default:
		event = logger.Info()
	}

	callerInfo := getCallerInfo(3)

	switch msg := message.(type) {
	case error:
		event.Msgf("%s %s", append([]interface{}{callerInfo, msg.Error()}, args...)...)
	case string:
		event.Msgf("%s %s", append([]interface{}{callerInfo, msg}, args...)...)
	default:
		event.Msgf("%s %s message %v has unknown type %v", append([]interface{}{callerInfo, level, message, msg}, args...)...)
	}
}

func getCallerInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}

	funcPath := runtime.FuncForPC(pc).Name()
	funcSlice := strings.Split(funcPath, "/")
	funcName := funcSlice[len(funcSlice)-1]

	fmt.Println(getLevel())
	if getLevel() == Level.Debug {
		return fmt.Sprintf("%s:%d  |  %s  |", file, line, funcName)
	}

	return fmt.Sprintf("%s:%d %s", file, line, funcName)

}
