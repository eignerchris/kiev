# Kiev

An in memory key/value document store built to learn go.

Uses a space delimited TCP protocol "CMD KEY [VALUE]"


## Usage

1. `go run kiev.go`

2. In another terminal,  set some data `echo -n "SET posts:1 {\"id\" : 1, \"title\": 'Catcher in the Rye'}" | nc localhost 8745`

3. Query some data, `echo -n "GET users:1" | nc localhost 8745`

## TODO

  * build small client
  * break kiev.go into logical abstrations: server.go, validation.go, config.go, etc.
  * add support for wildcard key fetching: `GET *` or `GET users:*`
  * add ability to configure port, max memory, etc.
  * check size of map before SET'ing to make sure not exceeding max memory
  * create binary that access cli options