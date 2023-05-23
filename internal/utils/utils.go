package utils

import (
	"Excel-C2/internal/configuration"
	"log"
	"os"
	"strconv"
	"time"
)

func GenerateNewSheetName() string {

	currentTime := time.Now()
	currentTimeS := currentTime.Format("02-01-2006")
	unixString := strconv.FormatInt(currentTime.Unix(), 10)
	hostname, err := os.Hostname()
	if err != nil {
		return currentTimeS + "-" + unixString[len(unixString)-5:]
	}

	return currentTimeS + "-" + hostname + "-" + unixString[len(unixString)-5:]
}

func LogDebug(message string) {
	if configuration.GetOptionsDebug() {
		log.Println(message)
	}
}

func LogFatalDebug(message string) {
	if configuration.GetOptionsDebug() {
		log.Fatal(message)
	}
}
