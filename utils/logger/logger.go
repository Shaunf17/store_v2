package logger

import (
	"errors"
	"io"
	"log"
	"os"
	"store/configs"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger

	writer io.Writer

	ErrCantOpenFile    = errors.New("Couldn't open specified log file")
	ErrFileAlreadyOpen = errors.New("Log file already open")
)

func Init(cfg configs.LoggerConfigs) error {
	file, err := os.OpenFile(cfg.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return ErrCantOpenFile
	}

	writer = io.Writer(io.Discard)

	if cfg.WriteToFile {
		writer = io.MultiWriter(writer, file)
	}
	if cfg.WriteToConsole {
		writer = io.MultiWriter(writer, os.Stdout)
	}

	InfoLogger = log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(writer, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(writer, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}
