package server

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/overlay"
	"github.com/ryogrid/gossip-port-forward/gossip-overlay"
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
	node    *gossip_overlay.Node
	forward ServerForward
	ID      mesh.PeerName
}

func New(forward ServerForward, gossipListenPort uint16) *Server {
	node, err := gossip_overlay.NewNode(nil, gossipListenPort)
	if err != nil {
		log.Fatalln(err)
	}

	return &Server{node, forward, node.Peer.GossipDataMan.Self}
}

func (s *Server) ListenAndSync() {
	defer func() {
		gossip_overlay.LoggerObj.Printf("mesh router stopping")
		s.node.Peer.Router.Stop()
	}()

	//errs := make(chan error)
	//
	//go func() {
	//	c := make(chan os.Signal)
	//	signal.Notify(c, syscall.SIGINT)
	//	errs <- fmt.Errorf("%s", <-c)
	//}()

	go func() {
		oserv, err := overlay.NewOverlayServer(s.node.Peer, s.node.Peer.GossipMM)
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
