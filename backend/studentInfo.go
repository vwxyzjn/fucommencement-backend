package backend

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "ec2-184-73-202-112.compute-1.amazonaws.com"
	port     = 5432
	user     = "bfhpwkwgaddttc"
	password = "2a0ed3c9f886553f54296475d08ca5f2bb2067449368d9df363706c3f4672a24"
	dbname   = "d3g2l997ob4u2o"
)

type StudentInfo struct {
	gorm.Model
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
}

// Migrate ..
func Migrate() {
	db := connect()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&StudentInfo{})

	// Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	// db.Delete(&product)
}

// CreateEntry ..
func (s *StudentInfo) CreateEntry() {
	db := connect()
	defer db.Close()

	db.Create(&s)

}

func connect() *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, user, dbname, "require", password))
	CheckErr(err)
	return db
}

// CheckErr ..
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
