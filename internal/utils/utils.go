package utils

import (
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
		return currentTimeS + "_" + unixString[len(unixString)-5:]
	}

	// return currentTimeS + "_" + hostname + "_" + unixString[len(unixString)-5:]
	return currentTimeS + "_" + hostname
}

func LogDebug(message string) {
	// if configuration.GetOptionsDebug() {
	log.Println(message)
	// }
}

func LogFatalDebug(message string) {
	// if configuration.GetOptionsDebug() {
	log.Fatal(message)
	// }
}

func LogFatalDebugError(message string, err error) {
	// if configuration.GetOptionsDebug() {
	log.Fatal(message, err)
	// }
}
