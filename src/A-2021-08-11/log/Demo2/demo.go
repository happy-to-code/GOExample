package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	// log.SetLevel(log.DebugLevel)

	log.Debug("debug=====>Useful debugging information.")
	log.Info("info=====>Something noteworthy happened!")
	log.Warn("warn=====>You should probably take a look at this.")
	log.Error("error=====>Something failed but I'm not quitting.")
}
