package vault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

// przechowywanie secretu w vault
func StoreSecret(key string, value string) error {
	// tworzenie klienta
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	// ustawienie tokena vault
	client.SetToken("root")
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
