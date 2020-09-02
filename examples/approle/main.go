package main

import (
	"log"
	"os"

	"github.com/k0kubun/pp"

	vault "github.com/riotgames/vault-go-client"
)

type Secret struct {
	Hello string `json:"hello"`
	vault.SecretMetadata
}

func main() {
	appRoleID := os.Getenv("VAULT_APPROLE_ID")              // "aeda62c7-22c0-b22a-677e-129a41a5c73c"
	appRoleSecretID := os.Getenv("VAULT_APPROLE_SECRET_ID") // "24904192-b26c-fd16-8be6-7453daa8e767"
	appRoleAuthPath := os.Getenv("VAULT_APPROLE_AUTH_PATH") // "example/approle"
	secretMountPath := os.Getenv("VAULT_SECRET_MOUNT_PATH") // "example/secrets"
	secretPath := os.Getenv("VAULT_SECRET_PATH")            // "approle_example"

	client, err := vault.NewClient(vault.DefaultConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err = client.Auth.AppRole.Login(vault.AppRoleLoginOptions{
		RoleID:    appRoleID,
		SecretID:  appRoleSecretID,
		MountPath: appRoleAuthPath}); err != nil {
		log.Fatal(err.Error())
	}

	secretMap := map[string]interface{}{
		"hello": "world",
	}

	if _, err = client.KV2.Put(vault.KV2PutOptions{
		MountPath:  secretMountPath,
		SecretPath: secretPath,
		Secrets:    secretMap,
	}); err != nil {
		log.Fatal(err.Error())
	}

	secret, err := client.KV2.Get(vault.KV2GetOptions{
		MountPath:  secretMountPath,
		SecretPath: secretPath,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	pp.Println(secret)

	secrets := &Secret{}

	if _, err = client.KV2.Get(vault.KV2GetOptions{
		MountPath:     secretMountPath,
		SecretPath:    secretPath,
		UnmarshalInto: secrets,
	}); err != nil {
		log.Fatal(err.Error())
	}
	pp.Println(secrets)
}
