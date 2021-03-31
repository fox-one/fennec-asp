# fennec-asp
Fennec Account Service Provider

## Run

```

// default config file: ./config/config.yaml
// default server port: 80
$fennec api --config xxxx --port 9000
```

## Deployment

* Enviroment
    1. local
    2. test
    3. prod
   
* Config file

Put the configuration file of the corresponding environment in the deploy directory. e.g. `config.prod.yaml`. See [config file template](deploy/config.yaml.tpl)

```
dapp:
  client_id: 
  session_id: 
  pin_token: 
  pin: ''
  private_key: ~
```

* Build

See [Makefile](Makefile)

```
// build prod executable file
$make build-prod

// build docker image
$make docker-build-prod

// push to the docker repository
$make deploy-prod
```

