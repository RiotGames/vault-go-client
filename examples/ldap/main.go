package main

import (
	"fmt"
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
	username := os.Getenv("LDAP_USERNAME")
	password := os.Getenv("LDAP_PASSWORD")
	secretMountPath := os.Getenv("VAULT_SECRET_MOUNT_PATH") // "example/secrets"
	secretPath := os.Getenv("VAULT_SECRET_PATH")            // "ldap_example"

	client, err := vaultClient.NewClient(vaultClient.DefaultConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, err := client.Auth.LDAP.Login(vaultClient.LDAPLoginOptions{
		Username: username,
		Password: password}); err != nil {
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
		SecretPath: secretPath})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pp.Println(secret)

	secrets := &Secret{}
	if _, err := client.KV2.Get(vaultClient.KV2GetOptions{
		MountPath:     secretMountPath,
		SecretPath:    secretPath,
		UnmarshalInto: secrets}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pp.Println(secrets)
}
