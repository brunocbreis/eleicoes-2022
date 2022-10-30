package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/text/cases"
)

func getResults(url string, target interface{}) error {
	c := &http.Client{Timeout: 10 * time.Second}

	// getting response
	res, err := c.Get(url)
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
		caser := cases.Title(lang)
		name = caser.String(name)

		// Clean up numbers
		num, _ := strconv.Atoi(m["n"].(string))
		votes, _ := strconv.Atoi(m["vap"].(string))

		// Put it all together
		switch name {
		case "Lula", "Fernando Haddad":
			res = append(res, Results{Nome: name, Numero: num, Votos: votes, Progress: progLula})

		case "Jair Bolsonaro", "Tarcísio de Freitas":
			res = append(res, Results{Nome: name, Numero: num, Votos: votes, Progress: progBolsonaro})

		default:
			res = append(res, Results{Nome: name, Numero: num, Votos: votes, Progress: progress.New()})
		}
	}

	res = UpdatePercentage(res)

	for i := range res {
		res[i].Progress.Empty = emptyProg
		res[i].Progress.ShowPercentage = emptyShowPercentage
	}

	return res
}

func UpdatePercentage(res []Results) []Results {
	total := SumVotes(res)

	for i := 0; i < len(res); i++ {
		votes := res[i].Votos
		res[i].Porcentagem = float64(votes) / float64(total)
		// res[i].Porcentagem = .25 + float64(i)*0.15
	}

	return res
}

func SumVotes(r []Results) int {
	var total int
	for _, res := range r {
		votes := res.Votos
		total = total + votes
	}

	return total
}
