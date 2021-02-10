package cli

import (
	"context"
	"path"

	ucli "github.com/urfave/cli/v2"
)

var (

	// SecretsCli contains operations and paramteres for the secret operations
	SecretsCli = []*ucli.Command{
		{
			Name:  "secrets",
			Usage: "Manage keyvault secrets",
			// Flags: []ucli.Flag{
			// 	&ucli.StringFlag{
			// 		Name:        "apiid",
			// 		Usage:       "name (api id) of the api to deploy",
			// 		Required:    true,
			// 		EnvVars:     []string{"APIID"},
			// 		Destination: &apiDef.APIID,
			// 	},
			// },
			Subcommands: []*ucli.Command{
				{
					Name:  "get",
					Usage: "Retrieve latest secret value",
					Action: func(c *ucli.Context) error {
						sb, err := kvClient.GetSecretBundle(c.String("secret"))
						if err != nil {
							return ucli.Exit(err, 1)
						}
						println(kvClient.GetSecretValue(sb))
						return nil
					},
					Before: func(c *ucli.Context) error {
						kvClient.Ctx = context.Background()
						err := kvClient.Authenticate()
						if err != nil {
							return ucli.Exit(err, 1)
						}
						return nil
					},
					Flags: []ucli.Flag{
						&ucli.StringFlag{
							Name:        "vault",
							Usage:       "Name of the keyvault",
							Required:    true,
							EnvVars:     []string{"VAULT"},
							Destination: &kvClient.Vault,
						},
						&ucli.StringFlag{
							Name:     "secret",
							Usage:    "Name of the secret",
							Required: true,
							EnvVars:  []string{"SECRET"},
						},
					},
				},
				{
					Name:  "set",
					Usage: "Set a secret value",
					Action: func(c *ucli.Context) error {
						err := kvClient.SetSecretValue(c.String("secret"), c.String("value"))
						if err != nil {
							return ucli.Exit(err, 1)
						}
						return nil
					},
					Before: func(c *ucli.Context) error {
						kvClient.Ctx = context.Background()
						err := kvClient.Authenticate()
						if err != nil {
							return ucli.Exit(err, 1)
						}
						return nil
					},
					Flags: []ucli.Flag{
						&ucli.StringFlag{
							Name:        "vault",
							Usage:       "Name of the keyvault",
							Required:    true,
							EnvVars:     []string{"VAULT"},
							Destination: &kvClient.Vault,
						},
						&ucli.StringFlag{
							Name:     "secret",
							Usage:    "Name of the secret",
							Required: true,
							EnvVars:  []string{"SECRET"},
						},
						&ucli.StringFlag{
							Name:     "value",
							Usage:    "Value of the secret",
							Required: true,
							EnvVars:  []string{"VALUE"},
						},
					},
				},
				{
					Name:  "rm",
					Usage: "Delete a secret",
					Action: func(c *ucli.Context) error {
						err := kvClient.DeleteSecret(c.String("secret"), c.Bool("purge"))
						if err != nil {
							return ucli.Exit(err, 1)
						}
						return nil
					},
					Before: func(c *ucli.Context) error {
						kvClient.Ctx = context.Background()
						err := kvClient.Authenticate()
						if err != nil {
							return ucli.Exit(err, 1)
						}
						return nil
					},
					Flags: []ucli.Flag{
						&ucli.StringFlag{
							Name:        "vault",
							Usage:       "Name of the keyvault",
							Required:    true,
							EnvVars:     []string{"VAULT"},
							Destination: &kvClient.Vault,
						},
						&ucli.StringFlag{
							Name:     "secret",
							Usage:    "Name of the secret",
							Required: true,
							EnvVars:  []string{"SECRET"},
						},
						&ucli.BoolFlag{
							Name:     "purge",
							Usage:    "purge the secret?",
							Required: false,
							Value:    false,
							EnvVars:  []string{"PURGE"},
						},
					},
				},
				{
					Name:  "ls",
					Usage: "List secrets",
					Action: func(c *ucli.Context) error {
						sl, err := kvClient.GetSecrets()
						if err != nil {
							return ucli.Exit(err, 1)
						}
						for _, s := range sl {
							println(path.Base(*s.ID))
						}
						return nil
					},
					Before: func(c *ucli.Context) error {
						kvClient.Ctx = context.Background()
						err := kvClient.Authenticate()
						if err != nil {
							return ucli.Exit(err, 1)
						}
						return nil
					},
					Flags: []ucli.Flag{
						&ucli.StringFlag{
							Name:        "vault",
							Usage:       "Name of the keyvault",
							Required:    true,
							EnvVars:     []string{"VAULT"},
							Destination: &kvClient.Vault,
						},
					},
				},
			},
		},
	}
)
