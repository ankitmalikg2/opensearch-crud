package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/opensearch-project/opensearch-go/opensearchutil"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
)

func main() {
	// Replace with your OpenSearch cluster details
	endpoint := "https://localhost:9200"
	username := "admin" // Leave empty if not using authentication
	password := "admin" // Leave empty if not using authentication

	// Create a client
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{endpoint},
		Username:  username,
		Password:  password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	if err != nil {
		fmt.Println("Error creating OpenSearch client:", err)
		return
	}

	// Index a document
	indexName := "dev-article"
	documentID := "1"
	document := map[string]interface{}{
		"title":   "Getting Started with OpenSearch",
		"content": "OpenSearch is a powerful open-source search and analytics engine...",
	}
	err = indexDocument(client, indexName, documentID, document)
	if err != nil {
		fmt.Println("Error indexing document:", err)
		return
	}
	fmt.Println("Document indexed:", documentID)
}

func indexDocument(client *opensearch.Client, indexName string, documentID string, document map[string]interface{}) error {

	_, err := client.Create(indexName, documentID, opensearchutil.NewJSONReader(document))

	return err
}
