package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	/*
	cobaHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini Adalah Handler Coba")) // function expression
	}
	*/

	mux.HandleFunc("/", handler.HandlerIndex) //membuat routing index
	mux.HandleFunc("/about", handler.AboutPage)
	mux.HandleFunc("/profile", handler.ProfilePage)
	mux.HandleFunc("/product", handler.ProductHandler)
	// mux.HandleFunc("/coba", cobaHandler)

	/*
	mux.HandleFunc("/anonim", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini Adalah ANONYMOUSE FUNCTION")) //anonymouse function
	})
	*/

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) //link u. assets a. bs digunakan di layout.html

	log.Println("Starting port :8888") //memberikan pesan di terminal

	err := http.ListenAndServe(":8888", mux) //perintah untuk menjalankan web server
	log.Fatal(err)
}

