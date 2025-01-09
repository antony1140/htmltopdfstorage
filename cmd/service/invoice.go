package service

import (
	"github.com/antony1140/htmltopdfstorage/data"
	"os"
	"fmt"
	"log"
)

func UploadInvoice(key string) (error) {
	path := "invoice/" + key
	client := data.InitS3()
	file, osErr := os.Open("./" + key)
	if osErr != nil {
		fmt.Println("error in upload invoice", osErr)
	return osErr
}
	err := data.UploadS3(client, file, path)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("file uploaded")
	return nil
}
func DownloadInvoice(key string) (error) {
	path := "invoice/" + key
	client := data.InitS3()
	downloadErr := data.DownloadS3(client, path)
	if downloadErr != nil {
		log.Println("error at download invoice", downloadErr)
		return downloadErr
	}
	return nil
}
	

