package main

import (
	"fmt"
	"net/http"

	"github.com/neekonsu/chooseyourownadventure"
)

func main() {
	mux := defaultMux()
	newMap := chooseyourownadventure.MakeMap(JSON)
	chooseyourownadventure.PrintMap(newMap)
	mapHandler := chooseyourownadventure.MapHandler(newMap, mux)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusFound)
	})
	return mux
}
