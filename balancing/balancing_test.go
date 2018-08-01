package balancing

import (
	"reflect"
	"testing"

	iphash "github.com/hlts2/ip-hash"
	leastconnections "github.com/hlts2/least-connections"
	roundrobin "github.com/hlts2/round-robin"
)

var testData = []string{
	"server-1",
	"server-2",
	"server-3",
}

func TestGetLeastConnections(t *testing.T) {
	rc, err := leastconnections.New(testData)
	if err != nil {
		t.Errorf("leastconnections.New is error: %v", err)
	}

	b := &Balancing{
		algorithm: rc,
	}

	got := b.GetLeastConnections()

	if !reflect.DeepEqual(rc, got) {
		t.Errorf("GetLeastConnections is wrong. expected: %v, got: %v", rc, got)
	}
}

func TestGetRoundRobin(t *testing.T) {
	rr, err := roundrobin.New(testData)
	if err != nil {
		t.Errorf("leastconnections.New is error: %v", err)
	}

	b := &Balancing{
		algorithm: rr,
	}

	got := b.GetRoundRobin()

	if !reflect.DeepEqual(rr, got) {
		t.Errorf("GetLeastConnections is wrong. expected: %v, got: %v", rr, got)
	}
}

func TestGetIPHash(t *testing.T) {
	ih, err := iphash.New(testData)
	if err != nil {
		t.Errorf("leastconnections.New is error: %v", err)
	}

	b := &Balancing{
		algorithm: ih,
	}

	got := b.GetIPHash()

	if !reflect.DeepEqual(ih, got) {
		t.Errorf("GetLeastConnections is wrong. expected: %v, got: %v", ih, got)
	}
}
