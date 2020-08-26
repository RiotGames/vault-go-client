package vaultClient

import vault "github.com/hashicorp/vault/api"

type auth struct {
	LDAP    *ldap
	IAM     *iam
	Token   *token
	AppRole *appRole
}

func NewAuth(vaultClient *vault.Client) *auth {
	return &auth{
		LDAP: &ldap{
			client: vaultClient},
		IAM: &iam{
			client: vaultClient},
		Token: &token{
			client: vaultClient},
		AppRole: &appRole{
			client: vaultClient}}
}
