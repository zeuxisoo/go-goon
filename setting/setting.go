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
        User        string
        Password    string
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

    s.Values.Server.Host       = serverSection.Key("HOST").MustString("")
    s.Values.Server.Port       = serverSection.Key("PORT").MustInt(22)
    s.Values.Server.User       = serverSection.Key("USER").MustString("")
    s.Values.Server.Password   = serverSection.Key("PASSWORD").MustString("")
    s.Values.Server.PrivateKey = serverSection.Key("PRIVATE_KEY").MustString("")
}
