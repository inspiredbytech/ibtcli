package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

type AwsEksCommand struct {
}

func (cmd AwsEksCommand) Command() *cli.Command {
	return &cli.Command{
		Name:  "eks-update-profile",
		Usage: "Refresh EKS",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "profile", Aliases: []string{"p"}, Required: true, EnvVars: []string{"AWS_PROFILE"}},
			&cli.StringFlag{Name: "cluster", Aliases: []string{"c"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("SSO Login: ", c.Args().First())

			//aws  login --profile $profile
			awsProfile := c.String("profile")
			eksCluster := c.String("cluster")

			cmd := exec.Command("aws", "sso", "login", "--profile", awsProfile)
			err := cmd.Run()
			//stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

			cmd = exec.Command("aws", "eks", "--profile", awsProfile, "update-kubeconfig", "--name", eksCluster)

			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(stdout))

			return nil
		},
	}
}

type AwsSsoCommand struct {
}

func (cmd AwsSsoCommand) Command() *cli.Command {
	return &cli.Command{
		Name:  "sso-login",
		Usage: "login using sso",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "profile", Aliases: []string{"p"}, Required: true, EnvVars: []string{"AWS_PROFILE"}},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("SSO Login: ", c.Args().First())
			cmd := exec.Command("aws", "sso", "login", "--profile", c.String("profile"))
			err := cmd.Run()
			//stdout, err := cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}
}

func main() {
	/*addCommand := cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task to the list",
		Action: func(c *cli.Context) error {
			fmt.Println("added task: ", c.Args().First())
			return nil
		},
	}*/

	awsEks := AwsEksCommand{}
	awsSso := AwsSsoCommand{}

	app := &cli.App{
		Name:  "CLI interface for common technical tasks",
		Usage: "AWS, Azure, others",
		Commands: []*cli.Command{
			{
				Name:    "aws",
				Aliases: []string{"aws"},
				Usage:   "aws cli handy commands",
				Subcommands: []*cli.Command{
					awsSso.Command(),
					awsEks.Command(),
				},
			},
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
