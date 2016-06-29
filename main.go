package main

import (
    "fmt"

    "github.com/Sirupsen/logrus"
    "github.com/fatih/color"

    "github.com/zeuxisoo/goon/setting"
    "github.com/zeuxisoo/goon/ssh"
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
    cyan    := color.New(color.FgCyan).SprintFunc()
    magenta := color.New(color.FgMagenta).SprintFunc()
    yellow  := color.New(color.FgYellow).SprintFunc()

    log.Info(cyan("Goon"))

    // Load setting
    setting := setting.NewSetting(log)
    setting.Load(configFile)

    log.Info(magenta("host       : ", setting.Values.Server.Host))
    log.Info(magenta("port       : ", setting.Values.Server.Port))
    log.Info(magenta("user       : ", setting.Values.Server.User))
    log.Info(magenta("private key: ", setting.Values.Server.PrivateKey))

    // Init ssh client
    ssh := ssh.NewSsh(ssh.Config{
        Host      : setting.Values.Server.Host,
        Port      : setting.Values.Server.Port,
        User      : setting.Values.Server.User,
        PrivateKey: setting.Values.Server.PrivateKey,
    })
    ssh.SetLogger(log)

    // Run command in ssh client
    result := ssh.RunCommand("ping -c 4 -t 15 hk.yahoo.com")

    // Display result
    log.Info(yellow("Result"))

    fmt.Println("\n" + result)
}
