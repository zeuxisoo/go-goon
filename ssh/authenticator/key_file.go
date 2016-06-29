package authenticator

import (
    "os"
    "io/ioutil"
    "time"
    "strconv"

    "github.com/zeuxisoo/goon/setting"

    "github.com/Sirupsen/logrus"
    "golang.org/x/crypto/ssh"
)

type KeyFile struct {
    logger          *logrus.Logger
    settingValues   setting.SettingValues
}

// Implement
func (k *KeyFile) SetLogger(logger *logrus.Logger) {
    k.logger = logger
}

func (k *KeyFile) SetSettingValues(settingValues setting.SettingValues) {
    k.settingValues = settingValues
}

func (k *KeyFile) SshClient() (client *ssh.Client) {
    privateKeyBytes := k.readPrivateKey()

    signer, _ := ssh.ParsePrivateKey(privateKeyBytes)

    clientConfig := &ssh.ClientConfig{
        User: k.settingValues.Server.User,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        },
        Timeout: 30 * time.Second,
    }

    address := k.settingValues.Server.Host + ":" + strconv.Itoa(k.settingValues.Server.Port)

    client, err := ssh.Dial("tcp", address, clientConfig)

    if err != nil {
        k.logger.WithFields(logrus.Fields{
            "error"  : err.Error(),
            "address": address,
        }).Error("Failed to create client")
    }

    return client
}

// Private
func (k *KeyFile) readPrivateKey() (privateKeyBytes []byte) {
    file, err := os.Open(k.settingValues.Server.PrivateKey)

    if err != nil {
        k.logger.WithField("error", err.Error()).Error("Failed to read private key file")
    }

    defer file.Close()

    buffer, _ := ioutil.ReadAll(file)

    return buffer
}
