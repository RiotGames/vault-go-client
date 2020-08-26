package vaultClient

import vault "github.com/hashicorp/vault/api"

type SecretMetadata struct {
	CreatedTime  string `json:"created_time"`
	DeletionTime string `json:"deletion_time"`
	Version      int
	Destroyed    bool
}

func DefaultConfig() *vault.Config {
	return vault.DefaultConfig()
}

type Client struct {
	client *vault.Client
	Auth   *auth
	KV2    *kv2
}

func NewClient(config *vault.Config) (*Client, error) {
	client, err := vault.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		Auth:   NewAuth(client),
		KV2:    &kv2{client: client}}, nil
}
