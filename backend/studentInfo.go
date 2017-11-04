package backend

import (
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
	ObjectID                  string
	Name                      string
	FurmanID                  int
	AnticipatedCompletionDate string
	DegreeExpected            string
	Majors                    string
	InterdisciplinaryMinor    string
	DiplomaFirstName          string
	DiplomaMiddleName         string
	DiplomaLastName           string
	HometownAndState          string
	PronounceFirstName        string
	PronounceMiddleName       string
	PronounceLastName         string
	RhymeFirstName            string
	RhymeMiddleName           string
	RhymeLastName             string
	PostGradAddress           string
	PostGradAddressTwo        string
	PostGradCity              string
	PostGradState             string
	PostGradPostalCode        string
	PostGradTelephone         string
	PostGradEmail             string
	IntentConfirm             string
	NamePronunciationPath     string
	ProfilePicturePath        string
	Honor                     string
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
		"honor":                     "",
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

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func Test() {
	// studentData := getEntryByFurmanID(991596)
	// fmt.Println(studentData.FurmanID)
	// fmt.Println()
}
