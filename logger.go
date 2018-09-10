// Copyright (c) 2018 cloud-spin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package logger

import (
	"errors"
	"log"
	"os"
)

const (
	// LevelCritical represents a log of 'critical' level.
	LevelCritical = byte(0)
	// LevelError represents a log of 'error' level.
	LevelError = byte(1)
	// LevelWarn represents a log of 'warn' level.
	LevelWarn = byte(2)
	// LevelInfo represents a log of 'info' level.
	LevelInfo = byte(3)
	// LevelDebug represents a log of 'debug' level.
	LevelDebug = byte(4)
	// criticalHeader holds the header printed before every log line for Critical logs.
	criticalHeader = "[CRITICAL] "
	// errorHeader holds the header printed before every log line for Error logs.
	errorHeader = "[ERROR] "
	// warnHeader holds the header printed before every log line for Warn logs.
	warnHeader = "[WARN] "
	// infoHeader holds the header printed before every log line for Info logs.
	infoHeader = "[INFO] "
	// debugHeader holds the header printed before every log line for Debug logs.
	debugHeader = "[DEBUG] "
)

// OnLogHandler is fired when a log is called.
type OnLogHandler = func(level byte, format string, v ...interface{})

// Configs holds Logger specific configs.
// Enabled holds whether the logger is enabled or not.
// Level holds the logging level to log. Valids levels are: 0 (Critical); 1 (Error); 2 (Warn); 3 (Info); 4 (Debug)
type Configs struct {
	Enabled bool
	Level   byte
}

// Logger provides standard logging methods around the standard log package.
// Each logged line is prefixed with the logging level (Debug, Info, Warn, ...).
// The Enabled and Level configurations are respected every time the logging methods are called.
type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Critical(format string, v ...interface{})
	RegisterOnLog(handler OnLogHandler)
}

// LoggerImpl implements a Logger.
type LoggerImpl struct {
	Configs        *Configs
	onLogHandler   OnLogHandler
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	warnLogger     *log.Logger
	errorLogger    *log.Logger
	criticalLogger *log.Logger
}

// NewConfigs initializes a new instance of Configs with default values.
func NewConfigs() *Configs {
	return &Configs{
		Enabled: true,
		Level:   LevelInfo,
	}
}

// NewLogger initializes a new instance of Logger.
func NewLogger(configs *Configs) (Logger, error) {
	if configs == nil {
		return nil, errors.New("configs are required")
	}
	if configs.Level < LevelCritical || configs.Level > LevelDebug {
		return nil, errors.New("configs.Level is invalid")
	}

	return &LoggerImpl{
		Configs:        configs,
		debugLogger:    log.New(os.Stdout, debugHeader, log.LstdFlags),
		infoLogger:     log.New(os.Stdout, infoHeader, log.LstdFlags),
		warnLogger:     log.New(os.Stdout, warnHeader, log.LstdFlags),
		errorLogger:    log.New(os.Stderr, errorHeader, log.LstdFlags),
		criticalLogger: log.New(os.Stderr, criticalHeader, log.LstdFlags),
	}, nil
}

// RegisterOnLog registers a handler to be executed every time log is called.
// Only one handler at a time is supported.
func (l *LoggerImpl) RegisterOnLog(handler OnLogHandler) {
	l.onLogHandler = handler
}

// Critical logs the message to the standard log package using Printf.
func (l *LoggerImpl) Critical(format string, v ...interface{}) {
	l.log(l.criticalLogger, LevelCritical, format, v...)
}

// Error logs a error message to the standard log package.
func (l *LoggerImpl) Error(format string, v ...interface{}) {
	l.log(l.errorLogger, LevelError, format, v...)
}

// Warn logs a warning message to the standard log package.
func (l *LoggerImpl) Warn(format string, v ...interface{}) {
	l.log(l.warnLogger, LevelWarn, format, v...)
}

// Info logs a info message to the standard log package.
func (l *LoggerImpl) Info(format string, v ...interface{}) {
	l.log(l.infoLogger, LevelInfo, format, v...)
}

// Debug logs a debug message to the standard log package.
func (l *LoggerImpl) Debug(format string, v ...interface{}) {
	l.log(l.debugLogger, LevelDebug, format, v...)
}

// log logs the message to the standard log package.
func (l *LoggerImpl) log(logger *log.Logger, level byte, format string, v ...interface{}) {
	if l.Configs.Enabled && l.Configs.Level >= level {
		logger.Printf(format, v...)

		if l.onLogHandler != nil {
			l.onLogHandler(level, format, v...)
		}
	}
}
