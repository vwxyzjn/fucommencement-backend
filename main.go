package main

import (
	"gitlab.com/vwxyzjn/fucommencement-backend/backend"
)

func main() {
	server := &backend.Server{
		ProfilePicturePath:    "./commencement/profilePicture/",
		NamePronunciationPath: "./commencement/namePronunciation/",
	}
	replica := &backend.Server{
		ProfilePicturePath:    "./commencement/profilePicture/",
		NamePronunciationPath: "./commencement/namePronunciation/",
	}

	replica.Initialize(
		"TH20RENZY1",
		"f6fc0cc56e0b7af1fc5e5d71ff207bf6",
		"student_by_custom_sorting",
	)
	// replica.Export()
	replica.NthEntryInIndex(2)

	server.Initialize(
		"TH20RENZY1",
		"f6fc0cc56e0b7af1fc5e5d71ff207bf6",
		"student",
	)
	server.Setup()
}
