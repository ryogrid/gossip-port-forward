# gossip-port-forward

- gossip-port-forward is command-line utility to transfer port between two hosts via different network / subnet peer-to-peer using gossip-overlay lib
  - gossip-port-forward is a fork of [studiokaiji/libp2p-port-forward](https://github.com/studiokaiji/libp2p-port-forward)

## Installation

WORK IN PROGRESS....

## Usage

```
Usage:
  gossip-port-forward [flags]
  gossip-port-forward [command]

Available Commands:
  help        Help about any command
  client      Startup client peer.  
  server      Startup server peer.
  relay       Startup relay peer.

Flags:
  -h, --help   help for gossip-port-forward
```

### Server

```
Usage:
  gossip-port-forward server [flags]

Flags:
  -h, --help                     help for server
  -a, --forward-address string   Address to forward (default "localhost")
  -f, --forward-port uint16      Port to forward (default 22)
```

### Client

```
Usage:
  gossip-port-forward client [flags]

Flags:
  -h, --help                 help for client
  -c, --connect-to string    PeerId of the server peer
  -l, --listen-port uint16   Listen server port (default 2222)
```

### Relay

```
Usage:
  gossip-port-forward relay [flags]

Flags:
  -h, --help                 help for relay
  -p, --gossip-port uint16   Relay listen port (default 9999)
```
- Now, relay node is running at ryogrid.net:9999 and its address is hard coded as bootstrap node for demonstration purpose
- So, you don't need to launch relay node by myself though data transfer is done without encryption
- If you want use self relay node, you should replace hard coded relay node address and run the relay node at your machine which is accessible from the Internet
