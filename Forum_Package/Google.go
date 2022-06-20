package Forum

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserGoogle struct {
	Id             int
	Email          string
	Verified_email string
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/google/callback",
	ClientID:     "76456776386-97d624l8e74c6bj95tacm4udrs0veiad.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func GoogleLogin(rw http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(rw)

	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(rw, r, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func GoogleCallBack(rw http.ResponseWriter, r *http.Request) {
	// check is method is correct
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
	var User UserGoogle
	rw.Header().Add("content-type", "application/json")

	// ERROR : Invalid OAuth State
	if state != oauthState.Value {
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(rw, "invalid oauth google state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := googleOauthConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(rw, "falied code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from google server
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

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
	fmt.Println(User)
	InsertIntoGoogleUsers(db, User.Name, User.Email)
	// send back response to browser
	fmt.Fprintln(rw, string(contents))
}
