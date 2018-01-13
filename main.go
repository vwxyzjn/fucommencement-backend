package main

import (
	"gitlab.com/vwxyzjn/fucommencement-backend/backend"
)

func main() {
	server := &backend.Server{
		Port:                   ":8080",
		ProfilePicturePath:     "./commencement/profilePicture/",
		NamePronunciationPath:  "./commencement/namePronunciation/",
		AlgoliaAppID:           "TH20RENZY1",
		AlgoliaKey:             "f6fc0cc56e0b7af1fc5e5d71ff207bf6",
		AlgoliaIndexName:       "student",
		AlgoliaSortedIndexName: "student_by_custom_sorting",
	}

	server.Setup()
}
