package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prams := mux.Vars(r)

	for index, item := range movies {
		if item.ID == prams["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for _, item := range movies {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(10000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Name: "ameiele", Director: &Director{FirstName: "Amr", LastName: "Effat"}})
	movies = append(movies, Movie{ID: "2", Name: "Matrix", Director: &Director{FirstName: "mohamed", LastName: "khan"}})

	r.HandleFunc("/Movies", getMovies).Methods("GET")
	r.HandleFunc("/Movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/Movies", createMovie).Methods("POST")
	r.HandleFunc("/Movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/Movies", deleteMovie).Methods("DELETE")

	fmt.Println("Starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"LastName"`
}
