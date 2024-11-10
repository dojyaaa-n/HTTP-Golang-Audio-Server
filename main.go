package main

import (
	"log"
	"io"
	"os"
	"net/http"
)

var fname = "audio.aac"

func main() {
	http.HandleFunc("/stream", stream)

	log.Println("Listening on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func stream(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Transfer-Encoding", "chunked")

	log.Printf("%s has connected to the audio stream\n", r.Host)

	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatal(err)
	}
}