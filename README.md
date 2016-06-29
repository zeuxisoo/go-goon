# Go Goon

a simple ssh execute commands tools

# Build

Install packages

    make env

Build file

    make build

# Run commands

Auth method for key file

    ./go-goon -a keyfile -c ./conf/app.ini -r "ping -c 4 -t 15 hk.yahoo.com"

Auth method for password

    ./go-goon -a password -c ./conf/app.ini -r "ping -c 4 -t 15 hk.yahoo.com"

Show help

    ./go-goon -h
