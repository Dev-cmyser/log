package main

import (
	log "github.com/Dev-Cmyser/logger"
)

func main() {

	log.SetLevel(log.Level.Debug)
	log.Warn("Hello, World!")
	// log.Trace("Hello, Trace!")
	// log.Info("Hello, World!")
	// log.Debug("Debugging message")
	// log.Warn("Warn")
	// log.Error("Error")
	// showLogsAgain()
	// test.ShowLogsAgain()

}
func showLogsAgain() {
	log.Trace("Hello, Trace!")
	log.Info("Hello, World!")
	log.Debug("Debugging message")
	log.Warn("Warn")
	log.Error("Error")

}