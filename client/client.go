package client

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/overlay"
	"github.com/weaveworks/mesh"
	"log"
	"net"

	"github.com/ryogrid/gossip-port-forward/util"
)

type ClientListen struct {
	Addr string
	Port uint16
}

type Client struct {
	peer   *overlay.OverlayPeer
	listen ClientListen
}

// func New(destPeerId uint64, clientListen ClientListen, gossipListenPort uint16) *Client {
func New(peer *overlay.OverlayPeer, clientListen ClientListen) *Client {
	return &Client{peer, clientListen}
}

func (c *Client) ConnectAndSync(targetPeerId mesh.PeerName) {
	log.Println("Creating listen server...")
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.listen.Addr, c.listen.Port))
	if err != nil {
		log.Fatalln(err)
	}

	tcpLn, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Created listen server")

	log.Println("You can connect with", tcpLn.Addr().String())

	go func() {
		for {
			tcpConn, err2 := tcpLn.AcceptTCP()
			if err2 != nil {
				log.Fatalln(err2)
			}

			stream := c.peer.OpenStreamToTargetPeer(targetPeerId, "")

			go util.Sync(tcpConn, stream)
		}
	}()
}
