package server

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/overlay"
	util2 "github.com/ryogrid/gossip-overlay/util"
	"github.com/ryogrid/gossip-port-forward/constants"
	"github.com/ryogrid/gossip-port-forward/util"
	"github.com/weaveworks/mesh"
	"log"
	"net"
)

type ServerForward struct {
	Addr string
	Port uint16
}

type Server struct {
	peer    *overlay.OverlayPeer
	forward ServerForward
	ID      mesh.PeerName
}

func New(forward ServerForward, gossipListenPort uint16) *Server {
	host := "0.0.0.0"
	peers := &util2.Stringset{}
	peers.Set(constants.BootstrapPeer)
	peer, err := overlay.NewOverlayPeer(&host, int(gossipListenPort), peers)
	if err != nil {
		log.Fatalln(err)
	}

	return &Server{peer, forward, peer.Peer.GossipDataMan.Self}
}

func (s *Server) ListenAndSync() {
	defer func() {
		fmt.Println("mesh router stopping")
		//s.peer.Peer.Router.Stop()
		s.peer.Peer.Stop()
	}()

	go func() {
		oserv, err := overlay.NewOverlayServer(s.peer.Peer, s.peer.Peer.GossipMM)
		if err != nil {
			panic(err)
		}

		for {
			channel, remotePeerName, streamID, err2 := oserv.Accept()
			if err2 != nil {
				panic(err2)
			}
			fmt.Println("accepted:", remotePeerName, streamID)

			go func(channel_ *overlay.OverlayStream) {
				log.Println("Got a new stream!")

				log.Println("Connecting forward server...")

				tcpConn, err3 := s.dialForwardServer()
				if err3 != nil {
					log.Fatalln(err3)
				}

				log.Println("Connected forward server.")
				util.Sync(tcpConn, channel_)
			}(channel)
		}
	}()

	log.Printf("Waiting for client to connect.\nYour PeerId is %d\n", s.ID)
	//gossip_overlay.LoggerObj.Print(<-errs)
}

func (s *Server) dialForwardServer() (*net.TCPConn, error) {
	raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.forward.Addr, s.forward.Port))
	if err != nil {
		panic(err)
	}

	return net.DialTCP("tcp", nil, raddr)
}
