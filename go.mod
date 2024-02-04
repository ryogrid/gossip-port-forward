module github.com/ryogrid/gossip-port-forward

go 1.8

replace github.com/ryogrid/gossip-port-forward => ./

require (
	github.com/libp2p/go-libp2p v0.30.0
	github.com/libp2p/go-libp2p-kad-dht v0.25.1
	github.com/multiformats/go-multiaddr v0.11.0
	github.com/spf13/cobra v1.1.3
)
