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

```yaml
host: 0.0.0.0
port: 80
balancing: round-robin
tls:
  enabled: true
  cert_key: ./cert.key
  key_key: ./key.key
servers:
  - scheme: http
    host: 192.168.33.11
    port: 1111
  - scheme: http
    host: 192.168.33.11
    port: 2222
  - scheme: http
    host: 192.168.33.11

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
$ go-LB serve -s config.yml
```

## CLI Usage

```
$ go-LB --help
NAME:
   go-LB - Load Balancer

USAGE:
   go-LB [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

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
   --set value, -s value  set the configuration file (default: "config.yml")
```

## TODO

- [ ] Helth check of service

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
go-LB released under MIT license, refer [LICENSE](https://github.com/hlts2/go-LB/blob/master/LICENSE) file.
