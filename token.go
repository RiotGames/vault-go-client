package vaultClient

import (
	vault "github.com/hashicorp/vault/api"
)

type token struct {
	client *vault.Client
}

type TokenOptions struct {
	Token string
}

func (a *token) Login(options TokenOptions) (*vault.Secret, error) {
	a.client.SetToken(options.Token)
	return a.client.Auth().Token().LookupSelf()
}
