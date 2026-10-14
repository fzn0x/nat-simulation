// Package packet defines the simulated packet that the client, router and
// server pass to each other.
//
// All three programs really run on localhost, so the kernel would report
// 127.0.0.1 as everyone's address. To make the masquerade visible, the
// "network" addresses live inside the packet itself as plain fields. The
// router rewrites those fields exactly the way a NAT router rewrites the
// IP header.
package packet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// Packet is one simulated datagram.
type Packet struct {
	// Src is the sender's address as "ip:port".
	Src string `json:"src"`
	// Dst is the receiver's address as "ip:port".
	Dst string `json:"dst"`
	// HWAddr is the hardware (MAC-style) address of the machine that sent
	// the original request. The router files each translation under it,
	// and the server copies it into its reply so the router can find the
	// translation again on the way back.
	HWAddr string `json:"hwaddr"`
	// Body is the payload.
	Body string `json:"body"`
}

// Reply builds the answer to p: source and destination swap places, the
// hardware address is carried back unchanged.
func (p Packet) Reply(body string) Packet {
	return Packet{Src: p.Dst, Dst: p.Src, HWAddr: p.HWAddr, Body: body}
}

func (p Packet) String() string {
	return fmt.Sprintf("src=%-21s dst=%-21s hw=%s body=%q", p.Src, p.Dst, p.HWAddr, p.Body)
}

// Host returns the IP part of an "ip:port" address. If addr has no port it
// is returned unchanged.
func Host(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return addr
	}
	return host
}

// Port returns the port part of an "ip:port" address, or "" if there is none.
func Port(addr string) string {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return ""
	}
	return port
}

// Decode reads one packet from r.
func Decode(r io.Reader) (Packet, error) {
	var p Packet
	if err := json.NewDecoder(r).Decode(&p); err != nil {
		return Packet{}, fmt.Errorf("decode packet: %w", err)
	}
	if p.Src == "" || p.Dst == "" {
		return Packet{}, fmt.Errorf("decode packet: src and dst are required")
	}
	return p, nil
}

// Write sends p as an HTTP response with the given status code.
func Write(w http.ResponseWriter, status int, p Packet) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(p)
}

var client = &http.Client{Timeout: 3 * time.Second}

// Send delivers p to the program listening at url and returns its reply
// packet and HTTP status code.
func Send(url string, p Packet) (Packet, int, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return Packet{}, 0, err
	}
	resp, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return Packet{}, 0, err
	}
	defer resp.Body.Close()
	reply, err := Decode(resp.Body)
	if err != nil {
		return Packet{}, resp.StatusCode, err
	}
	return reply, resp.StatusCode, nil
}
