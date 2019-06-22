package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles [2]Article

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	// var articles [2]Article
	Articles[0] = Article{Title: "Test1", Desc: "Article description1", Content: "article content"}
	Articles[1] = Article{Title: "Test2", Desc: "Article description1", Content: "article content"}

	handleRequests();
}

