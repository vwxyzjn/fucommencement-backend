package main

import (
	"gitlab.com/vwxyzjn/fucommencement-backend/backend"
)

func main() {
	server := &backend.Server{
		ProfilePicturePath:    "./commencement/profilePicture/",
		NamePronunciationPath: "./commencement/namePronunciation/",
		AlgoliaAppID:          "TH20RENZY1",
		AlgoliaKey:            "f6fc0cc56e0b7af1fc5e5d71ff207bf6",
		AlgoliaIndexName:      "student",
	}
	// server.Migrate()
	// backend.Export()
	// backend.GetSettings()
	server.Setup()
	// r.Run() // listen and serve on 0.0.0.0:8080
}
