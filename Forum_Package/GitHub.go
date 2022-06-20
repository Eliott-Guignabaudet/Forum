package Forum

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "05ec8d8cf8286d26c1d5",
		ClientSecret: "30c37a90aead69d383b81f662b627452dce6cf30",
		// select level of access you want https://developer.github.com/v3/oauth/#scopes
		Scopes:   []string{"user:email", "repo"},
		Endpoint: githuboauth.Endpoint,
	}
	oauthStateString = "thisshouldberandom"
)

func GitHubLogin(rw http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(rw)

	u := oauthConf.AuthCodeURL(oauthState)
	http.Redirect(rw, r, u, http.StatusTemporaryRedirect)
}

func GitHubCallBack(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get oauth state from cookie for this user
	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")

	rw.Header().Add("content-type", "application/json")

	// ERROR : Invalid OAuth State
	if state != oauthState.Value {
		http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(rw, "invalid oauth github state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := oauthConf.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(rw, "falied code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from facebook server
	response, err := http.Get("https://github.com/login/oauth/access_token" + token.TokenType)

	// ERROR : Unable to get user data from google
	if err != nil {
		fmt.Fprintf(rw, "failed getting user info: %s", err.Error())
		return
	}

	// Parse user data JSON Object
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(rw, "failed read response: %s", err.Error())
		return
	}

	// send back response to browser
	fmt.Println(response)
	fmt.Println(string(contents))
	fmt.Fprintln(rw, string(contents))
}
