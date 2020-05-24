package main

import (
	"context"
	"fmt"

	vision "cloud.google.com/go/vision/apiv1"
)

// Annonate an image file based on Cloud Vision API, return score and error if exists.
func annotate(uri string) (float32, error) {
	//Creates a client
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return 0.0, err
	}
	defer client.Close()

	image := vision.NewImageFromURI(uri)
	annotates, err := client.DetectFaces(ctx, image, nil, 1)
	if err != nil {
		return 0.0, err
	}

	if len(annotates) == 0 {
		fmt.Println("No faces found.")
		return 0.0, nil
	}
	return annotates[0].DetectionConfidence, nil
}
