package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tylers HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

var Articles []Article

// Existing code from above
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true) //Http request multiplexer
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Multiplexer router")
	Articles = []Article{
		Article{Title: "Hello", Desc: "Dummy desc", Content: "Dummy Data"},
		Article{Title: "Hello 2", Desc: "Dummy desc", Content: "Dummy Data"},
	}
	handleRequests()
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

}
func createArticle(w http.ResponseWriter, r *http.Request) {
	//post request
	//string reponse containing the request body
	requestBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(requestBody))
}
