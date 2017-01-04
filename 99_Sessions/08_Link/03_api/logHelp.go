package main

import (
	log1 "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

func NewLogger() *log1.Logger {
	if logger != nil {
		return logger
	}

	logger = log1.New()
	logger.Formatter = new(log1.JSONFormatter)
	logger.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		log1.InfoLevel:  logdir + "info.log",
		log1.ErrorLevel: logdir + "error.log",
	}))
	return logger
}
