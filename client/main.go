// Command client is a machine on the private network behind the router.
//
// Every few seconds it sends the same request twice: once through the
// router and once straight to the server. The server only trusts the
// router's address, so the routed request is answered with "Hello!" and
// the direct one with "Not Permitted". Same client, same request, different
// path: that difference is the masquerade.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fzn0x/nat-simulation/internal/packet"
)

func main() {
	addr := flag.String("addr", "192.168.1.10", "simulated private IP of this client")
	hw := flag.String("hw", "02:42:c0:a8:01:0a", "simulated hardware (MAC) address of this client")
	dst := flag.String("dst", "203.0.113.80:80", "simulated address of the server")
	routerURL := flag.String("router", "http://127.0.0.1:8000", "where the router really listens")
	serverURL := flag.String("server", "http://127.0.0.1:8080", "where the server really listens")
	body := flag.String("body", "hello", "request body")
	every := flag.Duration("every", 4*time.Second, "time between rounds")
	rounds := flag.Int("rounds", 0, "stop after this many rounds (0 runs forever)")
	flag.Parse()

	log.SetFlags(log.Ltime)
	log.SetPrefix("[client] ")
	log.Printf("%s (hw %s) talking to %s every %s", *addr, *hw, *dst, *every)

	paths := []struct {
		name, peer, url string
	}{
		{"via router", "router", *routerURL},
		{"direct    ", "server", *serverURL},
	}

	// Each round is a new conversation, so it gets a new source port, the
	// way an operating system hands out ephemeral ports.
	port := 49152
	for round := 1; *rounds == 0 || round <= *rounds; round++ {
		req := packet.Packet{
			Src:    fmt.Sprintf("%s:%d", *addr, port),
			Dst:    *dst,
			HWAddr: *hw,
			Body:   *body,
		}
		log.Printf("round %d: %s -> %s", round, req.Src, req.Dst)
		for _, path := range paths {
			reply, status, err := packet.Send(path.url, req)
			if err != nil {
				log.Printf("  %s  error: %s", path.name, explain(err, path.peer, path.url))
				continue
			}
			log.Printf("  %s  %d %-14s reply to %s", path.name, status, reply.Body, reply.Dst)
		}

		port++
		if port > 65535 {
			port = 49152
		}
		if *rounds == 0 || round < *rounds {
			time.Sleep(*every)
		}
	}
}

// explain turns a connection failure into a hint about which program to start.
func explain(err error, peer, url string) string {
	var op *net.OpError
	if errors.As(err, &op) && op.Op == "dial" {
		return fmt.Sprintf("can't reach the %s at %s (is %sexec running?)", peer, url, peer)
	}
	return err.Error()
}
