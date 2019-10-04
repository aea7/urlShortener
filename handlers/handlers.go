package handlers

import (
	"../models"
	"github.com/gorilla/mux"
	"net/http"
)

// GET Endpoint that takes url parameter and returns shortened URL

func ShortenURL(w http.ResponseWriter, r *http.Request) {

	url, exists := r.URL.Query()["url"]

	if !exists || len(url[0]) < 1 {
		http.Error(w, "You must enter url parameter", http.StatusInternalServerError)
		return
	}

	hostname := r.Host
	key, err := models.CreateURLShortener(url[0])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	urlString := hostname + "/" + key
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(urlString))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GET Endpoint that takes shortened URL and redirects to original URL

func RedirectURL(w http.ResponseWriter, r *http.Request) {

	key, exists := mux.Vars(r)["key"]
	if !exists || len(key) != 6 {
		http.Error(w, "You must enter 6 letter key", http.StatusBadRequest)
		return
	}

	url, err := models.GetUrlFromDB(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)

}
