package Forum

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetSession(w http.ResponseWriter, r *http.Request) {
	var user UserParams
	session, _ := store.Get(r, "cookie-name")
	auth := session.Values["authenticated"]
	if auth != nil {
		json.Unmarshal([]byte(auth.(string)), &user)
		w.Write([]byte("\"id\":\"" + strconv.Itoa(user.Id) + "\", \"pseudo\":\"" + user.Pseudo + "\", \"email\":\"" + user.Email + "\""))

	} else {
		w.Write([]byte("\"resp\":\"not authenticated\""))
	}
}
