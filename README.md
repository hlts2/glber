# go-LB [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![GoDoc](http://godoc.org/github.com/hlts2/go-LB?status.svg)](http://godoc.org/github.com/hlts2/go-LB) [![Go Report Card](https://goreportcard.com/badge/github.com/hlts2/go-LB)](https://goreportcard.com/report/github.com/hlts2/go-LB) [![Join the chat at https://gitter.im/hlts2/go-LB](https://badges.gitter.im/hlts2/go-LB.svg)](https://gitter.im/hlts2/go-LB?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

go-LB is a simple lightweight load balancer written in golang.

## Requirement
Go (>= 1.9)

## Installation

```shell
go get github.com/hlts2/go-LB
```

## Example

### Config file

Config file `config.yml` describes configuration of load balancer.
The following is a setting example.

```
servers:
  -
    scheme: http
    host: 192.168.33.10
    port: 1111

  -
    scheme: http
    host: 192.168.33.10
    port: 2222

  -
    scheme: http
    host: 192.168.33.10
    port: 3333

balancing: round-robin
```

### Balancing Algorithm

There are three possible algorithms for balancing

- [round-robin](https://github.com/hlts2/round-robin)
- [ip-hash](https://github.com/hlts2/ip-hash)
- [least-connections](https://github.com/hlts2/least-connections)

Please write algorithm name in `balancing` field of `config.yml` file

```
balancing: round-robin # or ip-hash or least-connections
```

### Basic Example

```
# Not TLS Mode
$ go-LB serve -s config.yml -H x.x.x.x -p 8080

# TLS Mode
$ go-LB serve -s config.yml -H x.x.x.x -p 8080 -tlspath ./tlsdirectory
```

## CLI Usage

```
$ go-LB --help
NAME:
  go-LB - Load Balancer

USAGE:
  go-LB [global options] command [command options] [arguments...]

VERSION:
  v0.0.1

COMMANDS:
  serve    serve load balancer
  help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
  --help, -h     show help
  --version, -v  print the version
```

### Serve command

```
$ go-LB serve --help
NAME:
  go-LB serve - serve load balancer

USAGE:
  go-LB serve [command options] [arguments...]

OPTIONS:
  --set value, -s value   set a config file of load balancer (default: "config.yml")
  --host value, -H value  set a host name or IP of load balancer (default: "127.0.0.1")
  --port value, -p value  set a port number of load balancer (default: "8080")
  --tlspath value         set a TLS directory of load balancer
```

## TODO

- [ ] Helth check of service

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
go-LB released under MIT license, refer [LICENSE](https://github.com/hlts2/go-LB/blob/master/LICENSE) file.
