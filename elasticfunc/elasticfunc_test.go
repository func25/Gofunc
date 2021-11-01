package elasticfunc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	r  map[string]interface{}
	wg sync.WaitGroup
)

type Student struct {
	Name         string  `json:"name"`
	Age          string  `json:"age"`
	AverageScore float64 `json:"averageScore"`
}

func init() {
	if _, err := Connect(); err != nil {
		log.Fatal("Shiet")
	}
}

func TestOthers(t *testing.T) {
	res, err := es.Cluster.Health()
	if err != nil {
		t.Error(err)
		return
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		t.Error(err)
		return
	}
	// fmt.Println(r)
	fmt.Println(res.String())
}

func TestCreateIndex(t *testing.T) {
}

func TestMe(t *testing.T) {
	res, err := es.Info()
	if err != nil {
		t.Error(err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		t.Error(res.String())
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		t.Error(err)
		return
	}

	fmt.Println("X:", elasticsearch.Version, r["version"].(map[string]interface{})["number"])

	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)
		go handleIndexRequest(i, title)
		break
	}
	wg.Wait()
	log.Println(strings.Repeat("-", 37))
}

type TestTitle struct {
	Title string `json:"title"`
}

func handleIndexRequest(i int, title string) error {
	defer wg.Done()

	// Build the request body.
	var b strings.Builder
	b.WriteString(`{"title" : "`)
	b.WriteString(title)
	b.WriteString(`"}`)

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "test",
		Body:    strings.NewReader(b.String()),
		Refresh: "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	return nil
}

func TestSearch(t *testing.T) {
	if err := handleSearchRequest(); err != nil {
		t.Error(err)
		return
	}
}

func handleSearchRequest() error {
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			return fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return err
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
	return nil
}
