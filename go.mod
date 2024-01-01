module github.com/ryogrid/gossip-port-forward

go 1.8

replace github.com/ryogrid/gossip-overlay => ../gossip-overlay

require (
	github.com/ryogrid/gossip-overlay v0.0.0-00010101000000-000000000000
	//github.com/ryogrid/gossip-overlay v0.0.6
	github.com/spf13/cobra v1.8.0
	github.com/weaveworks/mesh v0.0.0-20191105120815-58dbcc3e8e63
)
