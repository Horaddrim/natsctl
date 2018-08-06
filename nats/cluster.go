package nats

// Cluster is the JSON definition for any NATS cluster
type Cluster struct {
	Address              string `json:"addr"`
	Port                 int64  `json:"cluster_port"`
	AuthorizationTimeout int32  `json:"auth_timeout"`
}
