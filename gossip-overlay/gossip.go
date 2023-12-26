package gossip_overlay

import (
	"github.com/ryogrid/gossip-overlay/gossip"
	"github.com/weaveworks/mesh"
	"io"
	"log"
)

type Node struct {
	Peer *gossip.Peer
}

func New(addr string, port uint16) (*Node, error) {
	//strAddr := fmt.Sprintf("/ip4/%s/tcp/%d", addr, port)
	//listenAddr := libp2p.ListenAddrStrings(strAddr)
	//
	//var DefaultPeerstore libp2p.Option = func(cfg *libp2p.Config) error {
	//	ps, err := pstoremem.NewPeerstore()
	//	if err != nil {
	//		return err
	//	}
	//
	//	return cfg.Apply(libp2p.Peerstore(ps))
	//}
	//
	//node, err := libp2p.New(DefaultPeerstore, listenAddr)

	return nil, nil
}

func (node *Node) OpenStreamToTargetPeer(peerId mesh.PeerName) io.ReadWriteCloser {
	log.Println("Opening a stream to", peerId)

	//passId := peer_.ID
	//stream, err := n.Host.NewStream(ctx, passId, constants.Protocol)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println("Opened a stream to", peer_.ID)

	//return stream
	return nil
}
