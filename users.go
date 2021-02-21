package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func HomeUsersHandler(w http.ResponseWriter, r *http.Request) {
	endpoints := Endpoint{
		Title: "users",
		Endpoints: []string{"/", "/:id"},
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(endpoints)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	metaData, err := GetMetaData(ioutil.Discard, "aimi.kamitoon.ohori.me", "users/" + vars["id"] + "/" + vars["hash"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print(err)
		fmt.Fprintf(w, "An error occured:\n%v", err)
		return
	}

	data, err := DownloadFile(ioutil.Discard, "aimi.kamitoon.ohori.me", "users/test/test.jpg")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occured")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", metaData.ContentType)

	w.Write(data)
}

func UserTestHandler(w http.ResponseWriter, r *http.Request) {
	metaData, err := GetMetaData(ioutil.Discard, "aimi.kamitoon.ohori.me", "users/test/test.jpg")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print(err)
		fmt.Fprintf(w, "An error occured:\n%v", err)
		return
	}

	data, err := DownloadFile(ioutil.Discard, "aimi.kamitoon.ohori.me", "users/test/test.jpg")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occured")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", metaData.ContentType)

	w.Write(data)
}