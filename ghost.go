package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var file string
var private bool
var description string
var anonymous bool
var login bool

func init() {
	flag.StringVar(&file, "f", "", "set a filename for your gist")
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
	for _, file := range flag.Args() {
		fmt.Println(file)
	}

	if anonymous {
		fmt.Println("Posting anonymously")
	}

	if login {
		fmt.Println("Signing in")
	}

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	url := GITHUB_API + "/gists"
	fmt.Println("Sending data to: ", url)
   
    file, err := os.Open(flag.Arg(0))
    if err != nil {
            log.Fatal(err)
    }

	resp, err = http.Post(url, "application/json", file)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
