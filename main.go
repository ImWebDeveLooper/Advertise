package main

import (
	"agahi/internal/platform/application"
	log "github.com/sirupsen/logrus"
)

func init() {
	// set the log level to Debug level for showing log just for devs.
	log.SetLevel(log.DebugLevel)
}

func main() {
	//create a new instance of Application
	application.NewApp()

}
