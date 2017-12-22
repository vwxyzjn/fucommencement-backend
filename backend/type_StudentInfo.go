package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
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
func (s *Server) AddEntry(st *StudentInfo) {
	algoliaObject := algoliasearch.Object(structs.Map(st))
	if _, err := s.AlgoliaIndex.AddObject(algoliaObject); err != nil {
		panic(err)
	}
}

func (s *Server) getEntryByFurmanID(id int) *StudentInfo {
	params := algoliasearch.Map{
		"restrictSearchableAttributes": []string{
			"furmanID",
		},
	}
	res, err := s.AlgoliaIndex.Search(strconv.Itoa(id), params)
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

func (s *Server) getEntryByID(id string) *StudentInfo {
	data, err := s.AlgoliaIndex.GetObject(id, nil)
	if err != nil {
		panic(err)
	}
	var studentData StudentInfo
	if err := mapstructure.Decode(data, &studentData); err != nil {
		panic(err)
	}
	return &studentData
}

func (s *Server) deleteEntryByFrumanID(id int) {
	studentData := s.getEntryByFurmanID(id)
	_, err := s.AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	DeleteFile("." + studentData.ProfilePicturePath)
	DeleteFile("." + studentData.NamePronunciationPath)
}

// DeleteEntryByID delete algolia entry by objectID
func (s *Server) DeleteEntryByID(id string) {
	studentData := s.getEntryByID(id)
	_, err := s.AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	DeleteFile("." + studentData.ProfilePicturePath)
	DeleteFile("." + studentData.NamePronunciationPath)
}

// DeleteEntryByIDPreserveFiles delete algolia entry but keep the files
func (s *Server) DeleteEntryByIDPreserveFiles(id string) {
	studentData := s.getEntryByID(id)
	_, err := s.AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	fmt.Println("I am called")
}

func (s *Server) GetSettings() {
	// sortReplicaIndex()
	settings, err := s.AlgoliaIndex.GetSettings()
	if err != nil {
		panic(err)
	}
	PrettyPrint(settings.ToMap())
}

func (s *Server) addReplica() {
	settings := algoliasearch.Map{
		"replicas": []string{"student_by_custom_sorting"},
	}
	if _, err := s.AlgoliaIndex.SetSettings(settings); err != nil {
		panic(err)
	}
}

// func (s *Server) sortReplicaIndex() {
// 	settings := algoliasearch.Map{
// 		"ranking": []string{"asc(diplomaLastName)"},
// 	}
// 	if _, err := s.AlgoliaSortedIndex.SetSettings(settings); err != nil {
// 		panic(err)
// 	}
// }

func (s *Server) Export() {
	fmt.Println("Export being called")

	var hits []algoliasearch.Map

	it, err := s.AlgoliaIndex.BrowseAll(algoliasearch.Map{"query": ""})
	if err != nil {
		panic(err)
	}

	res, err := it.Next()
	for err != algoliasearch.NoMoreHitsErr {
		if err != nil {
			panic(err)
		}

		hits = append(hits, res)
		res, err = it.Next()
	}

	json, err := json.Marshal(hits)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("your_filename.json", json, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func Test() {
	// studentData := getEntryByFurmanID(991596)
	// fmt.Println(studentData.FurmanID)
}

// NthEntryInIndex fetches the n-th ranked entry in the index
// Notice this methods assume the default ``hitsPerPage`` is 20
func (s *Server) NthEntryInIndex(n int) {
	fmt.Println("NthEntryInIndex")
	res, _ := s.AlgoliaIndex.Search("", algoliasearch.Map{
		"page": n / 20, // 20 is the default page
	})
	nthResult := res.Hits[n%20]
	var studentData StudentInfo
	if err := mapstructure.Decode(nthResult, &studentData); err != nil {
		panic(err)
	}
	return &studentData
}
