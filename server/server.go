package server

import (
	"fmt"
	"github.com/ryogrid/gossip-port-forward/gossip-overlay"
	"github.com/weaveworks/mesh"
	"log"
	"net"
)

type ServerForward struct {
	Addr string
	Port uint16
}

type Server struct {
	peer    *gossip_overlay.Node
	forward ServerForward
	ID      mesh.PeerName
}

func New(forward ServerForward) *Server {
	node, err := gossip_overlay.NewNode()
	if err != nil {
		log.Fatalln(err)
	}

	return &Server{node, forward, node.Peer.GossipDataMan.Self}
}

func (s *Server) ListenAndSync() {
	// TODO: not implemented yet (Server::ListenAndSync)

	//s.peer.Host.SetStreamHandler(constants.Protocol, func(stream network2.Stream) {
	//	log.Println("Got a new stream!")
	//
	//	log.Println("Connecting forward server...")
	//
	//	tcpConn, err := s.dialForwardServer()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	log.Println("Connected forward server.")
	//	go util.Sync(tcpConn, stream)
	//})

	log.Println("Waiting for client to connect.\nYour PeerId is", s.ID)
}

func (s *Server) dialForwardServer() (*net.TCPConn, error) {
	raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.forward.Addr, s.forward.Port))
	if err != nil {
		panic(err)
	}

	return net.DialTCP("tcp", nil, raddr)
}
