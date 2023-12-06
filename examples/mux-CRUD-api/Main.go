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

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Status struct {
	Error   string `json:"error,omitempty"`
	Success string `json:"success,omitempty"`
}

const (
	PORT int = 8000
)

var movies []Movie

func databaseInit() {
	/** Adding some movies to the memory database */
	movies = append(movies, Movie{
		ID:    "2516163492471",
		ISBN:  "438227",
		Title: "Avatar",
		Director: &Director{
			Firstname: "James",
			Lastname:  "Cameron",
		},
	})

	movies = append(movies, Movie{
		ID:    "834338520728",
		ISBN:  "438228",
		Title: "Jurassic Park",
		Director: &Director{
			Firstname: "Steven",
			Lastname:  "Spielberg",
		},
	})
}

func main() {
	/** Adds the initial mock database entries */
	databaseInit()

	/** Create a Router */
	r := mux.NewRouter()

	/** Routes */
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")
	http.Handle("/", r)

	/** Start server */
	fmt.Printf("Starting server at http://localhost:%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprint(":", PORT), nil); err != nil {
		log.Fatal(err)
	}
}

/***************************/
/** Handler Functions */
/***************************/

/** Get All Movies */
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

/** Get a Movie by ID */
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID != params["ID"] {
			continue
		}
		json.NewEncoder(w).Encode(movie)
		return
	}
	msg, _ := json.Marshal(&Status{Error: "Movie not Found"})
	http.Error(w, string(msg), http.StatusNotFound)

}

/** Create a Movie */
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

/** Update a Movie */
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	params := mux.Vars(r)
	for i, movie := range movies {
		if movie.ID != params["ID"] {
			continue
		}
		/** Update the ID with the body */
		var item Movie
		_ = json.NewDecoder(r.Body).Decode(&item)
		item.ID = movie.ID

		/** Update Movie */
		movies = append(append(movies[:i], item), movies[i+1:]...)
		json.NewEncoder(w).Encode(item)
		return
	}
	msg, _ := json.Marshal(&Status{Error: "Movie not Found"})
	http.Error(w, string(msg), http.StatusNotFound)
}

/** Delete a Movie */
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, movie := range movies {
		if movie.ID != params["ID"] {
			continue
		}
		/** Leave out the ID that is matched */
		movies = append(movies[:i], movies[i+1:]...)
		json.NewEncoder(w).Encode(movies)
		return
	}
	msg, _ := json.Marshal(&Status{Error: "Movie not Found"})
	http.Error(w, string(msg), http.StatusNotFound)
}
