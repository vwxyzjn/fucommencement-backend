package backend

import (
	"fmt"
	"os"
	"strconv"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

// Client, Index
var (
	AlgoliaClient = algoliasearch.NewClient("TH20RENZY1", "f6fc0cc56e0b7af1fc5e5d71ff207bf6")
	AlgoliaIndex  = AlgoliaClient.InitIndex("student")
)

// StudentInfo ..
type StudentInfo struct {
	ObjectID                  string `form:"objectID" json:"objectID"`
	Name                      string `form:"name" json:"name" binding:"required"`
	FurmanID                  int    `form:"furmanID" json:"furmanID" binding:"required"`
	AnticipatedCompletionDate string `form:"anticipatedCompletionDate" json:"anticipatedCompletionDate" binding:"required"`
	DegreeExpected            string `form:"degreeExpected" json:"degreeExpected" binding:"required"`
	Majors                    string `form:"majors" json:"majors" binding:"required"`
	InterdisciplinaryMinor    string `form:"interdisciplinaryMinor" json:"interdisciplinaryMinor" binding:"required"`
	DiplomaFirstName          string `form:"diplomaFirstName" json:"diplomaFirstName" binding:"required"`
	DiplomaMiddleName         string `form:"diplomaMiddleName" json:"diplomaMiddleName" binding:"required"`
	DiplomaLastName           string `form:"diplomaLastName" json:"diplomaLastName" binding:"required"`
	HometownAndState          string `form:"hometownAndState" json:"hometownAndState" binding:"required"`
	PronounceFirstName        string `form:"pronounceFirstName" json:"pronounceFirstName" binding:"required"`
	PronounceMiddleName       string `form:"pronounceMiddleName" json:"pronounceMiddleName" binding:"required"`
	PronounceLastName         string `form:"pronounceLastName" json:"pronounceLastName" binding:"required"`
	RhymeFirstName            string `form:"rhymeFirstName" json:"rhymeFirstName" binding:"required"`
	RhymeMiddleName           string `form:"rhymeMiddleName" json:"rhymeMiddleName" binding:"required"`
	RhymeLastName             string `form:"rhymeLastName" json:"rhymeLastName" binding:"required"`
	PostGradAddress           string `form:"postGradAddress" json:"postGradAddress" binding:"required"`
	PostGradAddressTwo        string `form:"postGradAddressTwo" json:"postGradAddressTwo" binding:"required"`
	PostGradCity              string `form:"postGradCity" json:"postGradCity" binding:"required"`
	PostGradState             string `form:"postGradState" json:"postGradState" binding:"required"`
	PostGradPostalCode        string `form:"postGradPostalCode" json:"postGradPostalCode" binding:"required"`
	PostGradTelephone         string `form:"postGradTelephone" json:"postGradTelephone" binding:"required"`
	PostGradEmail             string `form:"postGradEmail" json:"postGradEmail" binding:"required"`
	IntentConfirm             string `form:"intentConfirm" json:"intentConfirm" binding:"required"`
	NamePronunciationPath     string `form:"namePronunciationPath" json:"namePronunciationPath"`
	ProfilePicturePath        string `form:"profilePicturePath" json:"profilePicturePath"`
	Honor                     string `form:"honor" json:"honor"`
}

// AddEntry ..
func (s *StudentInfo) AddEntry() {
	studentData := algoliasearch.Object{
		"name":                      s.Name,
		"furmanID":                  s.FurmanID,
		"anticipatedCompletionDate": s.AnticipatedCompletionDate,
		"degreeExpected":            s.DegreeExpected,
		"majors":                    s.Majors,
		"interdisciplinaryMinor":    s.InterdisciplinaryMinor,
		"diplomaFirstName":          s.DiplomaFirstName,
		"diplomaMiddleName":         s.DiplomaMiddleName,
		"diplomaLastName":           s.DiplomaLastName,
		"hometownAndState":          s.HometownAndState,
		"pronounceFirstName":        s.PronounceFirstName,
		"pronounceMiddleName":       s.PronounceMiddleName,
		"pronounceLastName":         s.PronounceLastName,
		"rhymeFirstName":            s.RhymeFirstName,
		"rhymeMiddleName":           s.RhymeMiddleName,
		"rhymeLastName":             s.RhymeLastName,
		"postGradAddress":           s.PostGradAddress,
		"postGradAddressTwo":        s.PostGradAddressTwo,
		"postGradCity":              s.PostGradCity,
		"postGradState":             s.PostGradState,
		"postGradPostalCode":        s.PostGradPostalCode,
		"postGradTelephone":         s.PostGradTelephone,
		"postGradEmail":             s.PostGradEmail,
		"intentConfirm":             s.IntentConfirm,
		"namePronunciationPath":     s.NamePronunciationPath,
		"profilePicturePath":        s.ProfilePicturePath,
		"honor":                     s.Honor,
	}
	if _, err := AlgoliaIndex.AddObject(studentData); err != nil {
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
	studentData := &StudentInfo{
		ObjectID:                  data["objectID"].(string),
		Name:                      data["name"].(string),
		FurmanID:                  int(data["furmanID"].(float64)),
		AnticipatedCompletionDate: data["anticipatedCompletionDate"].(string),
		DegreeExpected:            data["degreeExpected"].(string),
		Majors:                    data["majors"].(string),
		InterdisciplinaryMinor:    data["interdisciplinaryMinor"].(string),
		DiplomaFirstName:          data["diplomaFirstName"].(string),
		DiplomaMiddleName:         data["diplomaMiddleName"].(string),
		DiplomaLastName:           data["diplomaLastName"].(string),
		HometownAndState:          data["hometownAndState"].(string),
		PronounceFirstName:        data["pronounceFirstName"].(string),
		PronounceMiddleName:       data["pronounceMiddleName"].(string),
		PronounceLastName:         data["pronounceLastName"].(string),
		RhymeFirstName:            data["rhymeFirstName"].(string),
		RhymeMiddleName:           data["rhymeMiddleName"].(string),
		RhymeLastName:             data["rhymeLastName"].(string),
		PostGradAddress:           data["postGradAddress"].(string),
		PostGradAddressTwo:        data["postGradAddressTwo"].(string),
		PostGradCity:              data["postGradCity"].(string),
		PostGradState:             data["postGradState"].(string),
		PostGradPostalCode:        data["postGradPostalCode"].(string),
		PostGradTelephone:         data["postGradTelephone"].(string),
		PostGradEmail:             data["postGradEmail"].(string),
		IntentConfirm:             data["intentConfirm"].(string),
		NamePronunciationPath:     data["namePronunciationPath"].(string),
		ProfilePicturePath:        data["profilePicturePath"].(string),
		Honor:                     data["honor"].(string),
	}

	return studentData
}

func getEntryByID(id string) *StudentInfo {
	data, err := AlgoliaIndex.GetObject(id, nil)
	if err != nil {
		panic(err)
	}
	studentData := &StudentInfo{
		ObjectID:                  data["objectID"].(string),
		Name:                      data["name"].(string),
		FurmanID:                  int(data["furmanID"].(float64)),
		AnticipatedCompletionDate: data["anticipatedCompletionDate"].(string),
		DegreeExpected:            data["degreeExpected"].(string),
		Majors:                    data["majors"].(string),
		InterdisciplinaryMinor:    data["interdisciplinaryMinor"].(string),
		DiplomaFirstName:          data["diplomaFirstName"].(string),
		DiplomaMiddleName:         data["diplomaMiddleName"].(string),
		DiplomaLastName:           data["diplomaLastName"].(string),
		HometownAndState:          data["hometownAndState"].(string),
		PronounceFirstName:        data["pronounceFirstName"].(string),
		PronounceMiddleName:       data["pronounceMiddleName"].(string),
		PronounceLastName:         data["pronounceLastName"].(string),
		RhymeFirstName:            data["rhymeFirstName"].(string),
		RhymeMiddleName:           data["rhymeMiddleName"].(string),
		RhymeLastName:             data["rhymeLastName"].(string),
		PostGradAddress:           data["postGradAddress"].(string),
		PostGradAddressTwo:        data["postGradAddressTwo"].(string),
		PostGradCity:              data["postGradCity"].(string),
		PostGradState:             data["postGradState"].(string),
		PostGradPostalCode:        data["postGradPostalCode"].(string),
		PostGradTelephone:         data["postGradTelephone"].(string),
		PostGradEmail:             data["postGradEmail"].(string),
		IntentConfirm:             data["intentConfirm"].(string),
		NamePronunciationPath:     data["namePronunciationPath"].(string),
		ProfilePicturePath:        data["profilePicturePath"].(string),
		Honor:                     data["honor"].(string),
	}

	return studentData
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
	// studentData := getEntryByFurmanID(991596)
	// fmt.Println(studentData.FurmanID)
	// fmt.Println()
}
