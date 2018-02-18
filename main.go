package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"flag"
	"strings"
)

const locale string = "locale=en_US"
const url string = "https://us.api.battle.net/wow/boss/"

func main() {

	apikey, err := ioutil.ReadFile("apikey.txt")
	if err != nil {
		panic(err)
	}

	bossId := flag.String("boss", "10184", "boss id")
	
	flag.Parse()

	key := strings.TrimSpace(string(apikey))

	resp, err := http.Get(url + *bossId + "?" + locale + "&" + key)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}