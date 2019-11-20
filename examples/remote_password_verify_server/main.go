package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	err = r.ParseForm()
	if err != nil {
		log.Printf("parse form error:%v", err)
	}
	fmt.Println(r.Method)
	fmt.Printf("%v\n", r.PostForm)
	username := r.FormValue("username")
	password := r.FormValue("password")
	scope := r.FormValue("scope")
	if username == "" || password == "" {
		err = json.NewEncoder(w).Encode(map[string]interface{}{"Success": false, "Message": "need username and password"})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		return
	}
	if username == "123" && password == "321" && scope == "read" {
		err = json.NewEncoder(w).Encode(map[string]interface{}{"Success": true, "Message": "this_id_cool"})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{"Success": false, "Message": "username and password wrong"})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/user", myHandle)
	log.Fatal(http.ListenAndServe(":12321", nil))
}
