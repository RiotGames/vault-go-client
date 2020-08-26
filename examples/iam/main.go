package main

import (
	"log"
	"os"

	"github.com/k0kubun/pp"

	"github.com/RiotGames/vault-go-client/vaultClient"
)

type Secret struct {
	Hello string `json:"hello"`
	vaultClient.SecretMetadata
}

func main() {
	vaultRoleName := os.Getenv("VAULT_ROLE_NAME")           // "example_role"
	awsAuthPath := os.Getenv("VAULT_AWS_IAM_PATH")          // "example/aws"
	secretMountPath := os.Getenv("VAULT_SECRET_MOUNT_PATH") // "example/secrets"
	secretPath := os.Getenv("VAULT_SECRET_PATH")            // "iam_example"

	client, err := vaultClient.NewClient(vaultClient.DefaultConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := client.Auth.IAM.Login(vaultClient.IAMLoginOptions{
		Role:      vaultRoleName,
		MountPath: awsAuthPath}); err != nil {
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
