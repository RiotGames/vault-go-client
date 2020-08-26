package main

import (
	"log"
	"os"

	"github.com/RiotGames/vault-go-client/vaultClient"
	"github.com/k0kubun/pp"
)

type Secret struct {
	Hello string `json:"hello"`
	vaultClient.SecretMetadata
}

func main() {
	vaultToken := os.Getenv("VAULT_TOKEN")
	secretMountPath := os.Getenv("VAULT_SECRET_MOUNT_PATH") // "example/secrets"
	secretPath := os.Getenv("VAULT_SECRET_PATH")            // "token_example"

	client, err := vaultClient.NewClient(vaultClient.DefaultConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := client.Auth.Token.Login(vaultClient.TokenOptions{
		Token: vaultToken}); err != nil {
		log.Fatal(err.Error())
	}

	secretMap := map[string]interface{}{
		"hello": "world",
	}

	if _, err = client.KV2.Put(vaultClient.KV2PutOptions{
		MountPath:  secretMountPath,
		SecretPath: secretPath,
		Secrets:    secretMap,
	}); err != nil {
		log.Fatal(err.Error())
	}

	secret, err := client.KV2.Get(vaultClient.KV2GetOptions{
		MountPath:  secretMountPath,
		SecretPath: secretPath,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	pp.Println(secret)

	secrets := &Secret{}

	if _, err = client.KV2.Get(vaultClient.KV2GetOptions{
		MountPath:     secretMountPath,
		SecretPath:    secretPath,
		UnmarshalInto: secrets,
	}); err != nil {
		log.Fatal(err.Error())
	}
	pp.Println(secrets)
}
