package Forum

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var FacebookLoginConfig = oauth2.Config{
	ClientID:     "335285802103443",
	ClientSecret: "0e4162ef6bac0c2bda1d2272cdabe825",
	Endpoint:     facebook.Endpoint,
	RedirectURL:  "http://localhost:8080/facebook/callback",
	Scopes: []string{
		"email",
		"public_profile",
	},
}

type UserFacebook struct {
	Id    int
	Name  string
	Email string
}

func FbLogin(rw http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	oauthState := generateStateOauthCookie(rw)

	u := FacebookLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(rw, r, u, http.StatusTemporaryRedirect)
}

func FbCallBack(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get oauth state from cookie for this user
	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")
	db := InitDatabase("ForumDB.db")
	defer db.Close()
	var User UserFacebook
	rw.Header().Add("content-type", "application/json")

	// ERROR : Invalid OAuth State
	if state != oauthState.Value {
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(rw, "invalid oauth facebook state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := FacebookLoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(rw, "falied code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from facebook server
	response, err := http.Get("https://graph.facebook.com/v13.0/me?fields=id,name,email,picture&access_token&access_token=" + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		fmt.Fprintf(rw, "failed getting user info: %s", err.Error())
		return
	}

	// Parse user data JSON Object
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	json.Unmarshal(contents, &User)
	if err != nil {
		fmt.Fprintf(rw, "failed read response: %s", err.Error())
		return
	}
	InsertIntoFacebookUsers(db, User.Name, User.Email)
	// send back response to browser
	fmt.Println(string(contents))
	fmt.Fprintln(rw, string(contents))
}
