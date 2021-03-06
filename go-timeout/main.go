package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	log.Println("listening to port *:8080. Press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query().Encode()
	svcURL := os.Getenv("SERVICE_URL") + "/?" + qs
	resp, err := http.Get(svcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, `{"hello": "%s"}`, string(body))
}
