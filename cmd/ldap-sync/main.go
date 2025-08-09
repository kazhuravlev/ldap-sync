package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:  "ldap-sync",
		Usage: "Sync ldap group hierarchy from Yaml to LDAP server",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
