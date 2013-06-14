package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"path"
)

var private bool
var description string
var anonymous bool
var login bool

func init() {
	flag.BoolVar(&private, "p", false, "make the gist private")
	flag.StringVar(&description, "d", "", "the gist of the gist!")
	flag.BoolVar(&anonymous, "a", false, "anonymously post a gist, even while signed in")
	flag.BoolVar(&login, "login", false, "sign in to Github.com or your an instance of Github Enterprise")
}

const (
	GITHUB_API = "https://api.github.com"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // susceptible to man-in-the-middle
	}
	client := &http.Client{Transport: tr}

	url := GITHUB_API + "/gists"

	flag.Parse()

	if anonymous {
		fmt.Println("Posting anonymously")
	}

	if login {
		fmt.Println("Signing in")
		usr, err := user.Current()
		file := path.Join(usr.HomeDir, ".gist")
		auth := url + "/authorizations"

		dump := `{
            "scopes": [
                "gist"
                ],
            "note": "yet another cli gister"
            }`

		payload := bytes.NewBufferString(dump)
		resp, err := client.Post(auth, "application/json", payload)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		//TODO: get token from resp.Body

		err := ioutil.WriteFile(file, []bytes(token), 0777)
		if err != nil {
			log.Fatal(err)
		}
		url = url + "?access_token=" + token
	}

	file, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(file)
	resp, err := client.Post(url, "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
