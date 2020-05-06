package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GlLige struct {
	Lige []Lige `json:"lige"`
}
type Tipovi struct {
	Naziv string `json:"naziv"`
}
type Razrade struct {
	Tipovi []Tipovi `json:"tipovi"`
	Ponude []int    `json:"ponude"`
}
type Lige struct {
	Naziv   string    `json:"naziv"`
	Razrade []Razrade `json:"razrade"`
}
type GlPonude []struct {
	Broj          string     `json:"broj"`
	id            int        `json:"id"`
	Naziv         string     `json:"naziv"`
	Vrijeme       time.Time  `json:"vrijeme"`
	Tecajevi      []Tecajevi `json:"tecajevi"`
	TvKanal       string     `json:"tv_kanal,omitempty"`
	ImaStatistiku bool       `json:"ima_statistiku,omitempty"`
}
type Tecajevi struct {
	Tecaj float64 `json:"tecaj"`
	Naziv string  `json:"naziv"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	dataLige := new(GlLige)
	dataPonude := new(GlPonude) // structs in variabs

	getJSON("https://www.dropbox.com/s/2kqweiiqf6nbhfz/lige.json?dl=0", dataLige)
	getJSON("https://www.dropbox.com/s/wr9vnmt5e1jhwkq/ponude.json?dl=0", dataPonude) // getting data and populating structs

	fmt.Println(dataLige.Lige)
	fmt.Println(dataPonude.[1]id) // output fields/values of structs 
}	
	/*var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))*/


// lige.json = https://www.dropbox.com/s/2kqweiiqf6nbhfz/lige.json?dl=0
// ponude.json = https://www.dropbox.com/s/wr9vnmt5e1jhwkq/ponude.json?dl=0
