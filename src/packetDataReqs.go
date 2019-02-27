package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
	getMyIP returns the IP address and port of the connected client
	in the format:
		{"ip":"12.7.0.0.1", "port":"12345"}
*/
func getMyIP(w http.ResponseWriter, r *http.Request) {
	response := IPResp{}
	tmp := strings.Split(r.RemoteAddr, ":")
	if len(tmp) == 2 {
		response.IP = tmp[0]
		response.Port = tmp[1]
	} else if len(tmp) == 1 {
		response.IP = tmp[0]
	} else {
		fmt.Println("[!] ERR parsing IP getMyIP")
		http.Error(w, "", http.StatusNotFound)
		return
	}
	respJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("[!] ERR creating JSON object in getPATCHResp")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

/*

	getHeaders returns a list of all the headers sent by the client.
	An example response is shown below
		{"Headers":[{"Accept":"text/html,application/xhtml+xml,application/xml"},
		{"Accept-Language":"en-US,en;q=0.5"}],"count":2}
*/
func getHeaders(w http.ResponseWriter, r *http.Request) {
	response := HeaderResp{}

	// loop over all the headers and put them in the map @tmp
	// so it can be turned into a JSON object
	// header format is:
	//		Accept-Language en-Us,en;q=0.5
	// 		User-Agent mySuper Cool UserAgent
	// Then we increment a counter so we can send the user the total
	// number of headers in the request
	for name, headers := range r.Header {
		for _, h := range headers {
			tmp := make(map[string]string)
			tmp[name] = h
			response.Headers = append(response.Headers, tmp)
			response.Count = response.Count + 1
		}
	}

	respJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("[!] ERR creating JSON object in getPATCHResp")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

/*
	getUserAgent parses the user agent out of the clients request
	and returns it in the format
		{"User-Agent":"mySuperCoolUserAgent"}
*/
func getUserAgent(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)

	// the user agent is sent over in a header so loop over
	// them till we find it, then break out of the loop
	for name, headers := range r.Header {
		for _, h := range headers {
			if strings.Compare(name, "User-Agent") == 0 {
				response[name] = h
				break
			}
		}
	}
	respJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("[!] ERR creating JSON object in getPATCHResp")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}
