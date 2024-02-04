package server

import (
	"context"
	"fmt"
	"log"
	"net"

	network2 "github.com/libp2p/go-libp2p/core/network"
	peer2 "github.com/libp2p/go-libp2p/core/peer"
	"github.com/studiokaiji/libp2p-port-forward/constants"
	"github.com/studiokaiji/libp2p-port-forward/libp2p"
	"github.com/studiokaiji/libp2p-port-forward/util"
)

type ServerForward struct {
	Addr string
	Port uint16
}

type Server struct {
	node    libp2p.Node
	forward ServerForward
	ID      peer2.ID
}

func New(addr string, port uint16, forward ServerForward) *Server {
	node, err := libp2p.New(addr, port)
	if err != nil {
		log.Fatalln(err)
	}

	return &Server{node, forward, node.Host.ID()}
}

func (s *Server) ListenAndSync() {
	ctx := context.Background()

	log.Println("Announcing ourselves...")
	s.node.Advertise(ctx)
	log.Println("Successfully announced.")

	s.node.Host.SetStreamHandler(constants.Protocol, func(stream network2.Stream) {
		log.Println("Got a new stream!")

		log.Println("Connecting forward server...")

		tcpConn, err := s.dialForwardServer()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Connected forward server.")
		go util.Sync(tcpConn, stream)
	})

	log.Println("Waiting for client to connect.\nYour PeerId is", s.ID.Pretty())
}

func (s *Server) ListenAndSyncForRelay() {
	ctx := context.Background()

	s.node.AdvertiseForRelay(ctx)
	fmt.Println("Successfully finished AdvertiseForRelay method call.")

	//s.node.SetStreamHandler(constants.Protocol, func(stream network.Stream) {
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

	log.Println("Waiting other peers to connect.\nYour PeerId is", s.ID.Pretty())
}

func (s *Server) dialForwardServer() (*net.TCPConn, error) {
	raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.forward.Addr, s.forward.Port))
	if err != nil {
		panic(err)
	}

	return net.DialTCP("tcp", nil, raddr)
}
