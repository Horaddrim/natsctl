package nats

import (
	"fmt"
	"time"
)

const (
	// ConnectionRoute is the route to get connections
	ConnectionRoute = "connz"
	// VariableRoute is the route to get the server status
	VariableRoute = "varz"
	// NatsRoutesRoute is the route to get the NATS server routes
	NatsRoutesRoute = "routez"
	// SubscriptionsRoute is the route to get the NATS subscriptions details
	SubscriptionsRoute = "subsz"
)

var (
	// DefaultHeaders is the default header for the table
	DefaultHeaders []string
)

func init() {
	DefaultHeaders = []string{"SERVER", "TIMESTAMP", "AGE", "MEMORY USAGE"}
}

// Server is the definition struct to the Server info object.
type Server struct {
	ID                       string    `json:"server_id"`
	ServerTime               time.Time `json:"now"`
	Version                  string    `json:"version"`
	Proto                    int32     `json:"proto"`
	CommitVersion            string    `json:"git_commit"`
	GoRuntimeVersion         string    `json:"go"`
	Host                     string    `json:"host"`
	Address                  string    `json:"addr"`
	MaxConnection            int64     `json:"max_connections"`
	PingInterval             int64     `json:"ping_interval"`
	MaxPingCount             int32     `json:"ping_max"`
	HTTPHost                 string    `json:"http_host"`
	HTTPPort                 int64     `json:"http_port"`
	HTTPSPort                int64     `json:"https_port"`
	AuthorizationTimeout     int64     `json:"auth_timeout"`
	MaxControlLine           int64     `json:"max_control_line"`
	ClusterInfo              Cluster   `json:"cluster"`
	TLSTimeout               float32   `json:"tls_timeout"`
	Port                     int64     `json:"port"`
	PayloadMaximumSize       int64     `json:"max_payload"`
	StartTime                time.Time `json:"start"`
	Age                      string    `json:"uptime"`
	MemoryUsage              int64     `json:"mem"`
	CoresUsage               int32     `json:"cores"`
	CPUsUsing                int32     `json:"cpu"`
	ActiveConnections        int32     `json:"connections"`
	TotalConnections         int64     `json:"total_connections"`
	RoutesCount              int32     `json:"routes"`
	Remotes                  int32     `json:"remotes"`
	IncomingMessagesCount    int64     `json:"in_msgs"`
	OutcomingMessagesCount   int64     `json:"out_msgs"`
	IncomingBytesCount       int64     `json:"in_bytes"`
	OutcomingBytesCount      int64     `json:"out_bytes"`
	SlowConsumersCount       int32     `json:"slow_consumers"`
	MaxPending               int64     `json:"max_pending"`
	WriteDeadline            int64     `json:"write_deadline"`
	SubscriptionsCount       int32     `json:"subscriptions"`
	ConfigurationLoadingTime time.Time `json:"config_load_time"`
}

// PrepareHost receives the raw data of the host and port, and prepare it to
// be connected
func PrepareHost(protocol, rawHost, rawPort, route string) (finalURL string) {
	baseHost := fmt.Sprintf("%s://%s", protocol, rawHost)

	hostWithPort := fmt.Sprintf("%s:%s", baseHost, rawPort)

	finalURL = fmt.Sprintf("%s/%s", hostWithPort, route)
	return
}
