package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Respond to an HTTP POST request with
		{"POST":true}
	If the request ot this endpoint is not a POST request,
	return an error status and "BAD METHOD"
*/
func getPOSTResp(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]bool)
	if r.Method == "POST" {
		response["POST"] = true
		respJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("[!] ERR creating JSON object in getPOSTResp")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	} else {
		http.Error(w, "BAD METHOD", http.StatusBadRequest)
		return
	}
}

/*
	Respond to an HTTP GET request with
		{"GET":true}
	If the request ot this endpoint is not a GET request,
	return an error status and "BAD METHOD"
*/
func getGETResp(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]bool)
	if r.Method == "GET" {
		response["GET"] = true
		respJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("[!] ERR creating JSON object in getGETResp")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	} else {
		http.Error(w, "BAD METHOD", http.StatusBadRequest)
		return
	}
}

/*
	Respond to an HTTP PATCH request with
		{"PATCH":true}
	If the request ot this endpoint is not a PATCH request,
	return an error status and "BAD METHOD"
*/
func getPATCHResp(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]bool)
	if r.Method == "PATCH" {
		response["PATCH"] = true
		respJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("[!] ERR creating JSON object in getPATCHResp")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	} else {
		http.Error(w, "BAD METHOD", http.StatusBadRequest)
		return
	}
}

/*
	Respond to an HTTP PUT request with
		{"PUT":true}
	If the request ot this endpoint is not a PUT request,
	return an error status and "BAD METHOD"
*/
func getPUTResp(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]bool)
	if r.Method == "PUT" {
		response["PUT"] = true
		respJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("[!] ERR creating JSON object in getPUTResp")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	} else {
		http.Error(w, "BAD METHOD", http.StatusBadRequest)
		return
	}
}

/*
	Respond to an HTTP DELETE request with
		{"DELETE":true}
	If the request ot this endpoint is not a DELETE request,
	return an error status and "BAD METHOD"
*/
func getDELETEResp(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]bool)
	if r.Method == "DELETE" {
		response["DELETE"] = true
		respJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("[!] ERR creating JSON object in getDELETEResp")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	} else {
		http.Error(w, "BAD METHOD", http.StatusBadRequest)
		return
	}
}
