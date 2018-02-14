package main

import (
	"documentAnalyzer/pepper"
)

func main() {
	log := pepper.New()
	log.Info("Just is an info message")
	log.Error("This is an error message")
	log.Debug("This is a debug message")
}
