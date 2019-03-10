# ip-hash
ip-hash is balancing algorithm, based on [round-robin](https://github.com/hlts2/round-robin).

## Requrement

Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/ip-hash
```

## Example
```go
ip, _ := iphash.New([]*url.URL{
    {Host: "192.168.33.10"},
    {Host: "192.168.33.11"},
    {Host: "192.168.33.12"},
 })

ip.Next(&url.URL{Host: "192.168.33.10"})  // {Host: "192.168.33.10"}
ip.Next(&url.URL{Host: "192.168.33.10"})  // {Host: "192.168.33.10"}
ip.Next(&url.URL{Host: "192.168.33.44"})  // {Host: "192.168.33.11"}
ip.Next(&url.URL{Host: "192.168.33.44"})  // {Host: "192.168.33.11"}
```

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
ip-hash released under MIT license, refer [LICENSE](https://github.com/hlts2/ip-hash/blob/master/LICENSE) file.
