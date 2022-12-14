package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

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

func Load(url string) tea.Cmd {
	return func() tea.Msg {
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
			partido := Partido(num)

			votes, _ := strconv.Atoi(m["vap"].(string))

			// Put it all together
			switch name {

			case "Tarcísio":
				res = append(res, Results{Nome: "Tarcísio de Freitas", Partido: partido, Votos: votes, Progress: ProgressBar(coresPartidos[partido])})

			default:
				res = append(res, Results{Nome: name, Partido: partido, Votos: votes, Progress: ProgressBar(coresPartidos[partido])})
			}
		}

		// res = UpdatePercentage(res)

		for i := range res {
			res[i].Progress.Empty = emptyProg
			res[i].Progress.ShowPercentage = emptyShowPercentage
		}

		Sort(res)

		return res
	}
}
