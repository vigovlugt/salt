package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Response struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source Thread  `json:"_source"`
			Score  float64 `json:"_score"`
		} `json:"hits"`
	} `json:"hits"`
}

func runSearch(esClient *elasticsearch.Client) error {
	type Query struct {
		Query struct {
			Match struct {
				Title string `json:"title"`
			} `json:"match"`
		} `json:"query"`
	}

	queryString := os.Args[2]
	query := Query{}
	query.Query.Match.Title = queryString

	queryJson, err := json.Marshal(query)
	if err != nil {
		return err
	}

	res, err := esClient.Search(
		esClient.Search.WithIndex("salt"),
		esClient.Search.WithBody(bytes.NewReader(queryJson)),
		esClient.Search.WithPretty(),
		esClient.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	visualizeSearchResults(response)

	return nil
}

func visualizeSearchResults(response Response) {
	log.Printf("Total hits: %d", response.Hits.Total.Value)

	t := table.NewWriter()
	t.AppendHeader(table.Row{"ID", "Title", "Temp", "Url", "Score"})
	for _, hit := range response.Hits.Hits {
		t.AppendRow(table.Row{hit.Source.ThreadID, hit.Source.Title, hit.Source.Temperature, hit.Source.ShareableLink, hit.Score})
	}

	fmt.Println(t.Render())
}
