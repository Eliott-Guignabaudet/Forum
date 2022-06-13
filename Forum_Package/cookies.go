package Forum

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello World! "
	w.Write([]byte(html))
}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("ithinkidroppedacookie")
	if err != nil {
		w.Write([]byte("error in reading cookie : " + err.Error() + "\n"))
	} else {
		value := c.Value
		w.Write([]byte("cookie has : " + value + "\n"))
	}
}

// see https://golang.org/pkg/net/http/#Cookie
// Setting MaxAge<0 means delete cookie now.

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "login",
		MaxAge: -1}
	http.SetCookie(w, &c)

	w.Write([]byte("old cookie deleted!\n"))
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	var user UserParams
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		println("ERREUR : ", err)
	}

	c := http.Cookie{
		Name:   "login",
		Value:  `{"pseudo":` + user.Pseudo + ` ,"email":"` + user.Email + `"}`,
		MaxAge: 3600}
	http.SetCookie(w, &c)

	w.Write([]byte("new cookie created!\n"))
}
