package nats

import "time"

// Connection is the struct who defines a connection in NATS server
type Connection struct {
	ID                 int64     `json:"cid"`
	IP                 string    `json:"ip"`
	Port               int64     `json:"port"`
	StartTime          time.Time `json:"start"`
	LastHeartbeat      time.Time `json:"last_activity"`
	RTT                string    `json:"rtt"`
	Uptime             string    `json:"uptime"`
	Idle               string    `json:"idle"`
	PendingBytes       int64     `json:"pending_bytes"`
	InputMessageCount  int64     `json:"in_msgs"`
	OutputMessageCount int64     `json:"out_msgs"`
	InputBytesCount    int64     `json:"in_bytes"`
	OutputBytesCount   int64     `json:"out_bytes"`
	Subscriptions      int32     `json:"subscriptions"`
	ClientLanguage     string    `json:"lang,omitempty"`
	ClientVersion      string    `json:"version,omitempty"`
}
