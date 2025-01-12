package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"sync"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"golang.org/x/sync/semaphore"
)

func runIngest(esClient *elasticsearch.Client) error {
	page, err := GetPage(1)
	if err != nil {
		return err
	}

	log.Printf("Current total items: %d", page.Data.Pagination.Items)

	// drop index
	_, err = esClient.Indices.Delete([]string{"salt"})
	if err != nil {
		return err
	}

	indexConfig := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"title": map[string]interface{}{
					"type":     "text",
					"analyzer": "dutch",
				},
			},
		},
	}

	jsonConfig, err := json.Marshal(indexConfig)
	if err != nil {
		return err
	}

	// create index
	_, err = esClient.Indices.Create("salt",
		esClient.Indices.Create.WithBody(bytes.NewReader(jsonConfig)))
	if err != nil {
		return err
	}

	pagination := page.Data.Pagination

	wg := sync.WaitGroup{}
	sem := semaphore.NewWeighted(10)

	for i := 1; i <= pagination.LastPage; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			if err := sem.Acquire(context.Background(), 1); err != nil {
				log.Panicf("error %v", err)
			}
			defer sem.Release(1)

			page, err := GetPage(i)
			if err != nil {
				log.Panicf("error %v", err)
			}

			items, err := ParsePageContent(page.Data.Content)
			if err != nil {
				log.Panicf("error %v", err)
			}

			err = IngestItems(esClient, items)
			if err != nil {
				log.Panicf("error %v", err)
			}

			log.Printf("Page %d done with %d items", i, len(items))
		}(i)
	}

	wg.Wait()

	return nil
}
