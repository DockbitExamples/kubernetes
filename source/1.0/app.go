package main

import (
	"fmt"
	"net/http"
)

const version string = "1.0"

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

  jsonString, err := json.Marshal(jsonBody)
  fmt.Println(err)
  fmt.Fprintf(w, "%s\n", jsonString)

  w.WriteHeader(http.StatusOK)  
}

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
	http.HandleFunc("/reverse", dumpJsonRequestHandlerFunc)
	http.ListenAndServe(":8080", nil)
}
