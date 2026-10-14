package main

import (
	"fmt"
	"io"
	"sort"
	"sync"
	"text/tabwriter"
)

// Table is the router's translation table. Each entry remembers which
// inside address a hardware address belongs to, so a reply addressed to the
// router can be handed back to the right client.
//
// Linux keeps the same kind of record in its connection-tracking table
// (conntrack); see the README for how the real one is keyed.
type Table struct {
	mu      sync.Mutex
	entries map[string]string // hardware address -> original "ip:port"
}

func NewTable() *Table {
	return &Table{entries: make(map[string]string)}
}

// Record files inside under hw, replacing any older entry for that
// hardware address.
func (t *Table) Record(hw, inside string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.entries[hw] = inside
}

// Lookup returns the inside address filed under hw.
func (t *Table) Lookup(hw string) (string, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	inside, ok := t.entries[hw]
	return inside, ok
}

// WriteTo prints the table as aligned columns, sorted by hardware address.
func (t *Table) WriteTo(w io.Writer) (int64, error) {
	t.mu.Lock()
	hws := make([]string, 0, len(t.entries))
	for hw := range t.entries {
		hws = append(hws, hw)
	}
	sort.Strings(hws)
	rows := make([]string, len(hws))
	for i, hw := range hws {
		rows[i] = fmt.Sprintf("%s\t%s\n", hw, t.entries[hw])
	}
	t.mu.Unlock()

	cw := &countingWriter{w: w}
	tw := tabwriter.NewWriter(cw, 0, 0, 3, ' ', 0)
	fmt.Fprint(tw, "HWADDR\tINSIDE ADDRESS\n")
	for _, row := range rows {
		fmt.Fprint(tw, row)
	}
	err := tw.Flush()
	return cw.n, err
}

type countingWriter struct {
	w io.Writer
	n int64
}

func (c *countingWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.n += int64(n)
	return n, err
}
