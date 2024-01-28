package relay

import (
	"github.com/ryogrid/gossip-overlay/overlay"
	"github.com/ryogrid/gossip-overlay/util"
	"github.com/ryogrid/gossip-port-forward/constants"
	"github.com/weaveworks/mesh"
	"log"
	"math/rand"
)

type Relay struct {
	peer       *overlay.OverlayPeer
	ID         mesh.PeerName
	gossipPort uint16
}

func New(gossipPort uint16) *Relay {
	//host := "0.0.0.0"
	peers := &util.Stringset{}
	peers.Set(constants.BootstrapPeer)
	peer, err := overlay.NewOverlayPeer(uint64(rand.Int31n(10000)), int(gossipPort), peers, false)
	if err != nil {
		log.Fatalln(err)
	}

	return &Relay{peer, peer.Peer.GossipDataMan.Self, gossipPort}
}
