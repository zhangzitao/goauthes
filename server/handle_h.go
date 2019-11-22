package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/zhangzitao/goauthes/factory"
)

func errorToMap(message string) (m map[string]interface{}) {
	m = map[string]interface{}{"Type": "error", "Message": message}
	return m
}

func handlerToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	class := r.FormValue("grant_type")
	if class == "" {
		json.NewEncoder(w).Encode(errorToMap("grant type needed"))
		return
	}
	switch class {
	case "password":
		processPassword(w, r)
	case "refresh_token":
		processRefresh(w, r)
	default:
		json.NewEncoder(w).Encode(errorToMap("grant type not support"))
	}
}

func processPassword(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Println("error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()
	username := r.FormValue("username")
	password := r.FormValue("password")
	scope := r.FormValue("scope")
	if username == "" || password == "" {
		json.NewEncoder(w).Encode(errorToMap("password and username needed"))
		return
	}
	// input file to cred
	cred, err := factory.GenerateCredential("PassWord", "PassWord", username, password, scope)
	if err != nil {
		return
	}
	// cred to auth
	auth, err := factory.GenerateAuthorizeFromCredential(cred)
	if err != nil {
		return
	}
	// after verify, auth has userid
	authDone, err := auth.Verify(os.Getenv("GOAUTHES_AUTHORIZE_VERIFY_MODE"))
	if err != nil {
		log.Println("error:", err)
		json.NewEncoder(w).Encode(errorToMap(err.Error()))
		return
	}
	// if password verified false, return
	if !authDone.Verified() {
		json.NewEncoder(w).Encode(errorToMap("Password verify failed"))
		return
	}
	// create storage interface
	sto, err := factory.GenerateFromAuthorize("PassWord", authDone)
	if err != nil {
		return
	}
	// storage save to db
	flag, err := sto.Create()
	if err != nil {
		return
	}
	if !flag {
		return
	}
	tok, err := factory.GenerateTokenFromStorage(sto)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(tok)
}

func processRefresh(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Println("error:", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()
	refreshToken := r.FormValue("refresh_token")
	if refreshToken == "" {
		json.NewEncoder(w).Encode(errorToMap("refresh_token needed"))
		return
	}

	stoOld, err := factory.GetRefreshStorage(refreshToken)
	if err != nil {
		json.NewEncoder(w).Encode(errorToMap("refresh_token wrong"))
		return
	}
	// delete all tokens of this storage
	_, err = stoOld.Refresh()
	if err != nil {
		return
	}
	// rebuild storage
	sto, err := factory.MakeNewStorageByOld(stoOld)
	if err != nil {
		return
	}
	// storage save to db
	flag, err := sto.Create()
	if err != nil {
		return
	}
	if !flag {
		return
	}
	tok, err := factory.GenerateTokenFromStorage(sto)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(tok)
}
