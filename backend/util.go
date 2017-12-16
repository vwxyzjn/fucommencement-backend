package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

// HandleUpload reads a file from c.PostForm and return its stored path
func HandleUpload(file *multipart.FileHeader, student *StudentInfo, path string) string {
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

// Describe is a better println debugging method
func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// DeleteFile deletes file based on directory `path`
func DeleteFile(path string) {
	// check if file exist first.
	fmt.Println(path)
	if path == "." {
		return
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := os.Remove(path)
		if err != nil {
			panic(err)
		}
	}
}

// PrettyPrint prettily prints variable (struct, map, array, slice) in Golang. The easiest way is through MarshalIndent function in json package.
func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}
