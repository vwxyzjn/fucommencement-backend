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
	r.GET("/deleteEntryGET/:objectID", deleteEntryGET)
	r.POST("/commencementPOST", commencementPOST)
	r.POST("/updateEntryPOST", updateEntryPOST)
	r.StaticFS("/commencement", http.Dir("./commencement"))
	r.Run() // listen and serve on 0.0.0.0:8080
}

func testGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func commencementPOST(c *gin.Context) {
	var form StudentInfo
	if err := c.Bind(&form); err == nil {
		if namePronunciation, err := c.FormFile("namePronunciation"); err == nil {
			form.NamePronunciationPath = handleUpload(namePronunciation, &form, namePronunciationPath)
		}
		if profilePicture, err := c.FormFile("profilePicture"); err == nil {
			form.ProfilePicturePath = handleUpload(profilePicture, &form, profilePicturePath)
		}
		form.AddEntry()
		c.String(http.StatusOK, fmt.Sprintf("File %s", form.Name))
		fmt.Println(c.PostForm("name"))

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// temp, _ := strconv.Atoi(c.PostForm("furmanID"))
	// studentData.FurmanID = int(temp)
}

func deleteEntryGET(c *gin.Context) {
	objectID := c.Param("objectID")
	fmt.Println("deleted", objectID)
	DeleteEntryByID(objectID)
	c.String(http.StatusOK, "ok")
}

func updateEntryPOST(c *gin.Context) {
	fmt.Println(c.PostForm("objectID"))
	studentData := &StudentInfo{
		ObjectID: c.PostForm("objectID"),
		Name:     c.PostForm("name"),
		AnticipatedCompletionDate: c.PostForm("anticipatedCompletionDate"),
		DegreeExpected:            c.PostForm("degreeExpected"),
		Majors:                    c.PostForm("majors"),
		InterdisciplinaryMinor:    c.PostForm("interdisciplinaryMinor"),
		DiplomaFirstName:          c.PostForm("diplomaFirstName"),
		DiplomaMiddleName:         c.PostForm("diplomaMiddleName"),
		DiplomaLastName:           c.PostForm("diplomaLastName"),
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
		NamePronunciationPath:     c.PostForm("namePronunciationPath"),
		ProfilePicturePath:        c.PostForm("profilePicturePath"),
		Honor:                     c.PostForm("honor"),
	}
	temp, _ := strconv.Atoi(c.PostForm("furmanID"))
	studentData.FurmanID = int(temp)
	DeleteEntryByIDPreservePicture(studentData.ObjectID)
	studentData.AddEntry()
	c.String(http.StatusOK, fmt.Sprintf("File %s", studentData.Name))
}

// handleUpload reads a file from c.PostForm and return its stored path
func handleUpload(file *multipart.FileHeader, student *StudentInfo, path string) string {
	extentionName := getFileExtension(file)
	fileName := student.Name + "-" + strconv.Itoa(student.FurmanID) + "." + extentionName
	multipartFile, err := file.Open()
	if err != nil {
		panic(err)
	}
	saveFile(multipartFile, path, fileName)
	return path[1:] + fileName
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
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		panic(err)
	}
}
