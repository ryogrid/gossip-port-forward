package relay

import (
	"github.com/ryogrid/gossip-port-forward/gossip-overlay"
	"github.com/weaveworks/mesh"
	"log"
)

type Relay struct {
	peer       *gossip_overlay.Node
	ID         mesh.PeerName
	gossipPort uint16
}

func New(gossipPort uint16) *Relay {
	node, err := gossip_overlay.NewNode()
	if err != nil {
		log.Fatalln(err)
	}

	return &Relay{node, node.Peer.GossipDataMan.Self, gossipPort}
}
