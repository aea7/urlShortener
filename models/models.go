package models

import (
	"../config"
	"../middlewares"
	"errors"
	"strings"
)

type UrlShorten struct {
	Id  int
	Key string
	Url string
}

// Model package is for dealing with DB queries

// Following function Creates unique key for Shortened URL

func CreateURLShortener(url string) (string, error) {
	if url == "" {
		return "", errors.New("400 Bad Request. URL parameter cannot be empty")
	}

	err, key := GetNewKey(6)

	if err != nil {
		return "", err
	}

	_, err = config.DB.Exec("INSERT INTO urls VALUES (DEFAULT, $1, $2)", key, url)
	if err != nil {
		return "", err
	}

	return key, nil
}

// Checks if we have a duplicate key in the DB, if yes, returns another key

func GetNewKey(length int) (error, string) {
	key := middlewares.CreateKey(length)

	var exists bool
	_ = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE key = $1);", key).Scan(&exists)
	for exists == true {
		key = middlewares.CreateKey(length)
		_ = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE key = $1);", key).Scan(&exists)
	}

	return nil, key
}

// Fetches the original URL from the DB

func GetUrlFromDB(key string) (string, error) {
	urlShorten := UrlShorten{}
	if key == "" {
		return "", errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM urls WHERE key = $1", key)
	err := row.Scan(&urlShorten.Id, &urlShorten.Key, &urlShorten.Url)

	if err != nil {
		if urlShorten.Id != 0 && urlShorten.Url == "" {
			return urlShorten.Url, nil
		}
		return "", err
	}

	if strings.Contains(urlShorten.Url, "http://") {
		return urlShorten.Url, nil
	}else{
		finalUrl := "http://" + urlShorten.Url
		return finalUrl, nil
	}
}
