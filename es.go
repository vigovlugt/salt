package main

import (
	"bytes"
	"encoding/json"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

func CreateEsClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	return elasticsearch.NewClient(cfg)
}

func IngestItems(client *elasticsearch.Client, items []Thread) error {
	for _, item := range items {
		body, err := json.Marshal(item)
		if err != nil {
			return err
		}
		_, err = client.Index(
			"salt",
			bytes.NewReader(body),
			client.Index.WithDocumentID(item.ThreadID),
			client.Index.WithRefresh("true"),
		)
		if err != nil {
			return err
		}

	}
	return nil
}
