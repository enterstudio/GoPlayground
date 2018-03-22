package main

import (
	log1 "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

// NewLogger This function creates the proper persistent "gLogger"
//  This is used in the Main Code to attached this logger globally to the Server
// Note: this influences the Package level variable "gLogger"
func NewLogger() *log1.Logger {
	if gLogger != nil {
		return gLogger
	}

	gLogger = log1.New()
	gLogger.Formatter = new(log1.TextFormatter)
	gLogger.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		log1.InfoLevel:  logdir + "info.log",
		log1.ErrorLevel: logdir + "error.log",
	}))
	return gLogger
}
