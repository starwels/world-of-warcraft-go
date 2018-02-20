package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Npc struct
type Npc struct {
	ID                int
	Name              string
	URLSlug           string
	CreatureDisplayID int
}

// Boss struct
type Boss struct {
	ID                    int
	Name                  string
	URLSlug               string
	Description           string
	ZondeID               int
	AvailableInNormalMode bool
	AvailableInHeroicMode bool
	Health                int
	Level                 int
	HeroicLevel           int
	JournalID             int
	Npcs                  []Npc
}

const locale string = "locale=en_US"
const url string = "https://us.api.battle.net/wow/boss/"

func main() {
	b := Boss{}
	apikey, err := ioutil.ReadFile("apikey.txt")
	if err != nil {
		panic(err)
	}

	bossID := flag.String("boss", "24723", "boss id")

	flag.Parse()

	key := strings.TrimSpace(string(apikey))

	fullURL := url + *bossID + "?" + locale + "&" + key

	resp, err := http.Get(fullURL)

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(resp.Body).Decode(&b)

	defer resp.Body.Close()
	fmt.Println(b.Name)
	fmt.Println(b.Description)
}
