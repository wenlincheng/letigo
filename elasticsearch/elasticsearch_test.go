package elasticsearch

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	es := NewEs("", "", "http://localhost:9200")
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "C语言",
			},
		},
	}
	searchResult, err := es.Search("blog", query)
	if err != nil {
		return
	}

	fmt.Printf("%v", searchResult)

}
