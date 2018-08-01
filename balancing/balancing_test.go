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

func TestNew(t *testing.T) {
	lc, err := leastconnections.New(testData)
	if err != nil {
		t.Errorf("leastconnections.New is error: %v", err)
	}

	blc := New(lc)
	if blc == nil {
		t.Errorf("New(leastconnections) is nil")
	}

	rr, err := roundrobin.New(testData)
	if err != nil {
		t.Errorf("roundrobin.New is error: %v", err)
	}

	brr := New(rr)
	if brr == nil {
		t.Errorf("New(roundrobin) is nil")
	}

	ih, err := iphash.New(testData)
	if err != nil {
		t.Errorf("iphash.New is error: %v", err)
	}

	bih := New(ih)
	if bih == nil {
		t.Error("New(ip-hash) is nil")
	}

	bnone := New(nil)
	if bnone != nil {
		t.Errorf("New(nil) is wrong. expected: %v, got: %v", nil, bnone)
	}
}

func TestGetLeastConnections(t *testing.T) {
	lc, err := leastconnections.New(testData)
	if err != nil {
		t.Errorf("leastconnections.New is error: %v", err)
	}

	b := &Balancing{
		algorithm: lc,
	}

	got := b.GetLeastConnections()

	if !reflect.DeepEqual(lc, got) {
		t.Errorf("GetLeastConnections is wrong. expected: %v, got: %v", lc, got)
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
