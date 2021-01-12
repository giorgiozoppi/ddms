package routing

import (
	"time"

	"github.com/giorgiozoppi/ddms/pkg/common"
)

// Hop is a hop in a route.
type Hop struct {
	hop  common.ID
	time time.Time
}

// Result is the result vector of the routing
type Result struct {
	NodeAndNeigbours []common.ID
	Route            Hop
}

// MessagingAddress is tuple that contains address, port, hostname.
type MessagingAddress struct {
	HostAddress string
	HostName    string
	Port        int16
}

// Service is a service used by the node for joing the DHT and routing nodes.
type Service interface {
	Join(initalPeer MessagingAddress) (Result, error)
	Leave() error
	Route(target common.ID, candidates int) (Result, error)
}
