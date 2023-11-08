package storage

import (
	"context"
	"fmt"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirebaseBucket(keyPath string, name string) (*storage.BucketHandle, error) {
	// get credentials
	opt := option.WithCredentialsFile(keyPath)

	// initialize app
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return nil, err
	}

	// get storage client
	client, err := app.Storage(context.Background())
	if err != nil {
		fmt.Printf("error getting storage client: %v", err)
		return nil, err
	}

	// get storage bucket
	bucket, err := client.Bucket(name)

	return bucket, err
}
