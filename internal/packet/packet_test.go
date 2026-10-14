package packet

import (
	"strings"
	"testing"
)

func TestReplySwapsAddressesAndKeepsHWAddr(t *testing.T) {
	req := Packet{Src: "203.0.113.1:49152", Dst: "203.0.113.80:80", HWAddr: "02:42:c0:a8:01:0a", Body: "hello"}
	got := req.Reply("Hello!")
	want := Packet{Src: "203.0.113.80:80", Dst: "203.0.113.1:49152", HWAddr: "02:42:c0:a8:01:0a", Body: "Hello!"}
	if got != want {
		t.Fatalf("Reply() = %+v, want %+v", got, want)
	}
}

func TestHostAndPort(t *testing.T) {
	tests := []struct {
		addr, host, port string
	}{
		{"192.168.1.10:49152", "192.168.1.10", "49152"},
		{"203.0.113.1", "203.0.113.1", ""},
	}
	for _, tt := range tests {
		if got := Host(tt.addr); got != tt.host {
			t.Errorf("Host(%q) = %q, want %q", tt.addr, got, tt.host)
		}
		if got := Port(tt.addr); got != tt.port {
			t.Errorf("Port(%q) = %q, want %q", tt.addr, got, tt.port)
		}
	}
}

func TestDecodeRejectsMissingAddresses(t *testing.T) {
	_, err := Decode(strings.NewReader(`{"body":"hello"}`))
	if err == nil {
		t.Fatal("Decode() accepted a packet without src and dst")
	}
}
