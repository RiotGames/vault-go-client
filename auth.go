package vault

import hashivault "github.com/hashicorp/vault/api"

type auth struct {
	LDAP    *ldap
	IAM     *iam
	Token   *token
	AppRole *appRole
}

func NewAuth(client *hashivault.Client) *auth {
	return &auth{
		LDAP: &ldap{
			client: client},
		IAM: &iam{
			client: client},
		Token: &token{
			client: client},
		AppRole: &appRole{
			client: client}}
}
