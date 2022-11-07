package vault

import (
	hashivault "github.com/hashicorp/vault/api"
)

type Token struct {
	client *hashivault.Client
}

type TokenOptions struct {
	Token string
}

func (a *Token) Login(options TokenOptions) (*hashivault.Secret, error) {
	a.client.SetToken(options.Token)
	return a.client.Auth().Token().LookupSelf()
}

func (a *Token) Renew(increment int) (*hashivault.Secret, error) {
	return a.client.Auth().Token().RenewSelf(increment)
}
