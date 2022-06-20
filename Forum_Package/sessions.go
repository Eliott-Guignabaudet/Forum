package Forum

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetSession(w http.ResponseWriter, r *http.Request) {
	var user UserParams
	session, _ := store.Get(r, "session")
	auth := session.Values["authenticated"]
	if auth != nil {
		json.Unmarshal([]byte(auth.(string)), &user)
		w.Write([]byte("{\"id\":" + strconv.Itoa(user.Id) + ", \"pseudo\":\"" + user.Pseudo + "\", \"email\":\"" + user.Email + "\"}"))
		//w.Write([]byte(auth.(string)))

	} else {
		w.Write([]byte("{\"resp\":\"not authenticated\"}"))
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "session",
		MaxAge: -1}
	http.SetCookie(w, &c)
	w.Write([]byte("{\"resp\":\"session deleted\"}"))
}
