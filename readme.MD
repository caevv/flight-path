# Calculate the Flight Path of a Person

## Server

The service is running on port `8080` or can be changed with the env variable `APP_SERVER_PORT`

## Endpoints

`/calculate`

expected body is an array of a source and destination flight, example:

```json
[
  [
    "SFO",
    "EWR"
  ]
]
```

Response:

```json
[
  "SFO",
  "EWR"
]
```

## Other commands

`make test` will run all the tests

`make gorun` will use go to run the service

`make build` will build the service artifact
