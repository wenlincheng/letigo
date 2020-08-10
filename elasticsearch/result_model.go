package elasticsearch

type SearchResult struct {
	Took     float64 `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   Shard   `json:"_shards"`
	Hits     Hits    `json:"hits"`
}

type Shard struct {
	Total      float64 `json:"total"`
	Successful float64 `json:"successful"`
	Skipped    float64 `json:"skipped"`
	Failed     float64 `json:"failed"`
}

type Hits struct {
	Total    HitsTotal `json:"total"`
	MaxScore float64   `json:"max_score"`
	Hits     []Hit     `json:"hits"`
}

type Hit struct {
	Id      string                 `json:"_id"`
	Index   string                 `json:"_index"`
	Score   float64                `json:"_score"`
	Doctype string                 `json:"_type"`
	Source  map[string]interface{} `json:"_source"`
}

type HitsTotal struct {
	Relation string  `json:"relation"`
	Value    float64 `json:"value"`
}
