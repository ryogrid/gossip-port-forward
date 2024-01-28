package client

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/overlay"
	"github.com/weaveworks/mesh"
	"io"
	"log"
	"math"
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

			// read remote node address at start of stream wrote by application

			lenBuf := make([]byte, 1)
			// read 1 byte
			n, err3 := io.ReadFull(tcpConn, lenBuf)
			if err3 != nil || n != 1 {
				fmt.Println("DummyTCPListener::Accept failed (reading addres len)", err3)
			}
			addrStrLen := int(lenBuf[0])
			addrBuf := make([]byte, addrStrLen)
			// read addrStrLen bytes
			n, err4 := io.ReadFull(tcpConn, addrBuf)
			if err4 != nil || n != addrStrLen {
				fmt.Println("DummyTCPListener::Accept failed (reading address)", err4)
			}
			remoteAddrStr := string(addrBuf)

			targetPeerId_ := targetPeerId
			if targetPeerId == math.MaxUint64 {
				// destination peer can't be determinated at launch case (destination is not single peer)
				targetPeerId_ = mesh.PeerName(uint64(util.GenHashIDUint16(remoteAddrStr)))
			}

			if err2 != nil {
				log.Fatalln(err2)
			}

			stream := c.peer.OpenStreamToTargetPeer(targetPeerId_, remoteAddrStr)

			go util.Sync(tcpConn, stream)
		}
	}()
}
