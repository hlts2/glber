package slb

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var testYaml = `
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
    port: 3333
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
		ServerConfig: ServerConfig{
			Host: "0.0.0.0",
			Port: "80",
		},
		Balancing: "round-robin",
		TLSConfig: TLSConfig{
			Enabled: true,
			CertKey: "./cert.key",
			KeyKey:  "./key.key",
		},
		BackendServerConfigs: ServerConfigs{
			{
				Scheme: "http",
				Host:   "192.168.33.11",
				Port:   "1111",
			},
			{
				Scheme: "http",
				Host:   "192.168.33.11",
				Port:   "2222",
			},
			{
				Scheme: "http",
				Host:   "192.168.33.11",
				Port:   "3333",
			},
		},
	}

	if !reflect.DeepEqual(want, c) {
		t.Errorf("LoadConfig is wrong. want: %v, got: %v", want, c)
	}
}

func TestDuplicateExists(t *testing.T) {
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
		if want, got := test.want, duplicateExists(test.input); want != got {
			t.Errorf("tests[%d] - duplicateExists is wrong. want: %v, got: %v", i, want, got)
		}
	}
}

func createFile(filename string, data []byte) {
	ioutil.WriteFile(filename, data, os.ModePerm)
}

func deleteFile(filename string) {
	os.Remove(filename)
}
