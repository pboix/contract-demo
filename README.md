# Contract Testing Demo 🤝
In this demo we will showcase two different contract testing approaches, as showed in the Bugbusters contrac testing meetup.

## Prerequisites
* Node
* Yarn / NPM
* Go
* Your folder is inside $GOPATH

## Pact

```bash
cd pact
```

### Run consumer tests

```bash
cd consumer
yarn
yarn jest
```

This will generate the corresponding pact under the `pacts` folder, which will then be used to validate our provider.

### Run provider tests

Ensure you've installed pact's [CLI Tools](https://github.com/pact-foundation/pact-ruby-standalone/releases).

```bash
cd provider
go test . -v
```

## Apidocs (Swagger)

```bash
cd swagger
```

The API definition is under apidocs/specs/openapi.yaml

### Run consumer tests

To run our consumer, we will be using [Prism](https://github.com/stoplightio/prism).

Be sure to have Prism installed: `curl https://raw.githubusercontent.com/stoplightio/prism/2.x/install.sh | sh`

```bash
cd consumer
prism mock -s ../apidocs/specs/openapi.yaml
```

This will create a Prism mock server based on our specifications. Prism will remain running in the foreground. So, in another terminal...

```bash
cd consumer
yarn
yarn jest
```
This will validate our consumer against our Prism mock server, making sure that what the consumer expects is the same as the specifications in our swagger definition.

## Run provider tests

First, let's make sure we have [Dredd](https://github.com/apiaryio/dredd/) installed:
`yarn global add dredd`

In order to validate our provider, we will need to run the service first, so it can respond in its endpoints.

```bash
cd provider
go run provider.go
```
This will leave our provider running in the foreground, so in another terminal we can run dredd against the service:

```bash
cd provider
dredd ~/dev/go/src/github.com/pboix/contract-demo/swagger/apidocs/specs/openapi.yaml  http://localhost:8000
```

## Contributors
 [Toni Feliu](https://github.com/tonimellodic)

 [Daniel Giralt](https://github.com/daniel-giralt-len)