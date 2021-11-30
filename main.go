package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// Initialize minio client object.
	clt, err := minio.New("storage.googleapis.com", &minio.Options{
		Creds:  credentials.NewStaticV4("GOOGFLJHDAY4MHDNRLKN", "NV04ooirgR4cI3KIoQJcLSjl6ccFExNCGsVWcVR1", ""),
		Secure: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Upload the zip file
	objectName := "test.txt"
	filePath := "test.txt"
	contentType := "application/text"

	// Upload the zip file with FPutObject
	n, err := clt.FPutObject(context.Background(), "arribada-smart", objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s  %+v\n", objectName, n)
}
