package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	LEVEL_TRACE LogLevel = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_OFF
)

type Logger struct {
	*log.Logger

	Level  LogLevel
	Output *os.File
}

func NewLogger(f string, l LogLevel) (*Logger, error) {
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := &Logger{Output: file, Level: l}
	logger.Logger = log.New(file, "", log.Ldate|log.Ltime)

	return logger, nil
}

func NewLoggerFromConfig(file string) (*Logger, error) {
	props := make(map[string]string)
	err := LoadConfig(file, props)
	if err != nil {
		return nil, err
	}

	var level LogLevel
	switch props["level"] {
	case "TRACE", "trace":
		level = LEVEL_TRACE
	case "DEBUG", "debug":
		level = LEVEL_DEBUG
	case "INFO", "info":
		level = LEVEL_INFO
	case "WARN", "warn":
		level = LEVEL_WARN
	case "ERROR", "error":
		level = LEVEL_ERROR
	case "OFF", "off":
		level = LEVEL_ERROR
	default:
		return nil, errors.New(fmt.Sprintf("Error loading config: %s is not a valid level", props["level"]))
	}

	return NewLogger(props["file"], level)
}

func (logger *Logger) Close() {
	err := logger.Output.Close()
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
}

func (logger *Logger) Trace(v ...interface{}) {
	if logger.Level <= LEVEL_TRACE {
		logger.Logger.Println(append([]interface{}{"TRACE:"}, v...)...)
	}
}

func (logger *Logger) Tracef(format string, v ...interface{}) {
	if logger.Level <= LEVEL_TRACE {
		logger.Logger.Printf("TRACE: "+format+"\n", v...)
	}
}

func (logger *Logger) TraceBegin(f string) {
	if logger.Level <= LEVEL_TRACE {
		logger.Logger.Println("TRACE: Begin -", f)
	}
}

func (logger *Logger) TraceEnd(f string) {
	if logger.Level <= LEVEL_TRACE {
		logger.Logger.Println("TRACE: End -", f)
	}
}

func (logger *Logger) Debug(v ...interface{}) {
	if logger.Level <= LEVEL_DEBUG {
		logger.Logger.Println(append([]interface{}{"DEBUG:"}, v...)...)
	}
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	if logger.Level <= LEVEL_DEBUG {
		logger.Logger.Printf("DEBUG: "+format+"\n", v...)
	}
}

func (logger *Logger) Info(v ...interface{}) {
	if logger.Level <= LEVEL_INFO {
		logger.Logger.Println(append([]interface{}{"INFO:"}, v...)...)
	}
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	if logger.Level <= LEVEL_INFO {
		logger.Logger.Printf("INFO: "+format+"\n", v...)
	}
}

func (logger *Logger) Warn(v ...interface{}) {
	if logger.Level <= LEVEL_WARN {
		logger.Logger.Println(append([]interface{}{"WARN:"}, v...)...)
	}
}

func (logger *Logger) Warnf(format string, v ...interface{}) {
	if logger.Level <= LEVEL_WARN {
		logger.Logger.Printf("WARN: "+format+"\n", v...)
	}
}

func (logger *Logger) Error(v ...interface{}) {
	if logger.Level <= LEVEL_ERROR {
		logger.Logger.Println(append([]interface{}{"ERROR:"}, v...)...)
	}
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	if logger.Level <= LEVEL_ERROR {
		logger.Logger.Printf("ERROR: "+format+"\n", v...)
	}
}

func (logger *Logger) Fatal(v ...interface{}) {
	logger.Logger.Fatalln(append([]interface{}{"FATAL:"}, v...)...)
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.Logger.Fatalf("FATAL: "+format+"\n", v...)
}
