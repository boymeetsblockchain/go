package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1234848", Title: "Movie One", Director: &Director{Firstname: "pablo", Lastname: "mari"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1234848", Title: "Movie Two", Director: &Director{Firstname: "pablo", Lastname: "mari the second"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
