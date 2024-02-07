package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	f *os.File
}

func NewLogger(fileName string) (*Logger, error) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	mw := io.MultiWriter(f, os.Stdout)
	log.SetOutput(mw)
	return &Logger{f}, nil
}

func (l *Logger) Close() error {
	return l.f.Close()
}

func (l *Logger) Info(str string) {
	log.Println(str)
}

func (l *Logger) Error(err error) {
	log.Println(err)
}

func (l *Logger) Warn(str string) {
	log.Println(str)
}
