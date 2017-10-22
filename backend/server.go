package backend

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	profilePicturePath    = "./commencement/profilePicture/"
	namePronunciationPath = "./commencement/namePronunciation/"
)

// Setup setups the server http end points
func Setup() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", testGET)
	r.POST("/commencementPOST", commencementPOST)
	r.StaticFS("/commencement", http.Dir("./commencement"))
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
		Majors:                    c.PostForm("majors"),
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
	if namePronunciation, err := c.FormFile("namePronunciation"); err == nil {
		studentData.NamePronunciationPath = handleUpload(namePronunciation, studentData, namePronunciationPath)
	}
	if profilePicture, err := c.FormFile("profilePicture"); err == nil {
		studentData.ProfilePicturePath = handleUpload(profilePicture, studentData, profilePicturePath)
	}

	studentData.CreateEntry()
	c.String(http.StatusOK, fmt.Sprintf("File %s", studentData.Name))
	fmt.Println(c.PostForm("name"))

}

// handleUpload reads a file from c.PostForm and return its stored path
func handleUpload(file *multipart.FileHeader, student *StudentInfo, path string) string {
	extentionName := getFileExtension(file)
	fileName := student.Name + "-" + strconv.Itoa(student.FurmanID) + "." + extentionName
	multipartFile, err := file.Open()
	CheckErr(err)
	saveFile(multipartFile, path, fileName)
	return path + fileName
}

func getFileExtension(file *multipart.FileHeader) string {
	fileName := file.Filename
	fileNameSplit := strings.Split(fileName, ".")
	return fileNameSplit[len(fileNameSplit)-1]
}

func saveFile(file multipart.File, path string, fileName string) {
	// check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	// save the file
	dst, err := os.Create(path + fileName)
	CheckErr(err)
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		panic(err)
	}
}
