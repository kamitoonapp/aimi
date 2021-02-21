package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
)

// Download public object
func DownloadFile(w io.Writer, bucket, object string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 50)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)

	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	
	fmt.Fprintf(w, "Blob %v downloaded.\n", object)
	
	return data, nil
}

func GetMetaData(w io.Writer, bucket, object string) (*storage.ObjectAttrs, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucket).Object(object)

	attrs, err := o.Attrs(ctx)

	if err != nil {
		return nil, fmt.Errorf("Object(%q).Attrs: %v", object, err)
	}

	return attrs, nil
}