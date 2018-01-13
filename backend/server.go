package backend

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server is ...
type Server struct {
	Port                   string
	ProfilePicturePath     string
	NamePronunciationPath  string
	AlgoliaAppID           string
	AlgoliaKey             string
	AlgoliaIndexName       string
	AlgoliaSortedIndexName string
	AlgoliaClient          algoliasearch.Client
	AlgoliaIndex           algoliasearch.Index
	AlgoliaSortedIndex     algoliasearch.Index
}

// Setup setups the server http end points
func (s *Server) Setup() {
	s.AlgoliaClient = algoliasearch.NewClient(s.AlgoliaAppID, s.AlgoliaKey)
	s.AlgoliaIndex = s.AlgoliaClient.InitIndex(s.AlgoliaIndexName)
	s.AlgoliaSortedIndex = s.AlgoliaClient.InitIndex(s.AlgoliaSortedIndexName)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", testGET)
	r.GET("/deleteEntryGET/:objectID", s.deleteEntryGET)
	r.GET("/entryByFurmanIDGET/:furmanID", s.entryByFurmanIDGET)
	r.GET("/entryByRankGET/:rank", s.entryByRankGET)
	r.POST("/commencementPOST", s.commencementPOST)
	r.POST("/updateEntryPOST", s.updateEntryPOST)
	r.StaticFS("/commencement", http.Dir("./commencement"))
	r.Run(s.Port) // listen and serve on 0.0.0.0:8080
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
		s.AddEntry(&studentData)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// temp, _ := strconv.Atoi(c.PostForm("furmanID"))
	// studentData.FurmanID = int(temp)
}

func (s *Server) deleteEntryGET(c *gin.Context) {
	objectID := c.Param("objectID")
	fmt.Println("deleted", objectID)
	s.DeleteEntryByID(objectID)
	c.String(http.StatusOK, "ok")
}

func (s *Server) updateEntryPOST(c *gin.Context) {
	var studentData StudentInfo
	if err := c.Bind(&studentData); err == nil {
		s.DeleteEntryByIDPreserveFiles(studentData.ObjectID)
		s.AddEntry(&studentData)
		c.String(http.StatusOK, fmt.Sprintf("File %s", studentData.Name))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (s *Server) entryByFurmanIDGET(c *gin.Context) {
	furmanID := c.Param("furmanID")
	temp, err := strconv.Atoi(furmanID)
	if err == nil {
		studentData := s.getEntryByFurmanID(temp)
		c.JSON(http.StatusOK, studentData)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (s *Server) entryByRankGET(c *gin.Context) {
	rank := c.Param("rank")
	temp, err := strconv.Atoi(rank)
	if err == nil {
		studentData := s.GetNthEntryInIndex(temp)
		c.JSON(http.StatusOK, studentData)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
