package main

import (
	"sort"

	"github.com/charmbracelet/bubbles/progress"
)

type Results struct {
	Partido     Partido
	Nome        string
	Votos       int
	Porcentagem float64
	Progress    progress.Model
}

func (r Results) String() string {
	s := printer.Sprintf("\nCandidato: %s\nVotos: %d\n", r.Nome, r.Votos)
	return s
}

func Sort(r []Results) {
	sort.Slice(r, func(i, j int) bool { return r[i].Votos > r[j].Votos })
}
