package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	LOG_PATH       = "./logger/logs/"
	ENTRIES_CUTTER = "--------------------------------------------------------------"
	// LEVELS
	DEBUG   = "[DEBUG]"
	INFO    = "[INFO]"
	WARNING = "[WARNING]"
	ERROR   = "[ERROR]"
	FATAL   = "[FATAL]"
)

type LogEntry struct {
	DateTime time.Time
	Level    string
	Location string
	Content  string
}

func Logger(entry LogEntry) {
	date := fmt.Sprintf("(%v%v%v)", entry.DateTime.Day(), entry.DateTime.Month(), entry.DateTime.Year())

	logFile, err := os.OpenFile(LOG_PATH+"logby_"+date+fmt.Sprintf("(%s)", entry.Level)+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Cannot create a new log file!\n%s", err)
	}
	defer logFile.Close()
	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)

	log.Printf("\n" + ENTRIES_CUTTER + "\nLOG LEVEL: " + entry.Level +
		"\nLOCATION: " + entry.Location + fmt.Sprintf(" AT %v", entry.DateTime) + "\nCONTENT:\n" + entry.Content +
		"\n" + ENTRIES_CUTTER)
}
