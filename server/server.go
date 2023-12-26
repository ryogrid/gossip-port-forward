package server

import (
	"fmt"
	"github.com/ryogrid/gossip-port-forward/gossip-overlay"
	"log"
	"net"
)

type ServerForward struct {
	Addr string
	Port uint16
}

type Server struct {
	peer    gossip_overlay.Node
	forward ServerForward
	ID      peer2.ID
}

func New(addr string, port uint16, forward ServerForward) *Server {
	node, err := gossip_overlay.New(addr, port)
	if err != nil {
		log.Fatalln(err)
	}

	return &Server{node, forward, node.Host.ID()}
}

func (s *Server) ListenAndSync() {
	//ctx := context.Background()

	log.Println("Announcing ourselves...")
	//s.peer.Advertise(ctx)
	log.Println("Successfully announced.")

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

	log.Println("Waiting for client to connect.\nYour PeerId is", s.ID.Pretty())
}

func (s *Server) dialForwardServer() (*net.TCPConn, error) {
	raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.forward.Addr, s.forward.Port))
	if err != nil {
		panic(err)
	}

	return net.DialTCP("tcp", nil, raddr)
}
