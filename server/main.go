// Command server is the far end of the simulation: a host on the internet
// that only answers sources on its allow-list.
//
// It reads each packet's body, checks the source address, and replies
// "Hello!" to allowed sources and "Not Permitted" to everyone else. By
// default only the router's public address is allowed, so a client can
// only get through by hiding behind the router.
package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/fzn0x/nat-simulation/internal/packet"
)

func main() {
	listen := flag.String("listen", "127.0.0.1:8080", "real address to listen on")
	addr := flag.String("addr", "203.0.113.80:80", "simulated public address of this server")
	allow := flag.String("allow", "203.0.113.1", "comma-separated source IPs allowed to talk to this server")
	flag.Parse()

	log.SetFlags(log.Ltime)
	log.SetPrefix("[server] ")

	s := newServer(*addr, strings.Split(*allow, ","))
	log.Printf("%s listening on %s, allow-list: %s", *addr, *listen, *allow)
	log.Fatal(http.ListenAndServe(*listen, s))
}

type server struct {
	addr    string
	allowed map[string]bool
}

func newServer(addr string, allow []string) *server {
	s := &server{addr: addr, allowed: make(map[string]bool)}
	for _, ip := range allow {
		if ip = strings.TrimSpace(ip); ip != "" {
			s.allowed[ip] = true
		}
	}
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "send a packet with POST", http.StatusMethodNotAllowed)
		return
	}
	p, err := packet.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// The server only ever looks at the source IP. It has no idea whether
	// a router rewrote it on the way.
	src := packet.Host(p.Src)
	if !s.allowed[src] {
		log.Printf("DENY  %-20s not on the allow-list, replying Not Permitted", p.Src)
		packet.Write(w, http.StatusForbidden, p.Reply("Not Permitted"))
		return
	}
	log.Printf("ALLOW %-20s on the allow-list, replying Hello!", p.Src)
	packet.Write(w, http.StatusOK, p.Reply("Hello!"))
}
