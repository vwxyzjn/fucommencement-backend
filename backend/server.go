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
	var studentData StudentInfo
	if err := c.Bind(&studentData); err == nil {
		if namePronunciation, err := c.FormFile("namePronunciation"); err == nil {
			studentData.NamePronunciationPath = handleUpload(namePronunciation, &studentData, namePronunciationPath)
		}
		if profilePicture, err := c.FormFile("profilePicture"); err == nil {
			studentData.ProfilePicturePath = handleUpload(profilePicture, &studentData, profilePicturePath)
		}
		studentData.AddEntry()
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
	var studentData StudentInfo
	if err := c.Bind(&studentData); err == nil {
		DeleteEntryByIDPreserveFiles(studentData.ObjectID)
		studentData.AddEntry()
		c.String(http.StatusOK, fmt.Sprintf("File %s", studentData.Name))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
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
