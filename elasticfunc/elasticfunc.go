package elasticfunc

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

var es *elasticsearch.Client

func Connect() (*elasticsearch.Client, error) {
	var err error
	es, err = elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	fmt.Println(elasticsearch.Version)
	fmt.Println(es.Info())

	return es, nil
}
