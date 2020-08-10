package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strings"
)

type Elasticsearch struct {
	client *elasticsearch.Client
}

func NewEs(username, password string, urls ...string) *Elasticsearch {
	es := new(Elasticsearch)
	client, err := es.GetClient(username, password, urls)
	if err != nil {

	}
	es.client = client
	return es
}

func (es *Elasticsearch) GetClient(username, password string, urls []string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: urls,
		Username:  username,
		Password:  password,
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil, err
	}

	return esClient, nil
}

func (es *Elasticsearch) Search(index string, query map[string]interface{}) (*SearchResult, error) {
	// 构建请求体
	var buf bytes.Buffer
	var r map[string]interface{}
	result := SearchResult{}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// 搜索
	res, err := es.client.Search(
		es.client.Search.WithContext(context.Background()),
		es.client.Search.WithIndex(index),
		es.client.Search.WithBody(&buf),
		es.client.Search.WithTrackTotalHits(true),
		es.client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
			return nil, err
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return nil, err
	}
	b, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return nil, err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Fatalf("Error parsing map to struct: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))

	return &result, nil
}
