package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	file        *os.File
}

func NewLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
		file:        file,
	}, nil
}

func (l *Logger) Info(msg string) {
	l.infoLogger.Output(2, msg)
}

func (l *Logger) Error(msg string) {
	l.errorLogger.Output(2, msg)
}

func (l *Logger) Close() {
	l.file.Close()
}
