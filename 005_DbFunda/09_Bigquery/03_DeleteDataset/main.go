package main

import (
	"context"
	"log"
	"os"

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

	// Sets the name for the new dataset.
	datasetName := "my_new_dataset"

	if err := client.Dataset(datasetName).Delete(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Println(" Deleted Dataset '", datasetName, "' from project", projectID)
}
