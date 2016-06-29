package authenticator

import (
    "time"
    "strconv"

    "github.com/zeuxisoo/goon/setting"

    "github.com/Sirupsen/logrus"
    "golang.org/x/crypto/ssh"
)

type Password struct {
    logger          *logrus.Logger
    settingValues   setting.SettingValues
}

// Implement
func (p *Password) SetLogger(logger *logrus.Logger) {
    p.logger = logger
}

func (p *Password) SetSettingValues(settingValues setting.SettingValues) {
    p.settingValues = settingValues
}

func (p *Password) SshClient() (client *ssh.Client) {
    clientConfig := &ssh.ClientConfig{
        User: p.settingValues.Server.User,
        Auth: []ssh.AuthMethod{
            ssh.Password(p.settingValues.Server.Password),
        },
        Timeout: 30 * time.Second,
    }

    address := p.settingValues.Server.Host + ":" + strconv.Itoa(p.settingValues.Server.Port)

    client, err := ssh.Dial("tcp", address, clientConfig)

    if err != nil {
        p.logger.WithFields(logrus.Fields{
            "error"  : err.Error(),
            "address": address,
        }).Error("Failed to create client")
    }

    return client
}
