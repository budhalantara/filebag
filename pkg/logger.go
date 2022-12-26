package pkg

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

type Logger struct {
	log.Logger
}

func NewLogger() *Logger {
	baseLogger := log.New(os.Stdout, "", log.LstdFlags)
	return &Logger{
		Logger: *baseLogger,
	}
}

var Log = NewLogger()

func (l *Logger) Trace(err error) {
	newError := errors.New(err.Error())
	l.Logger.Printf("%+v", newError)
}

func (l *Logger) Fatal(err error) {
	newError := errors.New(err.Error())
	l.Logger.Fatalf("%+v", newError)
}
