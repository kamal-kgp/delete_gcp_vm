package main

import (
	"context"
	"fmt"
	"io"
	"os"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb" 
)

// deleteInstance sends a delete request to the Compute Engine API and waits for it to complete.
func deleteInstance(w io.Writer, projectID, zone, instanceName string) error {
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
			return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()  

	req := &computepb.DeleteInstanceRequest{
			Project:  projectID,
			Zone:     zone,
			Instance: instanceName,
	}

	op, err := instancesClient.Delete(ctx, req)
	if err != nil {
			return fmt.Errorf("unable to delete instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
			return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Instance deleted\n")

	return nil 
}

func main(){
	projectID  := "cloudsec-390404" ;
	zone := "us-central1-a" ;
	instanceName := "delete-test-instance" ;
	w := os.Stdout ;
	if err := deleteInstance(w, projectID, zone, instanceName); err != nil{
		fmt.Fprintf(w, "Failed to delete instance: %v", err) ;
	}
}