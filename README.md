# ibtcli
Common cli commands to automate repeatitive tasks

# Getting started
Inspired by https://github.com/urfave/cli/. 

1. To setup run the following:
```
    go mod init ibtcli
    go mod tidy
    go install ibtcli.go
```

2. You should be then able to run the cli
```
user@Macbook-Pro workspace % ibtcli

NAME:
   CLI interface for common technical tasks - AWS, Azure, others

USAGE:
   ibtcli [global options] command [command options] [arguments...]

COMMANDS:
   aws, aws  aws cli handy commands
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

```