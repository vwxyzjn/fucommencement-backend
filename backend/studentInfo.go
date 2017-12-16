package backend

import (
	"fmt"
	"os"
	"strconv"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

// Client, Index
var (
	AlgoliaClient = algoliasearch.NewClient("TH20RENZY1", "f6fc0cc56e0b7af1fc5e5d71ff207bf6")
	AlgoliaIndex  = AlgoliaClient.InitIndex("student")
)

// StudentInfo ..
type StudentInfo struct {
	ObjectID                  string `form:"objectID" json:"objectID" structs:"objectID"`
	Name                      string `form:"name" json:"name" structs:"name" binding:"required"`
	FurmanID                  int    `form:"furmanID" json:"furmanID" structs:"furmanID" binding:"required"`
	AnticipatedCompletionDate string `form:"anticipatedCompletionDate" json:"anticipatedCompletionDate" structs:"anticipatedCompletionDate" binding:"required"`
	DegreeExpected            string `form:"degreeExpected" json:"degreeExpected" structs:"degreeExpected" binding:"required"`
	Majors                    string `form:"majors" json:"majors" structs:"majors" binding:"required"`
	InterdisciplinaryMinor    string `form:"interdisciplinaryMinor" json:"interdisciplinaryMinor" structs:"interdisciplinaryMinor" binding:"required"`
	DiplomaFirstName          string `form:"diplomaFirstName" json:"diplomaFirstName" structs:"diplomaFirstName" binding:"required"`
	DiplomaMiddleName         string `form:"diplomaMiddleName" json:"diplomaMiddleName" structs:"diplomaMiddleName" binding:"required"`
	DiplomaLastName           string `form:"diplomaLastName" json:"diplomaLastName" structs:"diplomaLastName" binding:"required"`
	HometownAndState          string `form:"hometownAndState" json:"hometownAndState" structs:"hometownAndState" binding:"required"`
	PronounceFirstName        string `form:"pronounceFirstName" json:"pronounceFirstName" structs:"pronounceFirstName" binding:"required"`
	PronounceMiddleName       string `form:"pronounceMiddleName" json:"pronounceMiddleName" structs:"pronounceMiddleName" binding:"required"`
	PronounceLastName         string `form:"pronounceLastName" json:"pronounceLastName" structs:"pronounceLastName" binding:"required"`
	RhymeFirstName            string `form:"rhymeFirstName" json:"rhymeFirstName" structs:"rhymeFirstName" binding:"required"`
	RhymeMiddleName           string `form:"rhymeMiddleName" json:"rhymeMiddleName" structs:"rhymeMiddleName" binding:"required"`
	RhymeLastName             string `form:"rhymeLastName" json:"rhymeLastName" structs:"rhymeLastName" binding:"required"`
	PostGradAddress           string `form:"postGradAddress" json:"postGradAddress" structs:"postGradAddress" binding:"required"`
	PostGradAddressTwo        string `form:"postGradAddressTwo" json:"postGradAddressTwo" structs:"postGradAddressTwo" binding:"required"`
	PostGradCity              string `form:"postGradCity" json:"postGradCity" structs:"postGradCity" binding:"required"`
	PostGradState             string `form:"postGradState" json:"postGradState" structs:"postGradState" binding:"required"`
	PostGradPostalCode        string `form:"postGradPostalCode" json:"postGradPostalCode" structs:"postGradPostalCode" binding:"required"`
	PostGradTelephone         string `form:"postGradTelephone" json:"postGradTelephone" structs:"postGradTelephone" binding:"required"`
	PostGradEmail             string `form:"postGradEmail" json:"postGradEmail" structs:"postGradEmail" binding:"required"`
	IntentConfirm             string `form:"intentConfirm" json:"intentConfirm" structs:"intentConfirm" binding:"required"`
	NamePronunciationPath     string `form:"namePronunciationPath" json:"namePronunciationPath" structs:"namePronunciationPath" `
	ProfilePicturePath        string `form:"profilePicturePath" json:"profilePicturePath" structs:"profilePicturePath" `
	Honor                     string `form:"honor" json:"honor" structs:"honor" `
}

// AddEntry ..
func (s *StudentInfo) AddEntry() {
	algoliaObject := algoliasearch.Object(structs.Map(s))
	if _, err := AlgoliaIndex.AddObject(algoliaObject); err != nil {
		panic(err)
	}
}

func getEntryByFurmanID(id int) *StudentInfo {
	params := algoliasearch.Map{
		"restrictSearchableAttributes": []string{
			"furmanID",
		},
	}
	res, err := AlgoliaIndex.Search(strconv.Itoa(id), params)
	if err != nil {
		panic(err)
	}
	data := res.Hits[0]
	var studentData StudentInfo
	if err := mapstructure.Decode(data, &studentData); err != nil {
		panic(err)
	}
	return &studentData
}

func getEntryByID(id string) *StudentInfo {
	data, err := AlgoliaIndex.GetObject(id, nil)
	if err != nil {
		panic(err)
	}
	var studentData StudentInfo
	if err := mapstructure.Decode(data, &studentData); err != nil {
		panic(err)
	}
	return &studentData
}

func DeleteEntryByFrumanID(id int) {
	studentData := getEntryByFurmanID(id)
	_, err := AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	deleteFile("." + studentData.ProfilePicturePath)
	deleteFile("." + studentData.NamePronunciationPath)
}

func DeleteEntryByID(id string) {
	studentData := getEntryByID(id)
	_, err := AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	deleteFile("." + studentData.ProfilePicturePath)
	deleteFile("." + studentData.NamePronunciationPath)
}

func DeleteEntryByIDPreservePicture(id string) {
	studentData := getEntryByID(id)
	_, err := AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	fmt.Println("I am called")
}

func deleteFile(path string) {
	// check if file exist first.
	fmt.Println(path)
	if path == "." {
		return
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := os.Remove(path)
		if err != nil {
			panic(err)
		}
	}
}

func Test() {
	studentData := getEntryByFurmanID(991596)
	fmt.Println(studentData.FurmanID)
}
