package main

import (
	"fmt"
	"os"

	"gitlab.com/vwxyzjn/fucommencement-backend/backend"
)

func main() {
	fmt.Println(os.Getenv("APP_ENVIRONMENT"))

	var indexName = "student"
	if os.Getenv("APP_ENVIRONMENT") == "development" {
		indexName = "student_test"
	}
	var sortedIndexName = "student_by_custom_sorting"
	if os.Getenv("APP_ENVIRONMENT") == "development" {
		sortedIndexName = "student_by_custom_sorting_test"
	}

	server := &backend.Server{
		Port:                   ":8080",
		ProfilePicturePath:     "./commencement/profilePicture/",
		NamePronunciationPath:  "./commencement/namePronunciation/",
		AlgoliaAppID:           "TH20RENZY1",
		AlgoliaKey:             "f6fc0cc56e0b7af1fc5e5d71ff207bf6",
		AlgoliaIndexName:       indexName,
		AlgoliaSortedIndexName: sortedIndexName,
	}
	server.InitAlgolia()
	server.Migrate(indexName, sortedIndexName)
	// server.Setup()
}
