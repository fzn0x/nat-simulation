package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fzn0x/nat-simulation/internal/packet"
)

func init() { log.SetOutput(io.Discard) }

// fakeServer answers like the real server and remembers the last packet it saw.
func fakeServer(t *testing.T, seen *packet.Packet) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, err := packet.Decode(r.Body)
		if err != nil {
			t.Errorf("server got an unreadable packet: %v", err)
			return
		}
		*seen = p
		packet.Write(w, http.StatusOK, p.Reply("Hello!"))
	}))
	t.Cleanup(srv.Close)
	return srv
}

func route(t *testing.T, rt *router, p packet.Packet) (packet.Packet, int) {
	t.Helper()
	body, _ := json.Marshal(p)
	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body)))
	if rec.Code != http.StatusOK {
		return packet.Packet{}, rec.Code
	}
	reply, err := packet.Decode(rec.Body)
	if err != nil {
		t.Fatalf("router sent an unreadable reply: %v", err)
	}
	return reply, rec.Code
}

func TestRouterMasqueradesAndRestores(t *testing.T) {
	var seen packet.Packet
	srv := fakeServer(t, &seen)
	rt := &router{wan: "203.0.113.1", serverURL: srv.URL, table: NewTable()}

	reply, code := route(t, rt, packet.Packet{
		Src: "192.168.1.10:49152", Dst: "203.0.113.80:80", HWAddr: "02:42:c0:a8:01:0a", Body: "hello",
	})
	if code != http.StatusOK {
		t.Fatalf("router returned %d", code)
	}
	if seen.Src != "203.0.113.1:49152" {
		t.Errorf("server saw src %q, want the router's address 203.0.113.1:49152", seen.Src)
	}
	if reply.Dst != "192.168.1.10:49152" {
		t.Errorf("reply dst = %q, want the client's real address restored", reply.Dst)
	}
	if inside, _ := rt.table.Lookup("02:42:c0:a8:01:0a"); inside != "192.168.1.10:49152" {
		t.Errorf("table entry = %q, want 192.168.1.10:49152", inside)
	}
}

func TestRouterKeepsOneEntryPerHardwareAddress(t *testing.T) {
	var seen packet.Packet
	srv := fakeServer(t, &seen)
	rt := &router{wan: "203.0.113.1", serverURL: srv.URL, table: NewTable()}

	clients := []packet.Packet{
		{Src: "192.168.1.10:49152", Dst: "203.0.113.80:80", HWAddr: "02:42:c0:a8:01:0a"},
		{Src: "192.168.1.11:49152", Dst: "203.0.113.80:80", HWAddr: "02:42:c0:a8:01:0b"},
	}
	for _, c := range clients {
		reply, _ := route(t, rt, c)
		if reply.Dst != c.Src {
			t.Errorf("reply for %s went to %s", c.HWAddr, reply.Dst)
		}
	}

	var out strings.Builder
	rt.table.WriteTo(&out)
	for _, c := range clients {
		if !strings.Contains(out.String(), c.Src) {
			t.Errorf("table is missing %s:\n%s", c.Src, out.String())
		}
	}
}

func TestRouterRejectsPacketWithoutHardwareAddress(t *testing.T) {
	rt := &router{wan: "203.0.113.1", serverURL: "http://127.0.0.1:0", table: NewTable()}
	_, code := route(t, rt, packet.Packet{Src: "192.168.1.10:49152", Dst: "203.0.113.80:80"})
	if code != http.StatusBadRequest {
		t.Fatalf("router returned %d, want 400", code)
	}
}

func TestRouterDropsReplyWithUnknownHardwareAddress(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := packet.Decode(r.Body)
		reply := p.Reply("Hello!")
		reply.HWAddr = "02:00:00:00:00:99" // a reply the router never asked for
		packet.Write(w, http.StatusOK, reply)
	}))
	defer srv.Close()
	rt := &router{wan: "203.0.113.1", serverURL: srv.URL, table: NewTable()}

	_, code := route(t, rt, packet.Packet{Src: "192.168.1.10:49152", Dst: "203.0.113.80:80", HWAddr: "02:42:c0:a8:01:0a"})
	if code != http.StatusBadGateway {
		t.Fatalf("router returned %d, want 502", code)
	}
}
