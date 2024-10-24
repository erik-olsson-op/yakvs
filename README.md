# yakvs

Yet another key value store

Omegapoint Innovation Week 2024 Oct

## How to run

Create a .env file

* HTTPS_PORT=8443
* GRPCS_PORT=7443
* SERVER_CERT
* SERVER_KEY

## TLS

* Generated self-signed cert for 7 days for localhost

`
openssl req -x509 -newkey rsa:4096 -keyout x509/localhost-key.pem -out x509/localhost-cert.pem -days 7 -nodes -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
`

## Tools

To run some load tests in /test

* https://github.com/rakyll/hey
* https://github.com/bojand/ghz
