package main

import (
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	esClient, err := CreateEsClient()
	if err != nil {
		return err
	}
	_, err = esClient.Info()
	if err != nil {
		return err
	}

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [search|ingest]", os.Args[0])
		return nil
	}

	switch os.Args[1] {
	case "search":
		return runSearch(esClient)
	case "ingest":
		return runIngest(esClient)
	}

	log.Printf("Usage: %s [search|ingest]", os.Args[0])

	return nil
}
