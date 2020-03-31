package logger

import (
	"log"
	"os"
)

// type Trace func(format string, v ...interface{})
var (
	Trace func(format string, v ...interface{})
	Info  func(format string, v ...interface{})
	Warn  func(format string, v ...interface{})
	Error func(format string, v ...interface{})
)

func init() {
	Trace = log.New(os.Stdout, "[TRACE]: ", log.Ldate|log.Ltime|log.Lshortfile).Printf
	Info = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile).Printf
	Warn = log.New(os.Stdout, "[WARN]: ", log.Ldate|log.Ltime|log.Lshortfile).Printf
	Error = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile).Printf
}
