package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/company/list", fListCompany)
	http.HandleFunc("/company/view", fViewCompany)

	strHost := "localhost"
	strPort := "58099"

	log.Printf("starting at '%s:%s' \n", strHost, strPort)

	err := http.ListenAndServe(strHost + ":" + strPort, http.DefaultServeMux)

	if err != nil {
		log.Fatal(err)
	}
}
