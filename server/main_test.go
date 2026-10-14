package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fzn0x/nat-simulation/internal/packet"
)

func send(t *testing.T, h http.Handler, p packet.Packet) (packet.Packet, int) {
	t.Helper()
	body, _ := json.Marshal(p)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body)))
	reply, err := packet.Decode(rec.Body)
	if err != nil {
		t.Fatalf("server sent an unreadable reply: %v", err)
	}
	return reply, rec.Code
}

func TestServerAllowsListedSource(t *testing.T) {
	s := newServer("203.0.113.80:80", []string{"203.0.113.1"})
	reply, code := send(t, s, packet.Packet{Src: "203.0.113.1:49152", Dst: "203.0.113.80:80", HWAddr: "02:42:c0:a8:01:0a", Body: "hello"})

	if code != http.StatusOK || reply.Body != "Hello!" {
		t.Fatalf("got %d %q, want 200 \"Hello!\"", code, reply.Body)
	}
	if reply.Dst != "203.0.113.1:49152" {
		t.Errorf("reply dst = %q, want it addressed back to the sender", reply.Dst)
	}
	if reply.HWAddr != "02:42:c0:a8:01:0a" {
		t.Errorf("reply hwaddr = %q, want the request's hardware address echoed", reply.HWAddr)
	}
}

func TestServerRejectsUnlistedSource(t *testing.T) {
	s := newServer("203.0.113.80:80", []string{"203.0.113.1"})
	reply, code := send(t, s, packet.Packet{Src: "192.168.1.10:49152", Dst: "203.0.113.80:80", Body: "hello"})

	if code != http.StatusForbidden || reply.Body != "Not Permitted" {
		t.Fatalf("got %d %q, want 403 \"Not Permitted\"", code, reply.Body)
	}
}

func TestServerTrimsAllowList(t *testing.T) {
	s := newServer("203.0.113.80:80", []string{" 203.0.113.1 ", "", "192.168.1.10"})
	if !s.allowed["203.0.113.1"] || !s.allowed["192.168.1.10"] || len(s.allowed) != 2 {
		t.Fatalf("allow-list = %v", s.allowed)
	}
}
