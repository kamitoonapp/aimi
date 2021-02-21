package main

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

func init() {
	ctx = context.Background()

	var err error

	client, err = storage.NewClient(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connection to GCP Storage Established !")
}