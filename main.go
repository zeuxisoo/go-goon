package main

import (
    "github.com/Sirupsen/logrus"

    "github.com/zeuxisoo/goon/setting"
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

    setting := setting.NewSetting(log)
    setting.Load(configFile)

    log.Info(setting.Values.Server.Host)
    log.Info(setting.Values.Server.Port)
    log.Info(setting.Values.Server.PrivateKey)
}
