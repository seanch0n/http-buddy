package main

import (
	"fmt"
	"log"
	"net/http"
)

//HeaderResp is the format for /headers response
type HeaderResp struct {
	Headers []map[string]string `json:"Headers"`
	Count   int                 `json:"count"`
}

//IPResp is the format for /myip response
type IPResp struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}

func main() {
	fmt.Println("Serving on 0.0.0.0:8080")

	// serve a static page at 'http://domain.com/'
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/myip", getMyIP)
	http.HandleFunc("/headers", getHeaders)
	http.HandleFunc("/useragent", getUserAgent)
	http.HandleFunc("/images/jpg", getJPGImage)
	http.HandleFunc("/images/png", getPNGImage)
	http.HandleFunc("/methods/post", getPOSTResp)
	http.HandleFunc("/methods/get", getGETResp)
	http.HandleFunc("/methods/patch", getPATCHResp)
	http.HandleFunc("/methods/put", getPUTResp)
	http.HandleFunc("/methods/delete", getDELETEResp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
