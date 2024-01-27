package cmd

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/overlay"
	util2 "github.com/ryogrid/gossip-overlay/util"
	"github.com/ryogrid/gossip-port-forward/constants"
	"github.com/ryogrid/gossip-port-forward/relay"
	"github.com/weaveworks/mesh"
	"log"
	"os"
	"strconv"

	"github.com/ryogrid/gossip-port-forward/client"
	"github.com/ryogrid/gossip-port-forward/server"
	"github.com/ryogrid/gossip-port-forward/util"
	"github.com/spf13/cobra"
)

var gossipPort uint16 = 9999
var listenPort uint16
var forwardPort uint16
var forwardAddress string
var connectTo string
var selfPeerId uint16

var rootCmd = &cobra.Command{
	Use: "gossip-port-forward",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gossip-port-forward v0.1.0")
	},
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Startup client peer.",
	Run: func(cmd *cobra.Command, args []string) {
		listen := client.ClientListen{
			Addr: "127.0.0.1",
			Port: listenPort,
		}

		//destPeerId, err := strconv.ParseUint(connectTo, 10, 64) //Atoi(connectTo)
		//if err != nil {
		//	panic("Could not parse connect-to")
		//}
		//host := "0.0.0.0"
		peers := &util2.Stringset{}
		peers.Set(constants.BootstrapPeer)
		peer, err := overlay.NewOverlayPeer(uint64(selfPeerId), int(listenPort+1000), peers)
		if err != nil {
			log.Fatalln(err)
		}

		c := client.New(peer, listen)

		destNameNum, err := strconv.ParseUint(connectTo, 10, 64)
		if err != nil {
			panic("Could not parse Destname")
		}

		c.ConnectAndSync(mesh.PeerName(destNameNum))

		util.OSInterrupt()
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Startup server peer.",
	Run: func(cmd *cobra.Command, args []string) {
		forward := server.ServerForward{
			Addr: forwardAddress,
			Port: forwardPort,
		}

		peers := &util2.Stringset{}
		peers.Set(constants.BootstrapPeer)
		peer, err := overlay.NewOverlayPeer(uint64(selfPeerId), int(listenPort+1000), peers)
		if err != nil {
			log.Fatalln(err)
		}

		s := server.New(peer, forward)
		s.ListenAndSync()

		util.OSInterrupt()
	},
}

var bothCmd = &cobra.Command{
	Use:   "both",
	Short: "Startup client and server peer.",
	Run: func(cmd *cobra.Command, args []string) {
		// client initialization
		listen := client.ClientListen{
			Addr: "127.0.0.1",
			Port: listenPort,
		}

		peers := &util2.Stringset{}
		peers.Set(constants.BootstrapPeer)
		peer, err := overlay.NewOverlayPeer(uint64(selfPeerId), int(listenPort+1000), peers)
		if err != nil {
			log.Fatalln(err)
		}

		c := client.New(peer, listen)

		// server initialization
		forward := server.ServerForward{
			Addr: forwardAddress,
			Port: forwardPort,
		}
		s := server.New(peer, forward)

		// client and server launch

		destNameNum, err := strconv.ParseUint(connectTo, 10, 64)
		if err != nil {
			panic("Could not parse Destname")
		}
		c.ConnectAndSync(mesh.PeerName(destNameNum))

		s.ListenAndSync()

		util.OSInterrupt()
	},
}

var relayCmd = &cobra.Command{
	Use:   "relay",
	Short: "Startup relay peer.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = relay.New(gossipPort)

		util.OSInterrupt()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//overlay_setting.OVERLAY_DEBUG = true

	cobra.OnInitialize()

	clientCmd.Flags().Uint16VarP(
		&listenPort,
		"listen-port",
		"l",
		2222,
		"Listen server port",
	)
	clientCmd.Flags().StringVarP(
		&connectTo,
		"connect-to",
		"c",
		"",
		"PeerId of the server gossip peer",
	)
	clientCmd.MarkFlagRequired("connect-to")

	serverCmd.Flags().Uint16VarP(
		&forwardPort,
		"forward-port",
		"f",
		22,
		"Port to forward",
	)
	serverCmd.Flags().StringVarP(
		&forwardAddress,
		"forward-address",
		"a",
		"localhost",
		"Address to forward",
	)

	bothCmd.Flags().Uint16VarP(
		&selfPeerId,
		"self-peer-id",
		"s",
		1,
		"Peer ID of myself",
	)
	bothCmd.MarkFlagRequired("self-peer-id")
	bothCmd.Flags().AddFlagSet(clientCmd.Flags())
	bothCmd.Flags().AddFlagSet(serverCmd.Flags())

	relayCmd.Flags().Uint16VarP(
		&gossipPort,
		"gossip-port",
		"p",
		9999,
		"gossip relay peer port",
	)

	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(bothCmd)
	rootCmd.AddCommand(relayCmd)
}
