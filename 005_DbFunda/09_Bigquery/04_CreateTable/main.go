package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/api/iterator"

	// Imports the Google Cloud BigQuery client package.
	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
)

type playerPoints struct {
	Name       string
	Points     int64
	LastPlayed time.Time
	T          time.Time
}

func (p *playerPoints) Save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"Name":       p.Name,
		"Count":      p.Points,
		"LastPlayed": p.LastPlayed,
		"T":          p.T,
	}, "", nil
}

func main() {
	// Need to have Authentication Credentials in Place
	// Read the "Readme.md" file for more
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	//projectID := "<YOUR-PROJECT-ID-HERE>"
	projectID := os.Getenv("PROJECT_ID")
	//   Sets the name for Required dataset.
	datasetName := "my_new_dataset"
	//   Table to be used for storage of Player scores
	tableName := "player_score"

	{
		// # Creates a client.
		client := getBigQueryClient(projectID, ctx)
		defer client.Close()
		// # Create Dataset if not present

		ds := initDataset(client, datasetName, ctx)
		log.Println(" Dataset Ready :", ds.DatasetID)

		// # Create Table

		//   Get an Instance of the Table
		tb := initTable(client, ds, tableName, ctx)

		log.Println(" Table Ready :", tb.TableID)
	}

	// # Adding Description to Table
	{
		client := getBigQueryClient(projectID, ctx)
		defer client.Close()

		ds := initDataset(client, datasetName, ctx)
		tb := initTable(client, ds, tableName, ctx)
		//   1. Get Older Metadata
		tbmeta, err := tb.Metadata(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		//   2. Create the new Fields to change
		newMeta := bigquery.TableMetadataToUpdate{
			Description: "Stores the points scored by each player and last played time",
		}

		//   3. Update metadata using the older ETag
		_, err = tb.Update(ctx, newMeta, tbmeta.ETag)
		if err != nil {
			log.Fatalln(err)
		}

		//   4. Fetch the latest Description and Display
		tbmeta, _ = tb.Metadata(ctx)
		log.Printf(" Table Description: %v", tbmeta.Description)
	}

	// # List Tables in current Dataset
	{
		client := getBigQueryClient(projectID, ctx)
		defer client.Close()

		ds := initDataset(client, datasetName, ctx)
		// # List Tables in current Dataset
		tbls := ds.Tables(ctx)
		log.Println(" Following are the tables in", ds.DatasetID)
		lines := 0
		for {
			t, err := tbls.Next()
			if err == iterator.Done {
				break
			} else if err != nil {
				log.Fatalln(err)
			}
			log.Println(" - ", t.TableID)
			lines++
		}
		if lines == 0 {
			log.Println("   Looks like there are not Tables in", ds.DatasetID)
		}
	}

	// # Add Data to the Table
	{
		client := getBigQueryClient(projectID, ctx)
		defer client.Close()

		ds := initDataset(client, datasetName, ctx)
		tb := initTable(client, ds, tableName, ctx)

		data := []playerPoints{
			{"Hari", 200, time.Now().AddDate(0, 0, -1), time.Now()},
			{"Dev", 430, time.Now().AddDate(0, 0, -3), time.Now()},
			{"Radhika", 230, time.Now().AddDate(0, 0, -1), time.Now()},
		}
		//   Get the Uploader Handle
		upl := tb.Uploader()
		//   Add the Data
		err := upl.Put(ctx, data)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("  Data Added successfully to the Table")
	}

	// # List out the Rows in the Table
	{
		client := getBigQueryClient(projectID, ctx)
		defer client.Close()

		ds := initDataset(client, datasetName, ctx)
		tb := initTable(client, ds, tableName, ctx)

		q := client.Query(fmt.Sprintf(`
		SELECT *
		FROM %s.%s
		LIMIT 1000
	`, ds.DatasetID, tb.TableID))
		//  Read Query
		it, err := q.Read(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(" Following are records in the", tb.TableID, " Table")
		for {
			var row []bigquery.Value
			err := it.Next(&row)
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(" - ", row)
		}
	}
}

func getBigQueryClient(projectID string, ctx context.Context) *bigquery.Client {
	// # Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func initDataset(client *bigquery.Client, datasetName string, ctx context.Context) *bigquery.Dataset {
	ds := client.Dataset(datasetName)
	_, err := ds.Metadata(ctx)

	//   Check if the Dataset Does not Exists
	if err != nil {
		if strings.Contains(err.Error(), "notFound") {
			err = ds.Create(ctx, &bigquery.DatasetMetadata{})
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(" Created the Dataset :", datasetName)
		} else {
			log.Fatalln(err)
		}
	}

	return ds
}

func initTable(client *bigquery.Client, ds *bigquery.Dataset, tableName string, ctx context.Context) *bigquery.Table {
	tb := ds.Table(tableName)
	_, err := tb.Metadata(ctx)
	//   Check if the Dataset Does not Exists
	if err != nil {
		if strings.Contains(err.Error(), "notFound") {

			// Prepare DB Schema
			schema, err := bigquery.InferSchema(playerPoints{})
			if err != nil {
				log.Fatalln(err)
			}
			//   Create the Database with Fields
			err = tb.Create(ctx, &bigquery.TableMetadata{Schema: schema})
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(" Created the table :", tableName)
		} else {
			log.Fatalln(err)
		}
	}
	return tb
}
