package backend

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ProfilePicturePath    string
	NamePronunciationPath string
	Algolia               AlgoliaInstance
}

// Setup setups the server http end points
func (s *Server) Setup(AlgoliaAppID string, AlgoliaKey string, AlgoliaIndexName string) {
	// Setup Algolia
	s.Algolia.Initialize(AlgoliaAppID, AlgoliaKey, AlgoliaIndexName)

	// Setup Router
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", testGET)
	r.GET("/deleteEntryGET/:objectID", deleteEntryGET)
	r.POST("/commencementPOST", s.commencementPOST)
	r.POST("/updateEntryPOST", updateEntryPOST)
	r.StaticFS("/commencement", http.Dir("./commencement"))
	r.Run() // listen and serve on 0.0.0.0:8080
}

func testGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) commencementPOST(c *gin.Context) {
	var studentData StudentInfo
	if err := c.Bind(&studentData); err == nil {
		if namePronunciation, err := c.FormFile("namePronunciation"); err == nil {
			studentData.NamePronunciationPath = HandleUpload(namePronunciation, &studentData, s.NamePronunciationPath)
		}
		if profilePicture, err := c.FormFile("profilePicture"); err == nil {
			studentData.ProfilePicturePath = HandleUpload(profilePicture, &studentData, s.ProfilePicturePath)
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
