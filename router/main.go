// Command router is the masquerading NAT router, the heart of the
// simulation.
//
// Outbound, it files the packet's original source under the sender's
// hardware address, rewrites the source to the router's own public (WAN)
// address, and forwards the packet to the server. Inbound, it looks the
// hardware address up again, restores the client's real address as the
// destination, and hands the reply back to the client.
//
// This is what `iptables -t nat -A POSTROUTING -j MASQUERADE` asks the
// Linux kernel to do, minus the kernel.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fzn0x/nat-simulation/internal/packet"
)

func main() {
	listen := flag.String("listen", "127.0.0.1:8000", "real address to listen on (the LAN side)")
	wan := flag.String("wan", "203.0.113.1", "simulated public (WAN) IP of this router")
	server := flag.String("server", "http://127.0.0.1:8080", "where the server really listens")
	flag.Parse()

	log.SetFlags(log.Ltime)
	log.SetPrefix("[router] ")

	r := &router{wan: *wan, serverURL: *server, table: NewTable()}
	mux := http.NewServeMux()
	mux.Handle("/", r)
	mux.HandleFunc("/table", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		r.table.WriteTo(w)
	})

	log.Printf("WAN %s listening on %s, forwarding to %s", *wan, *listen, *server)
	log.Printf("translation table: curl http://%s/table", *listen)
	log.Fatal(http.ListenAndServe(*listen, mux))
}

type router struct {
	wan       string // public IP that every outgoing packet is rewritten to
	serverURL string
	table     *Table
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "send a packet with POST", http.StatusMethodNotAllowed)
		return
	}
	p, err := packet.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if p.HWAddr == "" {
		http.Error(w, "packet has no hardware address, so the router can't file it", http.StatusBadRequest)
		return
	}

	// Outbound: remember who sent it, then put the router's address on it.
	log.Printf("OUT  %s", p)
	inside := p.Src
	rt.table.Record(p.HWAddr, inside)
	p.Src = fmt.Sprintf("%s:%s", rt.wan, packet.Port(inside))
	log.Printf("     recorded %s -> %s", p.HWAddr, inside)
	log.Printf("     rewrote src %s -> %s", inside, p.Src)

	reply, status, err := packet.Send(rt.serverURL, p)
	if err != nil {
		log.Printf("     server unreachable: %v", err)
		http.Error(w, "router could not reach the server: "+err.Error(), http.StatusBadGateway)
		return
	}

	// Inbound: the reply is addressed to the router. Find the real client.
	log.Printf("IN   %s", reply)
	client, ok := rt.table.Lookup(reply.HWAddr)
	if !ok {
		log.Printf("     no entry for %s, dropping reply", reply.HWAddr)
		http.Error(w, "router has no translation for "+reply.HWAddr, http.StatusBadGateway)
		return
	}
	log.Printf("     restored dst %s -> %s", reply.Dst, client)
	reply.Dst = client
	packet.Write(w, status, reply)
}
