package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleName).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeader).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["PARAM"]
	if param != "" {
		fmt.Fprintf(w, "Hello, %s!", param)
	} else {
		fmt.Fprintf(w, "Empty body")
	}
	w.WriteHeader(http.StatusOK)
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	d, _ := io.ReadAll(r.Body)
	if len(d) != 0 {
		fmt.Fprintf(w, "I got message:\n%s", string(d))
	} else {
		fmt.Fprintf(w, "No body set")
	}
	w.WriteHeader(http.StatusOK)
}

func handleHeader(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	val1, err1 := strconv.Atoi(strings.Join(h["A"], ""))
	val2, err2 := strconv.Atoi(strings.Join(h["B"], ""))
	if err1 == nil && err2 == nil {
		w.Header().Set("a+b", strconv.Itoa(val1+val2))
	} else {
		fmt.Fprintf(w, "No headers set")
	}
	w.WriteHeader(http.StatusOK)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
