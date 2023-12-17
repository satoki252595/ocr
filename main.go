package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func main() {

	file := "sample.png"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	result := detectText(file)

	if result != nil {
		log.Fatal(result)
	}
}

// detectText gets text from the Vision API for an image at the given file path.
func detectText(file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Println("要素なし")
	} else {
		for _, annotation := range annotations {
			fmt.Printf("%q\n", annotation.Description)
		}
	}

	return nil
}
