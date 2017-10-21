package backend

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Setup setups the server http end points
func Setup() {
	r := gin.Default()
	r.GET("/ping", testGET)
	r.POST("/commencementPOST", commencementPOST)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func testGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func commencementPOST(c *gin.Context) {
	studentData := &StudentInfo{
		Name: c.PostForm("name"),
		AnticipatedCompletionDate: c.PostForm("anticipatedCompletionDate"),
		DegreeExpected:            c.PostForm("degreeExpected"),
		Majors:                    c.PostForm("Majors"),
		InterdisciplinaryMinor:    c.PostForm("interdisciplinaryMinor"),
		DiplomaFirstName:          c.PostForm("diplomafirstName"),
		DiplomaMiddleName:         c.PostForm("diplomamiddleName"),
		DiplomaLastName:           c.PostForm("diplomalastName"),
		HometownAndState:          c.PostForm("hometownAndState"),
		PronounceFirstName:        c.PostForm("pronounceFirstName"),
		PronounceMiddleName:       c.PostForm("pronounceMiddleName"),
		PronounceLastName:         c.PostForm("pronounceLastName"),
		RhymeFirstName:            c.PostForm("rhymeFirstName"),
		RhymeMiddleName:           c.PostForm("rhymeMiddleName"),
		RhymeLastName:             c.PostForm("rhymeLastName"),
		PostGradAddress:           c.PostForm("postGradAddress"),
		PostGradAddressTwo:        c.PostForm("postGradAddressTwo"),
		PostGradCity:              c.PostForm("postGradCity"),
		PostGradState:             c.PostForm("postGradState"),
		PostGradPostalCode:        c.PostForm("postGradPostalCode"),
		PostGradTelephone:         c.PostForm("postGradTelephone"),
		PostGradEmail:             c.PostForm("postGradEmail"),
		IntentConfirm:             c.PostForm("intentConfirm"),
	}
	temp, _ := strconv.ParseInt(c.PostForm("furmanID"), 10, 64)
	studentData.FurmanID = int(temp)
	studentData.NamePronunciation = readFile(c.FormFile("namePronunciation"))
	studentData.ProfilePicture = readFile(c.FormFile("namePronunciation"))

	studentData.CreateEntry()
	c.String(http.StatusOK, fmt.Sprintf("File %s %d", reflect.TypeOf(studentData.NamePronunciation), studentData.FurmanID))
	fmt.Println(studentData.FurmanID)

}

// readFile is a helper function to store images and sounds
func readFile(file *multipart.FileHeader, err error) []byte {
	multipartFile, err := file.Open()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, multipartFile)
	return buf.Bytes()
}
