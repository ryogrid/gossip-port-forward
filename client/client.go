package client

import (
	"context"
	"fmt"
	"github.com/ryogrid/gossip-port-forward/gossip-overlay"
	"log"
	"net"

	"github.com/ryogrid/gossip-port-forward/util"
)

type ClientListen struct {
	Addr string
	Port uint16
}

type Client struct {
	node   gossip_overlay.Node
	listen ClientListen
}

func New(addr string, port uint16, listen ClientListen) *Client {
	node, err := gossip_overlay.New(addr, port)
	if err != nil {
		log.Fatalln(err)
	}

	return &Client{node, listen}
}

func (c *Client) ConnectAndSync(ctx context.Context, targetPeerId peer2.ID) {
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

			stream := c.node.OpenStreamToTargetPeer(peerId)

			go util.Sync(tcpConn, stream)
		}
	}()
}
