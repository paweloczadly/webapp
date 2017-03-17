package main

import (
	"net/http"
	"log"
	"net/url"
	"html/template"
	"github.com/paweloczadly/webapp/utils"
)

var (
	left int
	right int
	influxEnabled bool = true
)

type Page struct {
	Left int
	Right int
}

func main() {
	utils.DumpAllEnvVars()

	mux := http.NewServeMux()
	log.Println("Listening on port :" + utils.AppPort())
	fs := http.FileServer(http.Dir(utils.Content()))

	mux.Handle("/", fs)
	mux.HandleFunc("/vote", vote)
	http.ListenAndServe(":" + utils.AppPort(), mux)
}

func vote(w http.ResponseWriter, r *http.Request) {
	values, _ := url.ParseQuery(r.URL.RawQuery)
	hand := values.Get("hand")

	log.Println(hand)
	if hand == "left" {
		left++
		if influxEnabled {
			utils.WriteToInflux(left, "left")
		}
	}
	if hand == "right" {
		right++
		if influxEnabled {
			utils.WriteToInflux(right, "right")
		}

	}

	page := Page{Left: left, Right: right}
	t, _ := template.ParseFiles(utils.Content() + "/index.html")
	t.Execute(w, page)
}



