package main

import (
	"os"

	"github.com/jinzhu/gorm"
	logging "github.com/op/go-logging"
)

// format is the logging formatter for the go-logging package.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{longfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// InitializeLogging configures the go-logging package for the current environment.
func InitializeLogging() {
	stdoutBackend := logging.NewBackendFormatter(logging.NewLogBackend(os.Stdout, "", 0), format)
	stderrBackend := logging.AddModuleLevel(logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))
	stderrBackend.SetLevel(logging.ERROR, "")

	logging.SetBackend(stdoutBackend, stderrBackend)
}

var gormLog = logging.MustGetLogger("orm")

type GormLogging struct {
	gorm.Logger
}

func (GormLogging) Print(args ...interface{}) {
	gormLog.Info(args)
}
