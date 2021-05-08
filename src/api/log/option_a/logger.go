package option_a

import (
	"fmt"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
)

var (
	Log *logrus.Logger
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level: level,
		Out: os.Stdout,
	}

	if os.Getenv("GO_ENVIRONMENT") == "PRODUCTION" {
		fmt.Println("======== IsProduction ========")
		// For JSON to index capable in elastic search
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		fmt.Println("======== IsStaging ========")
		Log.Formatter = &logrus.TextFormatter{}
	}
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)
}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)
}

func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.WithFields(parseFields(tags...)).Error(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}