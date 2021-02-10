package cli

import (
	"github.com/foryouandyourcustomers/kvutil/internal/kvclient"
)

var (
	kvClient kvclient.Client

	// Collection contains alli api commands from cli package
	Collection = append(SecretsCli)

	//
	// moved before function and global flags to
	// subcommands to ensure we dont have a freakyishly
	// odd command line like `kvutil --vault vaulrname secrets get --secret abc`
	//

	// // GlobalFlags contains the definition of all global parameters
	// GlobalFlags = []ucli.Flag{
	// 	// &ucli.StringFlag{
	// 	// 	Name:        "vault",
	// 	// 	Usage:       "Name of the keyvault",
	// 	// 	Required:    true,
	// 	// 	EnvVars:     []string{"VAULT"},
	// 	// 	Destination: &kvClient.Vault,
	// 	// },
	// }
	// // BeforeFunction is executed prior to execution of any subcommand
	// // BeforeFunction = func(c *ucli.Context) error {
	// // 	kvClient.Ctx = context.Background()
	// // 	err := kvClient.Authenticate()
	// // 	if err != nil {
	// // 		return err
	// // 	}
	// // 	return nil
	// }
)
