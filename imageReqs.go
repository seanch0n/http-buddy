package main

import (
	"io/ioutil"
	"net/http"
)

/*
	Read a JPG image (images/puppy.jpg) and send it back
	with the proper content type
*/
func getJPGImage(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("images/puppy.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "image/jpg")
	w.Write(file)
}

/*
	Read a PNG image (images/puppy.png) and send it back
	with the proper content type
*/
func getPNGImage(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("images/puppy.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "image/png")
	w.Write(file)
}
