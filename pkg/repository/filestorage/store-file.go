package filestorage

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	filePath = "../pkg/repository/filestorage/url-data.txt"
)

func StoreInFile(shortUrl, originalUrl string) {
	fmt.Printf("Request received by StoreInFile: shortUrl: %v originalUrl:%v\n", shortUrl, originalUrl)
	modifiedData := shortUrl + " " + originalUrl

	fileHandler, errorResponse := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
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

func FetchUrlFromFile(shortUrl string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make(map[string]string, 1024*4)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineRead := scanner.Text()
		stringsSlice := strings.Split(lineRead, " ")
		data[stringsSlice[0]] = stringsSlice[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data[shortUrl], err
}
