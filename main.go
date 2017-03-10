package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	dir := flag.String("d", "./", "directory with mock json")
	flag.Parse()

	log.Printf("Starting api-mock from dir %s\n", *dir)

	walker := NewWalker(*dir)
	filepath.Walk(*dir, walker.WalkDirs)

	log.Println("Creating handler for directory structure")
	handler := NewHandler(walker)

	router := mux.NewRouter()
	router.HandleFunc("/{path:.*}", handler.get).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
