package backend

import (
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

// Client, Index
var (
	AlgoliaClient = algoliasearch.NewClient("TH20RENZY1", "f6fc0cc56e0b7af1fc5e5d71ff207bf6")
	AlgoliaIndex  = AlgoliaClient.InitIndex("student")
)

// StudentInfo ..
type StudentInfo struct {
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
		"diplomafirstName":          s.DiplomaFirstName,
		"diplomamiddleName":         s.DiplomaMiddleName,
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
