package vaultClient

import (
	"encoding/json"
	"fmt"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

type kv2 struct {
	client *vault.Client
}

type KV2GetOptions struct {
	MountPath     string
	SecretPath    string
	UnmarshalInto interface{}
}

type KV2PutOptions struct {
	MountPath  string
	SecretPath string
	Secrets    map[string]interface{}
}

func (k *kv2) Put(options KV2PutOptions) (*vault.Secret, error) {
	mountPath := "secret"
	if options.MountPath != "" {
		mountPath = strings.Trim(options.MountPath, "/")
	}

	putPath := mountPath + "/data/" + strings.Trim(options.SecretPath, "/")
	secret, err := k.write(putPath, options.Secrets)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (k *kv2) Get(options KV2GetOptions) (*vault.Secret, error) {
	mountPath := "secret"
	if options.MountPath != "" {
		mountPath = strings.Trim(options.MountPath, "/")
	}

	readPath := mountPath + "/data/" + strings.Trim(options.SecretPath, "/")
	secret, err := k.read(readPath, map[string][]string{})
	if err != nil {
		return nil, err
	}

	if secret == nil {
		return nil, fmt.Errorf("No secret found at path: %s/%s", mountPath, options.SecretPath)
	}

	if options.UnmarshalInto != nil {
		dataBytes, err := json.Marshal(secret.Data["data"])
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(dataBytes, options.UnmarshalInto); err != nil {
			return nil, err
		}
		metadataBytes, err := json.Marshal(secret.Data["metadata"])
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(metadataBytes, options.UnmarshalInto); err != nil {
			return nil, err
		}

	}

	return secret, nil
}

func (k *kv2) write(path string, data map[string]interface{}) (*vault.Secret, error) {
	normalizedData := map[string]interface{}{
		"data":    data,
		"options": map[string]interface{}{},
	}
	return k.client.Logical().Write(path, normalizedData)
}

func (k *kv2) read(path string, data map[string][]string) (*vault.Secret, error) {
	if len(data) == 0 {
		return k.client.Logical().Read(path)
	}

	return k.client.Logical().ReadWithData(path, data)
}
