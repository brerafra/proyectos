package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Article struct {
	Title           string `csv:"Title"`
	URL             string `csv:"URL"`
	DocumentTags    string `csv:"Document tags"`
	SavedDate       string `csv:"Saved date"`
	ReadingProgress string `csv:"Reading progress"`
	Location        string `csv:"Location"`
	Seen            string `csv:"Seen"`
}

func ReadCsv() []*Article {
	csvFile, csvFileError := os.OpenFile("example.csv", os.O_RDWR, os.ModePerm)
	if csvFileError != nil {
		panic(csvFileError)
	}
	defer csvFile.Close()

	var articles []*Article
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &articles); unmarshalError != nil {
		panic(unmarshalError)
	}
	return articles
}

func GetInboxArticles(articles []*Article) []*Article {
	var inboxArticles []*Article

	for _, article := range articles {
		if article.Location == "inbox" {
			inboxArticles = append(inboxArticles, article)
		}
	}
	return inboxArticles
}

func WriteCsv(articles []*Article) {
	resultFile, resultFileError := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if resultFileError != nil {
		panic(resultFileError)
	}
	defer resultFile.Close()

	if marshalError := gocsv.MarshalFile(&articles, resultFile); marshalError != nil {
		panic(marshalError)
	}
}

func main() {
	articles := ReadCsv()
	filteredArticles := GetInboxArticles(articles)
	WriteCsv(filteredArticles)
}
