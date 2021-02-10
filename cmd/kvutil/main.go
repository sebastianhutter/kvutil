package main

import (
	"log"
	"os"

	"github.com/foryouandyourcustomers/kvutil/internal/cli"
	ucli "github.com/urfave/cli/v2"
)

// func init() {

// 	var v string
// 	var s string

// 	flag.StringVar(&v, "v", "vault", "Name of the keyvault (env var: VAULT)")
// 	flag.StringVar(&s, "s", "secret", "Name of the secret to retrieve (env var: SECRET)")
// 	flag.Parse()

// 	if os.Getenv("SECRET") != "" {
// 		s = os.Getenv("SECRET")
// 	}
// 	if os.Getenv("VAULT") != "" {
// 		v = os.Getenv("VAULT")
// 	}

// 	if (v == "") || (s == "") {
// 		flag.PrintDefaults()
// 		os.Exit(1)
// 	}

// 	err := kvClient.Authenticate()
// 	if err != nil {
// 		log.Fatalf("Unable to authenticate: %s", err)
// 	}
// 	err = kvClient.SetVault(v)
// 	if err != nil {
// 		log.Fatalf("Unable to find vault or access denied: %s", err)
// 	}

// }

// func main() {

// 	// kvclient

// 	// kvClient.ctx = context.Background()
// 	// kvClient.authenticate(vaultName)

// 	// s, err := kvClient.fetchSecret(secretName)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Print(s)
// }

func main() {
	app := &ucli.App{
		Name:  "kvutil",
		Usage: "Helper functions to manage keyvaults",
		//Flags:    cli.GlobalFlags,
		//Before:   cli.BeforeFunction,
		Commands: cli.Collection,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
