package setting

import (
    "github.com/Sirupsen/logrus"
    "gopkg.in/ini.v1"
)

type Setting struct {
    logger  *logrus.Logger

    Config  *ini.File

    Values  SettingValues
}

type SettingValues struct {
    Server struct {
        Host        string
        Port        int
        PrivateKey  string
    }
}

func NewSetting(logger *logrus.Logger) (s *Setting) {
    return &Setting{
        logger: logger,
    }
}

func (s *Setting) Load(configFile string) {
    config, err := ini.Load(configFile)

    if err != nil {
        s.logger.WithField("configFile", configFile).Fatal("Failed to parse config file")
    }

    s.Config = config

    //
    serverSection := config.Section("server")

    s.Values.Server.Host       = serverSection.Key("HOST").MustString("127.0.0.1")
    s.Values.Server.Port       = serverSection.Key("PORT").MustInt(22)
    s.Values.Server.PrivateKey = serverSection.Key("PRIVATE_KEY").MustString("")
}