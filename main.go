package main

import (
	"gitlab.com/vwxyzjn/fucommencement-backend/backend"
)

func main() {
	server := &backend.Server{
		ProfilePicturePath:    "./commencement/profilePicture/",
		NamePronunciationPath: "./commencement/namePronunciation/",
	}
	// server.Migrate()
	// backend.Export()
	// backend.GetSettings()
	server.Setup(
		"TH20RENZY1",
		"f6fc0cc56e0b7af1fc5e5d71ff207bf6",
		"student",
	)
	// r.Run() // listen and serve on 0.0.0.0:8080
}
