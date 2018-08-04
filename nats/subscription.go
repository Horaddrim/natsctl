package nats

// Subscription is the struct to holds information of the subscriptions
// in the NATS server
type Subscription struct {
	Count         int64   `json:"num_subscriptions"`
	CacheCount    int64   `json:"num_cache"`
	InsertsCount  int64   `json:"num_inserts"`
	RemovesCount  int64   `json:"num_removes"`
	MatchesCount  int64   `json:"num_matches"`
	CacheHitRate  float64 `json:"cache_hit_rate"`
	MaximunFanout int32   `json:"max_fanout"`
	AverageFanout int32   `json:"avg_fanout"`
}
