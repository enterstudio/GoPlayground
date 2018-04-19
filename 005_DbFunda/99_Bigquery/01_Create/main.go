package main

import (
	"fmt"
	"log"
	"os"

	// Imports the Google Cloud BigQuery client package.
	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
)

func main() {
	// Need to have Authentication Credentials in Place
	// Read the "Readme.md" file for more
	// 01_Create/Readme.md#getting-account-credentials
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	//projectID := "<YOUR-PROJECT-ID-HERE>"
	projectID := os.Getenv("PROJECT_ID")

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new dataset.
	datasetName := "my_new_dataset"

	// Creates the new BigQuery dataset.
	if err := client.Dataset(datasetName).Create(ctx, &bigquery.DatasetMetadata{}); err != nil {
		log.Fatalf("Failed to create dataset: %v", err)
	}

	fmt.Printf("Dataset created\n")
}
