package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tylers Homepage")
	fmt.Println("Endpoint Hit: homePage")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	//adding the article route\
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequests()
}

//struct structure of data fields
//with declared data
type Article struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}

//declaring an articles array
//simulate a db
var Articles []Article

//CRUD
//retiving all the articles
func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit! Return articles")
	json.NewEncoder(w).Encode(Articles)
}
