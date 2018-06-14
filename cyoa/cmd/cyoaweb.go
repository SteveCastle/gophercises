package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SteveCastle/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA app on")
	file := flag.String("file", "gopher.json", "The json file")
	flag.Parse()
	fmt.Printf("using the story in %s\n", *file)
	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	fmt.Printf("%+v\n", story)
}
