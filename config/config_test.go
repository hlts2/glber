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

func TestLoadConfig(t *testing.T) {
	const filename string = "test.yaml"
	defer deleteFile(filename)
	createFile(filename, []byte(testYaml))

	var c Config
	err := LoadConfig(filename, &c)
	if err != nil {
		t.Errorf("LoadConfig is faild. error: %v, c: %v", err, c)
	}

	expected := Config{
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

	if !reflect.DeepEqual(expected, c) {
		t.Errorf("LoadConfig is wrong. expected: %v, got: %v", expected, c)
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		servers  Servers
		expected []string
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
			expected: []string{
				"http://192.168.33.10:1111",
				"http://192.168.33.10:2222",
				"https://192.168.33.10:3333",
			},
		},
	}

	for i, test := range tests {
		got := test.servers.GetAddresses()

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - ToStringSlice is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestExistsDuplicateHost(t *testing.T) {
	tests := []struct {
		hosts    []string
		expected bool
	}{
		{
			hosts: []string{
				"192.168.33.10:1111",
				"192.168.33.10:2222",
			},
			expected: false,
		},
		{
			hosts: []string{
				"192.168.33.10:2222",
				"192.168.33.10:2222",
			},
			expected: true,
		},
		{
			hosts:    []string{},
			expected: false,
		},
	}

	for i, test := range tests {
		got := existsDuplicateHost(test.hosts)

		if test.expected != got {
			t.Errorf("tests[%d] - existsDuplicateHost is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func createFile(filename string, data []byte) {
	ioutil.WriteFile(filename, data, os.ModePerm)
}

func deleteFile(filename string) {
	os.Remove(filename)
}
