package vaultClient

import (
	"errors"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

type appRole struct {
	client *vault.Client
}

type AppRoleLoginOptions struct {
	RoleID    string
	SecretID  string
	MountPath string
}

func (a *appRole) Login(options AppRoleLoginOptions) (*vault.Secret, error) {
	authSecret, err := a.appRoleLogin(options)

	if err != nil {
		return nil, err
	}

	if authSecret.Auth == nil {
		return nil, errors.New("Vault AppRole Auth returned nil")
	}

	a.client.SetToken(authSecret.Auth.ClientToken)
	return authSecret, nil
}

func (a *appRole) appRoleLogin(options AppRoleLoginOptions) (*vault.Secret, error) {
	appRoleCreds := map[string]interface{}{
		"role_id":   options.RoleID,
		"secret_id": options.SecretID,
	}

	authPath := "auth/approle/login"
	if options.MountPath != "" {
		authPath = "auth/" + strings.Trim(options.MountPath, "/") + "/login"
	}

	authSecret, err := a.client.Logical().Write(authPath, appRoleCreds)

	if err != nil {
		return nil, err
	}

	if authSecret == nil {
		return nil, errors.New("empty response from Vault AppRole")
	}

	return authSecret, nil
}