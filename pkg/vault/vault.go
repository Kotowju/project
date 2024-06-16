package vault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

// StoreSecret stores a secret in Vault.
func StoreSecret(key string, value string) error {
	// Create a new Vault client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	// Set Vault token (replace with proper authentication method in production)
	client.SetToken("root")

	// Write the secret to Vault
	_, err = client.Logical().Write(fmt.Sprintf("secret/data/%s", key), map[string]interface{}{
		"data": map[string]interface{}{
			"value": value,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
