package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/bigquery"
)

func main() {
	// Need to have Authentication Credentials in Place
	// Read the "Readme.md" file for more
	ctx := context.Background()

	projectID := os.Getenv("PROJECT_ID")

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	log.Println(" The Datasets under the project", projectID, "are:")
	iter := client.Datasets(ctx)
	lines := 0
	for {
		ds, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		log.Println(" - ", ds.DatasetID)
		lines++
	}

	if lines == 0 {
		log.Println("  Looks like there are no Datasets available")
	}
}
