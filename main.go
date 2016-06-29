package main

import (
    "github.com/Sirupsen/logrus"
)

var (
    log = logrus.New()
)

func init() {
    log.Formatter = new(logrus.TextFormatter)
    log.Level     = logrus.DebugLevel
}

func main() {
    log.Info("Goon")
}
