## Prerequisites
* Node
* Yarn / NPM
* Go
* Your folder is inside GOPATH

## Pact

```bash
cd pact
```

### Run consumer

```bash
cd consumer
yarn
yarn jest
```

### Run provider

Ensure you've installed pact's [CLI Tools](https://github.com/pact-foundation/pact-ruby-standalone/releases).

```bash
cd provider
go test . -v
```

## Apidocs

```bash
cd swagger
```

The API definition is under apidocs/specs/openapi.yaml

### Run consumer

Be sure to have Prism installed: `curl https://raw.githubusercontent.com/stoplightio/prism/2.x/install.sh | sh`

```bash
cd consumer
prism mock -s ../apidocs/specs/openapi.yaml
```

Prism will remain hung up. So, in another terminal...

```bash
cd consumer
yarn
yarn jest
```

## Run provider

Be sure to have dredd installed: `yarn install -g dredd`

```bash
cd provider
go run provider.go
```

The provider will remain hung up. So, in another terminal...

```bash
cd provider
dredd ~/dev/go/src/github.com/pboix/contract-demo/swagger/apidocs/specs/openapi.yaml  http://localhost:8000
```
