package contract

import (
    "github.com/zeuxisoo/go-goon/setting"

    "github.com/Sirupsen/logrus"
    "golang.org/x/crypto/ssh"
)

type Authenticator interface {
    SetLogger(*logrus.Logger)
    SetSettingValues(setting.SettingValues)
    SshClient() *ssh.Client
}
