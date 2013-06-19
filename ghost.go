package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const GITHUB_API = "https://api.github.com/gists?access_token="

type Gist struct {
	Description string                       `json:"description"`
	Public      bool                         `json:"public"`
	Files       map[string]map[string]string `json:"files"`
}

func main() {

	var token string
	if token = os.Getenv("GITHUB_PAT"); token == "" {
		log.Fatal("The GITHUB_PAT environment variable must be set with your GitHub Personal API access token")
	}

	url := GITHUB_API + token

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // susceptible to man-in-the-middle
	}
	client := &http.Client{Transport: tr}

	flag.Parse()

	files := make(map[string]map[string]string)

	for _, name := range flag.Args() {

		contents, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		file := make(map[string]string)
		file["content"] = string(contents)
		files[name] = file
	}

	g := Gist{"", true, files}
	b, err := json.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(b)
	resp, err := client.Post(url, "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
