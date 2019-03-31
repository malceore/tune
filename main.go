package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
)

// The new router function creates the router and
func newRouter() *mux.Router {
	r := mux.NewRouter()
        r.HandleFunc("/", ClientHandler)
        r.HandleFunc("/ws", WebsocketHandler)
	return r
}

func ClientHandler(w http.ResponseWriter, r *http.Request) {
     // If it does then we need to return the contents of it's file sin a javascript array, injected into the index.html page.
    bytes, err := ioutil.ReadFile("template.html")
    if err != nil {
	fmt.Println("Could not open templates.html :: ", err)
    }
    var templatebody = string(bytes)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, templatebody)
}

func main() {
	r := newRouter()
        fmt.Println("Tune server started on: 3434")
	http.ListenAndServe(":3434", r)
}

