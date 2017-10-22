package main

import (
	"fmt"
	"net/http"
)

const version string = "1.0"

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Version %s of your application is running on Kubernetes.", version)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", version)
}

func main() {
	http.HandleFunc("/", getFrontpage)
	http.HandleFunc("/health", health)
	http.HandleFunc("/version", getVersion)
	http.ListenAndServe(":8080", nil)
}
