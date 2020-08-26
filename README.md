vault-go-client
=====
This is a Golang client for Vault. 

# Usage
## Creating a Client
The following will create a client with default configuration:
```
// Uses VAULT_ADDR env var to set the clients URL
client, err := vaultClient.NewClient(vaultClient.DefaultConfig())
if err != nil {
    log.Fatal(err.Error())
}
```

## Putting a Secret into Vault
The following will put a secret into Vault:
```
secretMap := map[string]interface{}{
    "hello": "world",
}

if _, err = client.KV2.Put(vaultClient.KV2PutOptions{
	MountPath:  secretMountPath,
	SecretPath: secretPath,
	Secrets:    secretMap,
}); err != nil {
	log.Fatal(err.Error())
}
```

## Retrieving a Secret from Vault
### Unmarshaling Approach
This approach unmarshals the secret from Vault into the provided struct. 
The embedded struct `vaultClient.SecretMetadata` is optional.
```
// Injecting vaultClient.SecretMetadata is optional
type Secret struct {
	Hello string `json:"hello"`
	vaultClient.SecretMetadata
}
...
secret := &Secret{}

if _, err = client.KV2.Get(vaultClient.KV2GetOptions{
	MountPath:     secretMountPath,
	SecretPath:    secretPath,
	UnmarshalInto: secret,
}); err != nil {
	log.Fatal(err.Error())
}
fmt.Printf("%v\n", secret)
```
### Raw Secret Approach
This approach returns a `Secret` defined in `github.com/hashicorp/vault/api`.
```
// This returns a 
secret, err := client.KV2.Get(vaultClient.KV2GetOptions{
	MountPath:  secretMountPath,
	SecretPath: secretPath,
})

if err != nil {
	log.Fatal(err.Error())
}
```

