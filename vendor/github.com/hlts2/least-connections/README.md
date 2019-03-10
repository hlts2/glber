# least-connections
least-connections is least-connections balancing algorithm written in golang

## Requrement

Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/least-connections
```

## Example

### Basic Example
```go
lc, err := New([]*url.URL{
    {Host: "192.168.33.10"},
    {Host: "192.168.33.11"},
    {Host: "192.168.33.12"},
})

src1, done1 := lc.Next() // {Host: "192.168.33.10"}

src2, done2 := lc.Next() // {Host: "192.168.33.11"}

done1() // Reduce connection of src1

src3, done3 := lc.Next() // {Host: "192.168.33.10"}

```


## Author
[hlts2](https://github.com/hlts2)

## LICENSE
least-connections released under MIT license, refer [LICENSE](https://github.com/hlts2/least-connections/blob/master/LICENSE) file.
