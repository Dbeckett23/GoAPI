package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}


var Articles [2]Article

func endPointHit(endPointName string){
	fmt.Println("Endpoint Hit: " + endPointName)
}
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: sayHello")
	fmt.Fprintf(w, "Hello!")
}

func currentTime(w http.ResponseWriter, r *http.Request){
	endPointHit("currentTime")
	fmt.Fprint(w, "The current time is: " + time.Now().String())
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/sayHello", sayHello)
	http.HandleFunc("/currentTime", currentTime)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	// var articles [2]Article
	Articles[0] = Article{Title: "Test1", Desc: "Article description1", Content: "article content"}
	Articles[1] = Article{Title: "Test2", Desc: "Article description1", Content: "article content"}

	handleRequests();
}

