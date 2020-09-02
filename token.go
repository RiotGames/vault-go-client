package vault

import (
	hashivault "github.com/hashicorp/vault/api"
)

type token struct {
	client *hashivault.Client
}

type TokenOptions struct {
	Token string
}

func (a *token) Login(options TokenOptions) (*hashivault.Secret, error) {
	a.client.SetToken(options.Token)
	return a.client.Auth().Token().LookupSelf()
}
