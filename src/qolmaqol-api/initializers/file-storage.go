package initializers

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
	"runtime"
)

func ConnectStorage() (*storage.Client, error) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	config, err := LoadConfig(".")
	credentialsPath := filepath.Join(basepath, config.FirebaseStorageFile)

	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatalf("Failed to create Firebase storage client: %v", err)
	}
	return client, nil
}
