package main

import (
    "errors"
)

type Argument struct {
    AuthMethod  string
    ConfigFile  string
    Command     string
}

func NewArgument() (c *Argument) {
    return &Argument{}
}

func (c *Argument) Check() (string, error) {
    if c.AuthMethod == "" {
        return "", errors.New("The auth method should be keyfile or password only")
    }

    if c.ConfigFile == "" {
        return "", errors.New("The config file path cannot be empty")
    }

    if c.Command == "" {
        return "", errors.New("The command must be provided")
    }

    return "", nil
}
