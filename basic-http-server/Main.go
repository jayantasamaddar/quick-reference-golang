package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	PORT int = 8080
)

type Status struct {
	Status int
}

/** Documentation: https://pkg.go.dev/encoding/json#Marshal */
type FormData struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func main() {
	// Simple static webserver:
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	/** Hello Route */
	http.HandleFunc("/hello", helloHandler)

	/** Form Submission Route */
	http.HandleFunc("/signup-submit", formHandler)

	/** Start server */
	fmt.Printf("Starting server at http://localhost:%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprint(":", PORT), nil); err != nil {
		log.Fatal(err)
	}
}

/***************************/
/** Handler Functions */
/***************************/

/** Hello Handler */
func helloHandler(w http.ResponseWriter, r *http.Request) {
	/** Error boundary */
	if r.URL.Path != "/hello" {
		b, err := json.Marshal(&Status{Status: 404})
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Error(w, string(b), http.StatusNotFound)
		return
	}

	/** Validations */
	if r.Method != "GET" {
		jsonResp, err := json.Marshal(&Status{Status: 404})
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Error(w, string(jsonResp), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello")
}

/** Form Handler */
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v\n", err)
		return
	}

	/** Validations */
	if r.Method != "POST" {
		jsonResp, err := json.Marshal(&Status{Status: 404})
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Error(w, string(jsonResp), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Form submitted!\n")

	/** Method 1: Send a JSON Response */
	fmt.Fprintln(w, "-----------------------")
	fmt.Fprintln(w, "JSON Response: Method 1")
	fmt.Fprintln(w, "-----------------------")
	w.Header().Set("Content-Type", "application/json")
	user := &FormData{
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
	}
	json.NewEncoder(w).Encode(user)

	/** Method 2: Send a JSON Response */
	fmt.Fprintln(w, "-----------------------")
	fmt.Fprintln(w, "JSON Response: Method 2")
	fmt.Fprintln(w, "-----------------------")
	jsonResp, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonResp))
}
