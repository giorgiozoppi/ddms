package algorithm

const (
	PASTRY_DEFAULT_ID_SIZE                                       = 16 // byte
	PASTRY_DEFAULT_USE_LEAF_SET                                  = true
	PASTRY_DEFAULT_LEAF_SET_ONE_SIDE_SIZE                        = 4
	PASTRY_DEFAULT_DO_PERIODIC_ROUTING_TABLE_MAINTENANCE         = true
	PASTRY_DEFAULT_ROUTING_TABLE_MAINTENANCE_INTERVAL            = 20 * 1000
	PASTRY_DEFAULT_ROUTING_TABLE_MAINTENANCE_INTERVAL_PLAY_RATIO = 0.5
	PASTRY_DEFAULT_UPDATE_LEAF_SET_FREQ                          = 3
)

type RoutingAlgorithmConfiguration interface {
}
type Algorithm struct {
	Name          string
	DefaultConfig RoutingAlgorithmConfiguration
}

func NewRoutingAlgorithmConfiguration() *RoutingAlgorithmConfiguration {
	return nil
}

func newPastryConfiguration() *RoutingAlgorithmConfiguration {
	return nil
}
