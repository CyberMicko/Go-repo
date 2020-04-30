package main

import (
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
	ID            int        `json:"id"`
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

func main() {
	dataLige, err := http.Get("https://www.dropbox.com/s/2kqweiiqf6nbhfz/lige.json?dl=0")
	dataPonude, err2 := http.Get("https://www.dropbox.com/s/wr9vnmt5e1jhwkq/ponude.json?dl=0")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	if err2 != nil {
		fmt.Println("File reading error", err2)
		return
	}
	fmt.Println("Contents of dataLige:", string(dataLige))
	fmt.Println("Contents of dataPonude:", string(dataPonude))
}

// lige.json = https://www.dropbox.com/s/2kqweiiqf6nbhfz/lige.json?dl=0
// ponude.json = https://www.dropbox.com/s/wr9vnmt5e1jhwkq/ponude.json?dl=0
/*
func readJSONFromUrl(url string) ([]Country, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var countryList []Country
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &countryList); err != nil {
		return nil, err
	}

	return countryList, nil
}

func main() {
	url := "https://www.dropbox.com/s/2kqweiiqf6nbhfz/lige.json?dl=0"
	//https://www.dropbox.com/s/wr9vnmt5e1jhwkq/ponude.json?dl=0
	countryList, err := readJSONFromUrl(url)
	if err != nil {
		panic(err)
	}

	for idx, row := range countryList {
		// skip header
		if idx == 0 {
			continue
		}

		if idx == 6 {
			break
		}

		fmt.Println(row.CountryName.Common)
	}
}
*/
