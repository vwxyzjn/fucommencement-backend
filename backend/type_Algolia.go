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

type AlgoliaInstance struct {
	AlgoliaClient algoliasearch.Client
	AlgoliaIndex  algoliasearch.Index
}

func (a *AlgoliaInstance) Initialize(AlgoliaAppID string, AlgoliaKey string, AlgoliaIndexName string) {
	a.AlgoliaClient = algoliasearch.NewClient(AlgoliaAppID, AlgoliaKey)
	a.AlgoliaIndex = a.AlgoliaClient.InitIndex(AlgoliaIndexName)
}

// AddEntry ..
func (a *AlgoliaInstance) AddEntry(s *StudentInfo) {
	algoliaObject := algoliasearch.Object(structs.Map(s))
	if _, err := a.AlgoliaIndex.AddObject(algoliaObject); err != nil {
		panic(err)
	}
}

func (a *AlgoliaInstance) getEntryByFurmanID(id int) *StudentInfo {
	params := algoliasearch.Map{
		"restrictSearchableAttributes": []string{
			"furmanID",
		},
	}
	res, err := a.AlgoliaIndex.Search(strconv.Itoa(id), params)
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

func (a *AlgoliaInstance) getEntryByID(id string) *StudentInfo {
	data, err := a.AlgoliaIndex.GetObject(id, nil)
	if err != nil {
		panic(err)
	}
	var studentData StudentInfo
	if err := mapstructure.Decode(data, &studentData); err != nil {
		panic(err)
	}
	return &studentData
}

func (a *AlgoliaInstance) deleteEntryByFrumanID(id int) {
	studentData := a.getEntryByFurmanID(id)
	_, err := a.AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	DeleteFile("." + studentData.ProfilePicturePath)
	DeleteFile("." + studentData.NamePronunciationPath)
}

// DeleteEntryByID delete algolia entry by objectID
func (a *AlgoliaInstance) DeleteEntryByID(id string) {
	studentData := a.getEntryByID(id)
	_, err := a.AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	DeleteFile("." + studentData.ProfilePicturePath)
	DeleteFile("." + studentData.NamePronunciationPath)
}

// DeleteEntryByIDPreserveFiles delete algolia entry but keep the files
func (a *AlgoliaInstance) DeleteEntryByIDPreserveFiles(id string) {
	studentData := a.getEntryByID(id)
	_, err := a.AlgoliaIndex.DeleteObject(studentData.ObjectID)
	if err != nil {
		panic(err)
	}
	fmt.Println("I am called")
}

func (a *AlgoliaInstance) GetSettings() {
	// sortReplicaIndex()
	settings, err := a.AlgoliaIndex.GetSettings()
	if err != nil {
		panic(err)
	}
	PrettyPrint(settings.ToMap())
}

func (a *AlgoliaInstance) addReplica() {
	settings := algoliasearch.Map{
		"replicas": []string{"student_by_custom_sorting"},
	}
	if _, err := a.AlgoliaIndex.SetSettings(settings); err != nil {
		panic(err)
	}
}

// func (a *AlgoliaInstance) sortReplicaIndex() {
// 	settings := algoliasearch.Map{
// 		"ranking": []string{"asc(diplomaLastName)"},
// 	}
// 	if _, err := a.AlgoliaSortedIndex.SetSettings(settings); err != nil {
// 		panic(err)
// 	}
// }

func (a *AlgoliaInstance) Export() {
	fmt.Println("Export being called")

	var hits []algoliasearch.Map

	it, err := a.AlgoliaIndex.BrowseAll(algoliasearch.Map{"query": ""})
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
