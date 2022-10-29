package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var pleito = 544 // segundo turno = 545
var url = fmt.Sprintf("https://resultados.tse.jus.br/oficial/ele2022/%d/dados-simplificados/br/br-c0001-e000%d-r.json", pleito, pleito)

func getResults(url string, target interface{}) error {

	// getting response
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching results.", err)
		os.Exit(1)
	}

	// getting bytes from response body
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// decoding to json (modifies target)
	return json.Unmarshal(bytes, target)
}

func Load() tea.Msg {
	// Retrieve raw JSON Data
	var r map[string]interface{}
	getResults(url, &r)
	candInts := r["cand"].([]interface{})

	// Turn it into maps of candidates
	var candMaps []map[string]interface{}
	for _, c := range candInts {
		candMaps = append(candMaps, c.(map[string]interface{}))
	}

	// Clean stuff up
	var res []Results
	for _, m := range candMaps {
		// Clean up Names
		name := m["nm"].(string)
		// name = strings.ToLower(name)

		caser := cases.Title(language.BrazilianPortuguese)
		name = caser.String(name)

		// Clean Felipe D'Ávila
		name = strings.Replace(name, "&Apos;", "'", 1)

		// Clean up numbers
		num, _ := strconv.Atoi(m["n"].(string))
		votes, _ := strconv.Atoi(m["vap"].(string))

		// Create candidates
		cd := Candidato{Nome: name, Numero: num}

		// Put it all together
		res = append(res, Results{Candidato: cd, Votos: votes})
	}

	return res
}
