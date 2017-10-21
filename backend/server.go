package backend

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

type studentInfo struct {
	name                      string
	furmanID                  int64
	anticipatedCompletionDate string
	degreeExpected            string
	majors                    string
	interdisciplinaryMinor    string
	diplomafirstName          string
	diplomamiddleName         string
	diplomalastName           string
	hometownAndState          string
	pronounceFirstName        string
	pronounceMiddleName       string
	pronounceLastName         string
	rhymeFirstName            string
	rhymeMiddleName           string
	rhymeLastName             string
	postGradAddress           string
	postGradAddressTwo        string
	postGradCity              string
	postGradState             string
	postGradPostalCode        string
	postGradTelephone         string
	postGradEmail             string
	intentConfirm             string
	namePronunciation         []byte
	profilePicture            []byte
}

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
	studentData := &studentInfo{
		name: c.PostForm("name"),
		anticipatedCompletionDate: c.PostForm("anticipatedCompletionDate"),
		degreeExpected:            c.PostForm("degreeExpected"),
		majors:                    c.PostForm("interdisciplinaryMinor"),
		diplomafirstName:          c.PostForm("diplomafirstName"),
		diplomamiddleName:         c.PostForm("diplomamiddleName"),
		diplomalastName:           c.PostForm("diplomalastName"),
		hometownAndState:          c.PostForm("hometownAndState"),
		pronounceFirstName:        c.PostForm("pronounceFirstName"),
		pronounceMiddleName:       c.PostForm("pronounceMiddleName"),
		pronounceLastName:         c.PostForm("pronounceLastName"),
		rhymeFirstName:            c.PostForm("rhymeFirstName"),
		rhymeMiddleName:           c.PostForm("rhymeMiddleName"),
		rhymeLastName:             c.PostForm("rhymeLastName"),
		postGradAddress:           c.PostForm("postGradAddress"),
		postGradAddressTwo:        c.PostForm("postGradAddressTwo"),
		postGradCity:              c.PostForm("postGradCity"),
		postGradState:             c.PostForm("postGradState"),
		postGradPostalCode:        c.PostForm("postGradPostalCode"),
		postGradTelephone:         c.PostForm("postGradTelephone"),
		postGradEmail:             c.PostForm("postGradEmail"),
		intentConfirm:             c.PostForm("intentConfirm"),
	}
	studentData.namePronunciation = readFile(c.FormFile("namePronunciation"))
	studentData.profilePicture = readFile(c.FormFile("namePronunciation"))
	studentData.furmanID, _ = strconv.ParseInt(c.PostForm("furmanID"), 10, 64)

	// SaveData(buf.Bytes())
	err := SaveData(studentData)
	CheckErr(err)
	c.String(http.StatusOK, fmt.Sprintf("File %s %d", reflect.TypeOf(studentData.namePronunciation), studentData.furmanID))
	fmt.Println(studentData.furmanID)

}

// readFile is a helper function to store images and sounds
func readFile(file *multipart.FileHeader, err error) []byte {
	multipartFile, err := file.Open()
	CheckErr(err)
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, multipartFile)
	return buf.Bytes()
}
