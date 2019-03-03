package config

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var testYaml = `
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

balancing: ip-hash
`

func TestLoad(t *testing.T) {
	const filename string = "test.yaml"
	defer deleteFile(filename)
	createFile(filename, []byte(testYaml))

	var c Config
	err := Load(filename, &c)
	if err != nil {
		t.Errorf("LoadConfig is faild. error: %v, c: %v", err, c)
	}

	want := Config{
		Servers: Servers{
			{
				Scheme: "http",
				Host:   "192.168.33.10",
				Port:   "1111",
			},
			{
				Scheme: "http",
				Host:   "192.168.33.10",
				Port:   "2222",
			},
			{
				Scheme: "http",
				Host:   "192.168.33.10",
				Port:   "3333",
			},
		},
		Balancing: "ip-hash",
	}

	if !reflect.DeepEqual(want, c) {
		t.Errorf("LoadConfig is wrong. want: %v, got: %v", want, c)
	}
}

func TestGetAddresses(t *testing.T) {
	tests := []struct {
		servers Servers
		want    []string
	}{
		{
			servers: Servers{
				{
					Scheme: "http",
					Host:   "192.168.33.10",
					Port:   "1111",
				},
				{
					Scheme: "http",
					Host:   "192.168.33.10",
					Port:   "2222",
				},
				{
					Scheme: "https",
					Host:   "192.168.33.10",
					Port:   "3333",
				},
			},
			want: []string{
				"http://192.168.33.10:1111",
				"http://192.168.33.10:2222",
				"https://192.168.33.10:3333",
			},
		},
	}

	for i, test := range tests {
		got := test.servers.GetAddresses()

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("tests[%d] - ToStringSlice is wrong. want: %v, got: %v", i, test.want, got)
		}
	}
}

func TestExistsDuplicateHost(t *testing.T) {
	tests := []struct {
		input []string
		want  bool
	}{
		{
			input: []string{
				"192.168.33.10:1111",
				"192.168.33.10:2222",
			},
			want: false,
		},
		{
			input: []string{
				"92.168.33.10:2222",
				"92.168.33.10:2222",
			},
			want: true,
		},
		{
			input: []string{},
			want:  false,
		},
	}

	for i, test := range tests {
		got := duplicateExists(test.input)

		if test.want != got {
			t.Errorf("tests[%d] - existsDuplicateHost is wrong. want: %v, got: %v", i, test.want, got)
		}
	}
}

func createFile(filename string, data []byte) {
	ioutil.WriteFile(filename, data, os.ModePerm)
}

func deleteFile(filename string) {
	os.Remove(filename)
}
