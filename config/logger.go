package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (L *Logger) Debug(v ...interface{}) {
	L.debug.Println(v...)
}

func (L *Logger) Info(v ...interface{}) {
	L.info.Println(v...)
}

func (L *Logger) Warning(v ...interface{}) {
	L.warning.Println(v...)
}

func (L *Logger) Error(v ...interface{}) {
	L.err.Println(v...)
}

// Create Format Enabled Logs
func (L *Logger) DebugF(format string, v ...interface{}) {
	L.debug.Printf(format, v...)
}

func (L *Logger) InfoF(format string, v ...interface{}) {
	L.info.Printf(format, v...)
}

func (L *Logger) WarningF(format string, v ...interface{}) {
	L.warning.Printf(format, v...)
}

func (L *Logger) ErrorF(format string, v ...interface{}) {
	L.err.Printf(format, v...)
}
