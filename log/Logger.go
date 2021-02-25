package log

import (
	"fmt"
	"time"
)

var level *Level

func CreateLogger(setLevel Level) {
	if level == nil {
		level = &setLevel
	} else {
		log(LevelError, "Can't change the logging level at runtime!")
	}
}

func log(currLevel Level, msg string) {
	if level == nil {
		panic("LOGGER NOT INITIALIZED!")
	}

	if level.Value >= currLevel.Value {
		date := time.Now().Format("2006/01/02 - 15:04:05")
		text := fmt.Sprintf("%v [%v%v%v] - %v", date, currLevel.Color, currLevel.Text, reset, msg)

		// Might do log to file, etc...
		fmt.Println(text)
	}
}

func Debug(msg string) {
	log(LevelDebug, msg)
}

func Info(msg string) {
	log(LevelInfo, msg)
}

func InfoAlways(msg string) {
	log(LevelInfoAlwaysOn, msg)
}

func Warn(msg string) {
	log(LevelWarn, msg)
}

func Error(msg string) {
	log(LevelError, msg)
}