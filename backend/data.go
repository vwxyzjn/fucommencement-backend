package backend

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-184-73-202-112.compute-1.amazonaws.com"
	port     = 5432
	user     = "bfhpwkwgaddttc"
	password = "2a0ed3c9f886553f54296475d08ca5f2bb2067449368d9df363706c3f4672a24"
	dbname   = "d3g2l997ob4u2o"
)

// Migrate migrates the database schema
func Migrate() {
	db := makeConnection()
	defer db.Close()
	_, err := db.Exec(`
		CREATE TABLE student (
			"name" TEXT,
			"furmanID" INT PRIMARY KEY NOT NULL,
			"anticipatedCompletionDate" TEXT,
			"degreeExpected" TEXT,
			"majors" TEXT,
			"interdisciplinaryMinor" TEXT,
			"diplomafirstName" TEXT,
			"diplomamiddleName" TEXT,
			"diplomalastName" TEXT,
			"hometownAndState" TEXT,
			"pronounceFirstName" TEXT,
			"pronounceMiddleName" TEXT,
			"pronounceLastName" TEXT,
			"rhymeFirstName" TEXT,
			"rhymeMiddleName" TEXT,
			"rhymeLastName" TEXT,
			"postGradAddress" TEXT,
			"postGradAddressTwo" TEXT,
			"postGradCity" TEXT,
			"postGradState" TEXT,
			"postGradPostalCode" TEXT,
			"postGradTelephone" TEXT,
			"postGradEmail" TEXT,
			"intentConfirm" TEXT,
			"namePronunciation" bytea,
			"profilePicture" bytea
		)
	`)
	CheckErr(err)
}

func test() {
	db := makeConnection()
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertId int
	err := db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	CheckErr(err)
	fmt.Println("last inserted id =", lastInsertId)

	fmt.Println("# Updating")
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	CheckErr(err)

	res, err := stmt.Exec("astaxieupdate", lastInsertId)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect, "rows changed")

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM userinfo")
	CheckErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}

	fmt.Println("# Deleting")
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	CheckErr(err)

	res, err = stmt.Exec(lastInsertId)
	CheckErr(err)

	affect, err = res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect, "rows changed")
}

func makeConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	CheckErr(err)

	return db
}

// CheckErr ..
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// SaveData ..
func SaveData(data []byte) error {
	db := makeConnection()
	defer db.Close()
	_, err := db.Exec("INSERT INTO k(file_name, blob, file_size) VALUES($1,$2,$3)", "test", data, 4)
	CheckErr(err)
	return err
}
