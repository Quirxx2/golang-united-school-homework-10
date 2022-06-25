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

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/name/{PARAM}", handleName).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/header", handleHeader).Methods(http.MethodGet)
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
	if len(r.Header) != 0 {
		keys := ""
		values := 0
		for k, v := range r.Header {
			keys = keys + k + "+"
			tmpval, _ := strconv.Atoi(strings.Join(v, ""))
			values = values + tmpval
		}
		keys = keys[:len(keys)-1]
		fmt.Fprintf(w, "%s: %s", keys, strconv.Itoa(values))
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
