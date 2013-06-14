package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	flag.Parse()

	if anonymous {
		fmt.Println("Posting anonymously")
	}

	if login {
		fmt.Println("Signing in")
	}

	url := GITHUB_API + "/gists"

	file, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // susceptible to man-in-the-middle
	}
	client := &http.Client{Transport: tr}

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
