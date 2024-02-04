package constants

import "github.com/multiformats/go-multiaddr"

var BootstrapPeers = [1]multiaddr.Multiaddr{
	//multiaddr.StringCast("/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN"),
	//multiaddr.StringCast("/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa"),
	//multiaddr.StringCast("/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb"),
	//multiaddr.StringCast("/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt"),
	//multiaddr.StringCast("/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"),
	multiaddr.StringCast("/ip4/127.0.0.1/tcp/20001/p2p/Qmc5gskqBKrbqu8BSgtESAyZ8iC6PBn895jJFZxA8Lyj84"),
}

const Protocol = "/gossip-port-forward/v0"
