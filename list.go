package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
)

func fListCompany(w http.ResponseWriter, r *http.Request)  {
	slFile, err := ioutil.ReadDir("./csv")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		log.Println(err)
		return
	}

	var slFileName []string

	for _, oFile := range slFile {
		slFileName = append(slFileName, oFile.Name()[:len(oFile.Name()) - 4])
	}

	var jsonb []byte
	jsonb, err = json.Marshal(slFileName)
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		log.Println(err)
		return
	}

	fmt.Fprint(w, string(jsonb))
}
