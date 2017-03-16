package main

import (
	"net/http"
	"log"
	"net/url"
	"html/template"
)

var (
	left int
	right int
)

type Page struct {
	Left int
	Right int
}

func main() {
	mux := http.NewServeMux()
	//fs := http.FileServer(http.Dir("public"))
	log.Println("Listening on port :3000")

	mux.HandleFunc("/", home)
	mux.HandleFunc("/vote", vote)

	http.ListenAndServe(":3000", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	page := Page{Left: left, Right: right}
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(w, page)
}

func vote(w http.ResponseWriter, r *http.Request) {
	values, _ := url.ParseQuery(r.URL.RawQuery)
	hand := values.Get("hand")

	log.Println(hand)
	if hand == "left" {
		left++
	}
	if hand == "right" {
		right++
	}

	page := Page{Left: left, Right: right}
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(w, page)
}