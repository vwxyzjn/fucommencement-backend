package backend

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentInfo struct {
	name                      string
	furmanID                  int
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
	intentConfirmOptions      bool
	intentConfirm             string
	namePronunciation         string
	profilePicture            string
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
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
}
