package main

import (
	"io/ioutil"
        "log"
	"fmt"
	"net/http"
)

const version string = "1.0"

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://kubeapp-canary-service/")
	if err != nil {
	// handle error
		log.Fatalln(err)
	}	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
        // handle error
		log.Fatalln(err)
        }

	fmt.Fprintf(w, "Congratulations! Version %s of your application is running on Kubernetes. \n from canary:\n %s", version, string(body))
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
