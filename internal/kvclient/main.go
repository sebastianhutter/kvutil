package kvclient

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest/azure"
)

type Client struct {
	Ctx          context.Context
	Vault        string
	cli          keyvault.BaseClient
	VaultBaseURL string
}

func (c *Client) Authenticate() error {
	a, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		a, err = auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return err
		}
	}

	c.cli.Authorizer = a
	c.VaultBaseURL = fmt.Sprintf("https://%s.%s", c.Vault, azure.PublicCloud.KeyVaultDNSSuffix)
	return nil
}

// GetSecretBundle retrieve specified secret from keyvault
func (c *Client) GetSecretBundle(s string) (keyvault.SecretBundle, error) {
	sb, err := c.cli.GetSecret(c.Ctx, c.VaultBaseURL, s, "")
	if err != nil {
		return sb, err
	}
	return sb, nil
}

// GetSecretValue get decrypted secret value from bundle
func (c *Client) GetSecretValue(sb keyvault.SecretBundle) string {
	return *sb.Value
}

// SetSecretValue set the secret value to the specified string
func (c *Client) SetSecretValue(s string, v string) error {
	p := &keyvault.SecretSetParameters{Value: &v}
	_, err := c.cli.SetSecret(c.Ctx, c.VaultBaseURL, s, *p)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSecret - delete a secret from the vault, if purge is set to true, purge it.
func (c *Client) DeleteSecret(s string, p bool) error {
	dsb, err := c.cli.DeleteSecret(c.Ctx, c.VaultBaseURL, s)
	// if no secret was found / deleted check for the deletedsecret
	// so we can purge it if purge is true
	if dsb.StatusCode != 404 {
		dsb, err = c.cli.GetDeletedSecret(c.Ctx, c.VaultBaseURL, s)
	}
	// if we still receive a 404 we can safely assume the secret either never
	// existed or was deleted and purged already
	if err != nil && dsb.StatusCode != 404 {
		return err
	}

	if p {
		_, err := c.cli.PurgeDeletedSecret(c.Ctx, c.VaultBaseURL, s)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSecrets - return all secrets names
func (c *Client) GetSecrets() ([]keyvault.SecretItem, error) {
	sl := make([]keyvault.SecretItem, 0)
	it, err := c.cli.GetSecretsComplete(c.Ctx, c.VaultBaseURL, nil)
	if err != nil {
		return sl, err
	}

	for it.NotDone() {
		sl = append(sl, it.Value())
		it.Next()

	}
	return sl, nil
}
