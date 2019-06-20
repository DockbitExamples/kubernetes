package main

import (
  "io/ioutil"
  "log"
  "net/http"
  "io"
  "encoding/json"
  "fmt"
  "strconv"
  "math/rand"
)

const version string = "2.0"

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func dumpJsonRequestHandlerFunc(w http.ResponseWriter, req *http.Request){
  //Validate request
  if req.Method != "POST" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if req.Header.Get("Content-Type") != "application/json" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  //To allocate slice for request body
  length, err := strconv.Atoi(req.Header.Get("Content-Length"))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  //Read body data to parse json
  body := make([]byte, length)
  length, err = req.Body.Read(body)
  if err != nil && err != io.EOF {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  //parse json
  var jsonBody map[string]interface{}
  err = json.Unmarshal(body[:length], &jsonBody)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  // myString := jsonBody["message"].(string)
  jsonBody["message"] = reverse(jsonBody["message"].(string))
  jsonBody["rand"] = rand.Float64()

  jsonString, err := json.Marshal(jsonBody)
  fmt.Println(err)
  fmt.Fprintf(w, "%s\n", jsonString)

  w.WriteHeader(http.StatusOK)  
}

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://kubeapp-production-service/")
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

	fmt.Fprintf(w, "Congratulations! Version %s of your application is running on Kubernetes. \n from another:\n %s", version, string(body))
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
	http.HandleFunc("/api", dumpJsonRequestHandlerFunc)
	http.ListenAndServe(":8080", nil)
}
