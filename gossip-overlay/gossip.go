package gossip_overlay

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/gossip"
	"github.com/ryogrid/gossip-overlay/overlay"
	"github.com/ryogrid/gossip-overlay/util"
	"github.com/ryogrid/gossip-port-forward/constants"
	"github.com/weaveworks/mesh"
	"io"
	"log"
	"math"
	"net"
	"os"
)

var LoggerObj *log.Logger

type Node struct {
	Peer *gossip.Peer
}

func NewNode(destPeerId *uint16, gossipListenPort uint16) (*Node, error) {
	nicAddr := util.MustHardwareAddr()
	name, err := mesh.PeerNameFromString(nicAddr)
	if err != nil {
		panic("Failed to get PeerName from NIC address")
	}

	meshConf := mesh.Config{
		Host:               "0.0.0.0",
		Port:               int(gossipListenPort),
		ProtocolMinVersion: mesh.ProtocolMaxVersion,
		Password:           nil,
		ConnLimit:          64,
		PeerDiscovery:      true,
		TrustedSubnets:     []*net.IPNet{},
	}

	LoggerObj = log.New(os.Stderr, "gossip> ", log.LstdFlags)
	emptyStr := ""
	meshListen := "local"
	var destPeerId_ uint64 = math.MaxUint64
	if destPeerId != nil {
		destPeerId_ = uint64(*destPeerId)
	}
	peers := &util.Stringset{}
	peers.Set(constants.BootstrapPeer)
	p := gossip.NewPeer(name, LoggerObj, mesh.PeerName(destPeerId_), &emptyStr, &emptyStr, &meshListen, &meshConf, peers)

	return &Node{p}, nil
}

func (node *Node) OpenStreamToTargetPeer(peerId mesh.PeerName) io.ReadWriteCloser {
	LoggerObj.Println(fmt.Sprintf("Opening a stream to %d", peerId))

	oc, err := overlay.NewOverlayClient(node.Peer, node.Peer.Destname, node.Peer.GossipMM)
	if err != nil {
		panic(err)
	}

	channel, streamID, err2 := oc.OpenChannel(math.MaxUint16)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(fmt.Sprintf("opened: %d", streamID))

	//passId := peer_.ID
	//stream, err := n.Host.NewStream(ctx, passId, constants.Protocol)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println("Opened a stream to", peer_.ID)
	//return stream
	return channel
}
