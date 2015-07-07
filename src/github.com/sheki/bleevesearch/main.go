package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var m map[string]interface{}
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Println(m)
	io.WriteString(w, `{ "success" : true}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", mux)
}
