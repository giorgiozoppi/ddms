package config

import "github.com/giorgiozoppi/dms/pkg/common"

// NodeType can be of types of nodes: SuperNode and StandardNode
type NodeType string

const (
	// SuperNode is a resolver node: it communicates its views to other nodes
	SuperNode NodeType = "SuperNode"
	// Node is a simple node
	Node NodeType = "StandardNode"
)

// NodeConfig is a node configuration
type NodeConfig struct {
	// PeerID is a peer identifier
	PeerID string
	// Hostname is a string identifier
	Hostname string
	// Port is unique integer
	Port int
	// BootPeers is a list of string
	BootPeers []string
	// PeerSecret is a secret key for peers
	PeerSecret string
	// Type is the type of a node
	Type NodeType
	// alive nodes
	AliveNodes common.NodeMap
	// dead nodes
	DeadNodes common.NodeMap
}
