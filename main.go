package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const(
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
)

func FileUploadHandler(w http.ResponseWriter,r *http.Request){
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("error getting a file from form:",err)
		return
	}
	//po emolchaniu fail otkrit
	defer file.Close()

	outFile, pathError := os.Create("upload/Upload_"+header.Filename)
	if pathError != nil {
		log.Println("error creating a file for writing:",pathError)
	}
	defer outFile.Close()

	_,copyFileError := io.Copy(outFile,file)
	if copyFileError != nil {
		log.Println("copyFile error:",pathError)
		return
	}
	fmt.Fprintf(w,"File uploaded successfully "+header.Filename)
}

func UploadPageFormHandler(w http.ResponseWriter,r *http.Request){
	parsedTemplate,_ := template.ParseFiles("templates/upload.html")
	err := parsedTemplate.Execute(w,nil)
	if err != nil {
		log.Println("error parsing upload template", err)
		return
	}
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/",UploadPageFormHandler).Methods("GET","POST")
	router.HandleFunc("/upload",FileUploadHandler).Methods("GET","POST")

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT,router)
	if err != nil {
		log.Fatal("error starting server:",err)
		return
	}
}
