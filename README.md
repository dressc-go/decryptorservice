# decryptorservice

Simple helper daemon to decrypt strings via Web API.

## Usage

The started binary reads a configuration file with its settings.

The location of the file can be given in the environment variable ConfigFile or it will try ./config.yml and /run/secrets/config.yml

config.yml:
```
IpAddress: 127.0.1.15          # IP Adress of the server
Port: 8167                     # Port of the server
DecPrivateKeyFile: privkey.pem # Decryption key
TLSCertFile: pubcert.pem       # To activate TLS, CertFile and ...
TLSKeyFile: privkey.pem        # key must be given
```

The service is queried via JSON in a POST request:
```
{
  "c": "base64OeapSha1",
  "ctxt": "ZU1QVE...aIce6Ryvwkwey8Fkw="
}
```
With "c" selecting the cipher (only base64OeapSha1 supported for now) and "ctxt" holding the ciphertext.

The server answers with the cleartext in JSON