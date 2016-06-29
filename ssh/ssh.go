package ssh

import (
    "os"
    "io/ioutil"
    "strconv"
    "time"

    "github.com/Sirupsen/logrus"
    "golang.org/x/crypto/ssh"
)

type Config struct {
    Host        string
    Port        int
    User        string
    PrivateKey  string
}

type Ssh struct {
    config  Config
    logger  *logrus.Logger
}

func NewSsh(config Config) (s *Ssh) {
    return &Ssh{
        config: config,
    }
}

func (s *Ssh) SetLogger(logger *logrus.Logger) {
    s.logger = logger
}

func (s *Ssh) RunCommand(command string) (result string) {
    privateKeyBytes := s.readPrivateKey()
    sshClient       := s.createSshClient(privateKeyBytes)
    sshSession      := s.createSshSession(sshClient)

    buffer, err := sshSession.CombinedOutput(command)

    if err != nil {
        s.logger.WithFields(logrus.Fields{
            "error"  : err.Error(),
            "command": command,
        }).Error("Failed to execute command")
    }

    sshSession.Close()
    sshClient.Close()

    return string(buffer)
}

func (s *Ssh) readPrivateKey() (privateKeyBytes []byte) {
    file, err := os.Open(s.config.PrivateKey)

    if err != nil {
        s.logger.WithField("error", err.Error()).Error("Failed to read private key file")
    }

    defer file.Close()

    buffer, _ := ioutil.ReadAll(file)

    return buffer
}

func (s *Ssh) createSshClient(privateKeyBytes []byte) (client *ssh.Client) {
    signer, _ := ssh.ParsePrivateKey(privateKeyBytes)

    clientConfig := &ssh.ClientConfig{
        User: s.config.User,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        },
        Timeout: 30 * time.Second,
    }

    address := s.config.Host + ":" + strconv.Itoa(s.config.Port)

    client, err := ssh.Dial("tcp", address, clientConfig)

    if err != nil {
        s.logger.WithFields(logrus.Fields{
            "error"  : err.Error(),
            "address": address,
        }).Error("Failed to create client")
    }

    return client
}

func (s *Ssh) createSshSession(sshClient *ssh.Client) (session *ssh.Session) {
    session, err := sshClient.NewSession()

    if err != nil {
        s.logger.WithField("error", err.Error()).Error("Failed to create session")
    }

    return session
}
