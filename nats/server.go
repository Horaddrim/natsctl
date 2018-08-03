package nats

import "time"

// Server is the definition struct to the Server info object.
type Server struct {
	ID                string       `json:"server_id"`
	ServerTime        time.Time    `json:"now"`
	ActiveConnections int64        `json:"num_connections"`
	TotalConnections  int64        `json:"total"`
	Offset            int32        `json:"offset"`
	Limit             int32        `json:"limit"`
	Connections       []Connection `json:"connections"`
}
