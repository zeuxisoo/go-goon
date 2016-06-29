package main

import (
    "fmt"
    "flag"
    "os"

    "github.com/Sirupsen/logrus"
    "github.com/fatih/color"

    "github.com/zeuxisoo/go-goon/setting"
    "github.com/zeuxisoo/go-goon/ssh"
    "github.com/zeuxisoo/go-goon/ssh/authenticator"
    "github.com/zeuxisoo/go-goon/ssh/contract"
)

var (
    version = "0.1.0"

    log     = logrus.New()
)

func usage() {
    const usage = `Goon: a simple ssh execute commands tools
Usage:
    go-goon [-a] [-c CONFIG_FILE] [-r COMMAND]
    go-goon -h | --help
Options:
    -a,             Authorization method
    -c,             Configure file path
    -r,             Command for execute
    -h, --help      Output help information
`

    fmt.Printf(usage)
    os.Exit(0)
}

func init() {
    log.Formatter = new(logrus.TextFormatter)
    log.Level     = logrus.DebugLevel
}

func main() {
    var authMethod, configFile, command string
    var help bool

    argument := NewArgument()

    flag.StringVar(&authMethod, "a",    "",    "Authorization method")
    flag.StringVar(&configFile, "c",    "",    "Configure file path")
    flag.StringVar(&command,    "r",    "",    "Command for execute")
    flag.BoolVar(&help,         "h",    false, "Show help message")
    flag.BoolVar(&help,         "help", false, "Show help message")

    flag.Usage = usage

    flag.Parse()

    if help {
        usage()
    }

    argument.AuthMethod = authMethod
    argument.ConfigFile = configFile
    argument.Command    = command

    if _, err := argument.Check(); err != nil {
        color.Red("Argument error: %s", err)
    }else{
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
        log.Info(magenta("password   : ", setting.Values.Server.Password))
        log.Info(magenta("private key: ", setting.Values.Server.PrivateKey))

        var auth contract.Authenticator

        if authMethod == "keyfile" {
            // Create ssh authenticator for key file
            auth = new(authenticator.KeyFile)
            auth.SetLogger(log)
            auth.SetSettingValues(setting.Values)
        }else{
            // Create ssh authenticator for password
            auth = new(authenticator.Password)
            auth.SetLogger(log)
            auth.SetSettingValues(setting.Values)
        }

        // Create ssh agent
        sshAgent := ssh.NewSsh(ssh.Config{
            Host      : setting.Values.Server.Host,
            Port      : setting.Values.Server.Port,
            User      : setting.Values.Server.User,
            Password  : setting.Values.Server.Password,
            PrivateKey: setting.Values.Server.PrivateKey,
        })
        sshAgent.SetLogger(log)
        sshAgent.SetAuthenticator(auth)

        // Run command using ssh agent
        result := sshAgent.RunCommand(command)

        // Display result
        log.Info(yellow("Result"))
        color.White("\n%s", result)
    }
}
