package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"path"
)

const (
	GITHUB_API = "https://api.github.com"
)

type Gist struct {
	Description string                       `json:"description"`
	Public      bool                         `json:"public"`
	Files       map[string]map[string]string `json:"files"`
}

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // susceptible to man-in-the-middle
	}
	client := &http.Client{Transport: tr}

	url := GITHUB_API + "/gists"

	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	file := path.Join(usr.HomeDir, ".gist")

	token, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	url = url + "?access_token=" + string(token)

	contents, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	m["content"] = string(contents)
	n := make(map[string]map[string]string)
	n[flag.Arg(0)] = m

	g := Gist{"a file", true, n}
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
