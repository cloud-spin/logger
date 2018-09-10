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
	"testing"
)

type logData struct {
	level  byte
	format string
	v      []interface{}
}

var (
	invalidConfigsTests = []struct {
		name    string
		configs *Configs
	}{
		{"nil configs", nil},
		{"invalid level", &Configs{Level: 10}},
	}

	loggingTests = []struct {
		name    string
		logFunc func(l Logger, format string, v ...interface{})
		logData *logData
	}{
		{
			"Log Debug",
			func(l Logger, format string, v ...interface{}) {
				l.Debug(format, v)
			},
			&logData{level: LevelDebug, format: "Debug message %s", v: []interface{}{"Debug format"}},
		},
		{
			"Log Info",
			func(l Logger, format string, v ...interface{}) {
				l.Info(format, v)
			},
			&logData{level: LevelInfo, format: "Info message %s", v: []interface{}{"Info format"}},
		},
		{
			"Log Warn",
			func(l Logger, format string, v ...interface{}) {
				l.Warn(format, v)
			},
			&logData{level: LevelWarn, format: "Warn message %s", v: []interface{}{"Warn format"}},
		},
		{
			"Log Error",
			func(l Logger, format string, v ...interface{}) {
				l.Error(format, v)
			},
			&logData{level: LevelError, format: "Error message %s", v: []interface{}{"Error format"}},
		},
		{
			"Log Critical",
			func(l Logger, format string, v ...interface{}) {
				l.Critical(format, v)
			},
			&logData{level: LevelCritical, format: "Critical message %s", v: []interface{}{"Critical format"}},
		},
	}
)

func TestNewConfigsShouldReturnConfigsWithDefaultValuesSet(t *testing.T) {
	configs := NewConfigs()

	if configs.Enabled != true {
		t.Errorf("Expected: %t; Got: %t", true, configs.Enabled)
	}
	if configs.Level != LevelInfo {
		t.Errorf("Expected: %d; Got: %d", LevelInfo, configs.Level)
	}
}

func TestNewLoggerWithInvalidConfigsShouldReturnError(t *testing.T) {
	for _, tt := range invalidConfigsTests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.configs)
			if err == nil {
				t.Errorf("Expected: %s; Got: %s", "error", "success")
			}
			if logger != nil {
				t.Errorf("Expected: %s; Got: %s", "nil", "instance")
			}
		})
	}
}

func TestLogAllLevelsShouldLogMessages(t *testing.T) {
	configs := &Configs{
		Enabled: true,
		Level:   LevelDebug,
	}
	logger, err := NewLogger(configs)
	if err != nil {
		t.Errorf("Expected: logger initialized; Got: %s", err.Error())
	}
	var lastLoggedData *logData
	logger.RegisterOnLog(func(level byte, format string, v ...interface{}) {
		lastLoggedData = &logData{
			level:  level,
			format: format,
			v:      v,
		}
	})

	for _, tt := range loggingTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logFunc(logger, tt.logData.format, tt.logData.v)

			if lastLoggedData == nil {
				t.Error("Expected: logged data; Got: nil")
			}
			if lastLoggedData.level != tt.logData.level {
				t.Errorf("Expected: %d; Got: %d", tt.logData.level, lastLoggedData.level)
			}
			if lastLoggedData.format != tt.logData.format {
				t.Errorf("Expected: %s; Got: %s", tt.logData.format, lastLoggedData.format)
			}
			if len(lastLoggedData.v) != len(tt.logData.v) {
				t.Errorf("Expected: %d; Got: %d", len(lastLoggedData.v), len(tt.logData.v))
			}
		})
	}
}

func TestLogLevelShouldBeRespected(t *testing.T) {
	configs := &Configs{
		Enabled: true,
		Level:   LevelCritical,
	}
	logger, err := NewLogger(configs)
	if err != nil {
		t.Errorf("Expected: logger initialized; Got: %s", err.Error())
	}
	var lastLoggedData *logData
	logger.RegisterOnLog(func(level byte, format string, v ...interface{}) {
		lastLoggedData = &logData{
			level:  level,
			format: format,
			v:      v,
		}
	})

	logger.Debug("Debug", nil)
	if lastLoggedData != nil {
		t.Error("Expected: nil; Got: logged data")
	}
	logger.Info("Info", nil)
	if lastLoggedData != nil {
		t.Error("Expected: nil; Got: logged data")
	}
	logger.Warn("Warn", nil)
	if lastLoggedData != nil {
		t.Error("Expected: nil; Got: logged data")
	}
	logger.Error("Error", nil)
	if lastLoggedData != nil {
		t.Error("Expected: nil; Got: logged data")
	}
	logger.Critical("Critical", nil)
	if lastLoggedData == nil {
		t.Error("Expected: logged data; Got: nil")
	}
}

func TestLogDisabledShouldBeRespected(t *testing.T) {
	configs := &Configs{
		Enabled: false,
		Level:   LevelDebug,
	}
	logger, err := NewLogger(configs)
	if err != nil {
		t.Errorf("Expected: logger initialized; Got: %s", err.Error())
	}
	logger.RegisterOnLog(func(level byte, format string, v ...interface{}) {
		t.Error("Expected: no logged data; Got: logged data")
	})

	logger.Debug("Debug", nil)
	logger.Info("Info", nil)
	logger.Warn("Warn", nil)
	logger.Error("Error", nil)
	logger.Critical("Critical", nil)
}
