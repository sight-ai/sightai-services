package log

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger
var colorer *color.Color

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: true}

	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s,", i)
	}

	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}

	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s,", i)
	}

	l := zerolog.New(output).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	logger = l

	colorer = color.New()
}

func Initialize(appName string, noColor bool, level zerolog.Level) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: noColor}

	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s,", i)
	}

	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}

	output.FormatLevel = func(i interface{}) string {
		return fmt.Sprintf("%s %s", appName, levelToString(i))
	}

	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s,", i)
	}

	l := zerolog.New(output).Level(level).With().Timestamp().Logger()
	logger = l
}

func InitializeJson(level zerolog.Level, appName string) {
	f, err := os.OpenFile(fmt.Sprintf("%s.log", appName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	logger = zerolog.New(f).Level(level).With().Timestamp().Logger()
}

func levelToString(i interface{}) string {
	var l string
	if ll, ok := i.(string); ok {
		switch ll {
		case "debug":
			l = colorer.Magenta("DBG")
		case "info":
			l = colorer.Green("INF")
		case "warn":
			l = colorer.Yellow("WRN")
		case "error":
			l = colorer.Red("ERR")
		case "fatal":
			l = colorer.RedBg("FTL")
		case "panic":
			l = "PNC"
		default:
			l = "???"
		}
	} else {
		if i == nil {
			l = "???"
		} else {
			l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
		}
	}
	return l
}

// Debug ...
func Debug(args ...interface{}) *zerolog.Event {
	// do not log line for debug
	//e := logger.Debug().Str("location", logLocation())
	e := logger.Debug()
	e = reqID(e, args...)

	return e
}

// Info ...
func Info(args ...interface{}) *zerolog.Event {
	// do not log line for info
	//e := logger.Info().Str("location", logLocation())
	e := logger.Info()
	e = reqID(e, args...)

	return e
}

// Warn ...
func Warn(args ...interface{}) *zerolog.Event {
	e := logger.Warn().Str("location", logLocation())
	e = reqID(e, args...)

	return e
}

// Error ...
func Error(args ...interface{}) *zerolog.Event {
	e := logger.Error().Str("location", logLocation())
	e = reqID(e, args...)

	return e
}

// Fatal ...
func Fatal(args ...interface{}) *zerolog.Event {
	e := logger.Fatal().Str("location", logLocation())
	e = reqID(e, args...)

	return e
}

// Panic ...
func Panic(args ...interface{}) *zerolog.Event {
	e := logger.Panic().Str("location", logLocation())
	e = reqID(e, args...)

	return e
}

func GetLogger() zerolog.Logger {
	return logger
}

func logLocation() string {
	if pc, file, line, ok := runtime.Caller(2); ok {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return "?:?"
		}
		fNameSlice := strings.Split(file, "/")
		fName := fNameSlice[len(fNameSlice)-1]
		fnNameSlice := strings.Split(fn.Name(), ".")
		fnName := fnNameSlice[len(fnNameSlice)-1]
		return fmt.Sprintf("%v(%v):%v", fName, line, fnName)
	}
	return "?:?"
}

func reqID(e *zerolog.Event, args ...interface{}) *zerolog.Event {
	if len(args) > 0 && args[0] != nil {
		if c, ok := args[0].(echo.Context); ok {
			reqID := c.Get(echo.HeaderXRequestID).(string)
			return e.Str("req_id", reqID)
		} else if c, ok := args[0].(context.Context); ok {
			reqID := c.Value(echo.HeaderXRequestID)
			if reqID != nil && reqID.(string) != "" {
				return e.Str("req_id", reqID.(string))
			}
		}
	}
	return e
}
