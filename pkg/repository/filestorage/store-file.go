package filestorage

import (
	"fmt"
	"os"
)

func StoreInFile(shortUrl, originalUrl string) {
	fmt.Printf("Request received by StoreInFile: shortUrl: %v originalUrl:%v\n", shortUrl, originalUrl)
	modifiedData := shortUrl + ":" + originalUrl

	fileHandler, errorResponse := os.OpenFile("../pkg/repository/filestorage/url-data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if errorResponse != nil {
		fmt.Println("StoreInFile | Error while opening file", errorResponse)
		return
	}
	noOfWrittenBytes, errorResponse := fmt.Fprintln(fileHandler, modifiedData)
	fmt.Println("No. of bytes written", noOfWrittenBytes)
	if errorResponse != nil {
		fmt.Println("StoreInFile | Error while writing to file: %v", errorResponse)
		fileHandler.Close()
		return
	}
	errorResponse = fileHandler.Close()
	if errorResponse != nil {
		fmt.Println("StoreInFile | Error while closing file connection", errorResponse)
	}
	fmt.Println("Request successfully processed by StoreInFile")
}
