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
	vaultRoleName := os.Getenv("VAULT_ROLE_NAME")           // "example_role"
	awsAuthPath := os.Getenv("VAULT_AWS_IAM_PATH")          // "example/aws"
	secretMountPath := os.Getenv("VAULT_SECRET_MOUNT_PATH") // "example/secrets"
	secretPath := os.Getenv("VAULT_SECRET_PATH")            // "iam_example"

	client, err := vault.NewClient(vault.DefaultConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := client.Auth.IAM.Login(vault.IAMLoginOptions{
		Role:      vaultRoleName,
		MountPath: awsAuthPath}); err != nil {
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
