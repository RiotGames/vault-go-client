package vault

import hashivault "github.com/hashicorp/vault/api"

type SecretMetadata struct {
	CreatedTime  string `json:"created_time"`
	DeletionTime string `json:"deletion_time"`
	Version      int
	Destroyed    bool
}

func DefaultConfig() *hashivault.Config {
	return hashivault.DefaultConfig()
}

type Client struct {
	client *hashivault.Client
	Auth   *auth
	KV2    *kv2
}

func NewClient(config *hashivault.Config) (*Client, error) {
	client, err := hashivault.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		Auth:   NewAuth(client),
		KV2:    &kv2{client: client}}, nil
}
