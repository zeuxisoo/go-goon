package ssh

import (
    "github.com/zeuxisoo/goon/ssh/contract"

    "github.com/Sirupsen/logrus"
    "golang.org/x/crypto/ssh"
)

type Config struct {
    Host        string
    Port        int
    User        string
    Password    string
    PrivateKey  string
}

type Ssh struct {
    config          Config
    logger          *logrus.Logger
    authenticator   contract.Authenticator
}

func NewSsh(config Config) (s *Ssh) {
    return &Ssh{
        config: config,
    }
}

func (s *Ssh) SetLogger(logger *logrus.Logger) {
    s.logger = logger
}

func (s *Ssh) SetAuthenticator(authenticator contract.Authenticator) {
    s.authenticator = authenticator
}

func (s *Ssh) RunCommand(command string) (result string) {
    sshClient  := s.authenticator.SshClient()
    sshSession := s.createSshSession(sshClient)

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

func (s *Ssh) createSshSession(sshClient *ssh.Client) (session *ssh.Session) {
    session, err := sshClient.NewSession()

    if err != nil {
        s.logger.WithField("error", err.Error()).Error("Failed to create session")
    }

    return session
}
