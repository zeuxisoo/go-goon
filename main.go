package main

import (
    "github.com/Sirupsen/logrus"
    "gopkg.in/ini.v1"
)

var (
    log        = logrus.New()
    configFile = "conf/app.ini"
)

func init() {
    log.Formatter = new(logrus.TextFormatter)
    log.Level     = logrus.DebugLevel
}

func main() {
    log.Info("Goon")

    cfg, err := ini.Load(configFile)

    if err != nil {
        log.WithField("configFile", configFile).Fatal("Failed to parse config file")
    }

    serverSection := cfg.Section("server")

    log.Info(serverSection.Key("HOST"))
    log.Info(serverSection.Key("PORT"))
    log.Info(serverSection.Key("PRIVATE_KEY"))
}
