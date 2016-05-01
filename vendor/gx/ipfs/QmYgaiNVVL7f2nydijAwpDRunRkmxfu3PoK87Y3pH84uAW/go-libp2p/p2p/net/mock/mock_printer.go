package mocknet

import (
	"fmt"
	"io"

	inet "gx/ipfs/QmYgaiNVVL7f2nydijAwpDRunRkmxfu3PoK87Y3pH84uAW/go-libp2p/p2p/net"
	peer "gx/ipfs/QmZwZjMVGss5rqYsJVGy18gNbkTJffFyq2x1uJ4e4p3ZAt/go-libp2p-peer"
)

// separate object so our interfaces are separate :)
type printer struct {
	w io.Writer
}

func (p *printer) MocknetLinks(mn Mocknet) {
	links := mn.Links()

	fmt.Fprintf(p.w, "Mocknet link map:\n")
	for p1, lm := range links {
		fmt.Fprintf(p.w, "\t%s linked to:\n", peer.ID(p1))
		for p2, l := range lm {
			fmt.Fprintf(p.w, "\t\t%s (%d links)\n", peer.ID(p2), len(l))
		}
	}
	fmt.Fprintf(p.w, "\n")
}

func (p *printer) NetworkConns(ni inet.Network) {

	fmt.Fprintf(p.w, "%s connected to:\n", ni.LocalPeer())
	for _, c := range ni.Conns() {
		fmt.Fprintf(p.w, "\t%s (addr: %s)\n", c.RemotePeer(), c.RemoteMultiaddr())
	}
	fmt.Fprintf(p.w, "\n")
}
