package vaultClient

import (
	"errors"
	"fmt"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

type ldap struct {
	client *vault.Client
}

type LDAPLoginOptions struct {
	Username  string
	Password  string
	MountPath string
}

func (l *ldap) Login(options LDAPLoginOptions) (*vault.Secret, error) {
	authSecret, err := l.ldapLogin(options)

	if err != nil {
		return nil, err
	}

	if authSecret.Auth == nil {
		return nil, errors.New("Vault LDAP Auth returned nil")
	}

	l.client.SetToken(authSecret.Auth.ClientToken)
	return authSecret, nil
}

func (l *ldap) ldapLogin(options LDAPLoginOptions) (*vault.Secret, error) {
	ldapCreds := map[string]interface{}{
		"password": options.Password,
	}
	pathFormatString := "auth/ldap/login/%s"
	if options.MountPath != "" {
		pathFormatString = "auth/" + strings.Trim(options.MountPath, "/") + "/login/%s"
	}
	normalizedPath := fmt.Sprintf(pathFormatString, options.Username)

	authSecret, err := l.client.Logical().Write(normalizedPath, ldapCreds)

	if err != nil {
		return nil, err
	}

	if authSecret == nil {
		return nil, errors.New("empty response from vault ldap")
	}

	return authSecret, nil
}
