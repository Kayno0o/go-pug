package main

import (
	"log"
	"net/http"

	_ "net"

	_ "github.com/Joker/hpp"
	_ "github.com/valyala/bytebufferpool"

	"kaynooo.xyz/go-pug/src/jade"
)

func main() {
	// serve static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// pages
	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		jade.Index(map[string]string{
			"title": "Test title",
			"jade":  "true",
			"name":  req.URL.Query().Get("q"),
		}, wr)
	})

	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
