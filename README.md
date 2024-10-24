# yakvs
Yet another key value store

Omegapoint Innovation Week 2024 Oct

## How to run
Create a .env file
* HTTPS_PORT=8443
* GRPC_PORT=7443

## Tls

* Generated self signed cert for 7 days

`
openssl req -x509 -newkey rsa:4096 -keyout privatekey.pem -out cert.pem -days 7 -nodes -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
`

## Tools
To run some load tests in /test
* https://github.com/rakyll/hey
* https://github.com/bojand/ghz
