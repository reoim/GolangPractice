// Used DefaultClient with service accont json key.

package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
)

func main() {

	// Sets your Google Cloud Platform project ID.
	projectID := "terraform-reo"
	zone := "us-central1-a"
	// Creates a client.
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	instacnesList, err := computeService.Instances.List(projectID, zone).Do()
	if err != nil {
		fmt.Printf("Failed to get list of instances: %v", err)
	}
	fmt.Printf("%+v\n", instacnesList.Items[0].Name)

}
