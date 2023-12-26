# gossip-port-forward

gossip-port-forward command-line utility to transfer port between two hosts via different network / subnet peer-to-peer using gossip-overlay lib.

## Installation

WORK IN PROGRESS....

## Usage

```
Usage:
  gossip-port-forward [flags]
  gossip-port-forward [command]

Available Commands:
  client      Startup client node.
  help        Help about any command
  server      Startup server node.

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
  -c, --connect-to string    PeerId of the server libp2p node
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
