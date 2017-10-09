package logger

import (
	"io"
	"os"
	"time"

	"github.com/kwri/go-workflow/modules/fs"
	"github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("workflow")
	// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
	Fatal = log.Fatal

	// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
	Fatalf = log.Fatalf

	// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
	Panic = log.Panic

	// Panicf is equivalent to l.Critical followed by a call to panic().
	Panicf = log.Panicf

	// Critical logs a message using CRITICAL as log level.
	Critical = log.Critical

	// Criticalf logs a message using CRITICAL as log level.
	Criticalf = log.Criticalf

	// Error logs a message using ERROR as log level.
	Error = log.Error

	// Errorf logs a message using ERROR as log level.
	Errorf = log.Errorf

	// Warning logs a message using WARNING as log level.
	Warning = log.Warning

	// Warningf logs a message using WARNING as log level.
	Warningf = log.Warningf

	// Notice logs a message using NOTICE as log level.
	Notice = log.Notice

	// Noticef logs a message using NOTICE as log level.
	Noticef = log.Noticef

	// Info logs a message using INFO as log level.
	Info = log.Info

	// Infof logs a message using INFO as log level.
	Infof = log.Infof

	// Debug logs a message using DEBUG as log level.
	Debug = log.Debug

	// Debugf logs a message using DEBUG as log level.
	Debugf = log.Debugf
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func Initialize(command string) {
	logging.MustGetLogger("workflow")
	// Example format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	// For demo purposes, create two backend for os.Stderr.
	backend := logging.NewLogBackend(getLogOutput(command), "", 0)

	// For messages written to backend we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Set the backends to be used.
	logging.SetBackend(backendFormatter)
}

func getLogOutput(command string) io.Writer {
	filename := generateFileOutputPath(command)
	ensureFileOutput(filename)
	file, err := fs.OpenFile(filename)
	if err != nil {
		return os.Stderr
	}
	return io.MultiWriter(file, os.Stderr)
}

func generateFileOutputPath(command string) string {
	time := time.Now().Format("yyyy_mm_dd")
	fileName := "/var/log/workflow/" + command + "_" + time + ".log"
	return fileName
}

func ensureFileOutput(filename string) error {
	if fs.FileExists(filename) {
		return nil
	}

	return fs.CreateFile(filename)
}
