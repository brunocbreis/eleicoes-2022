package main

import (
	"github.com/charmbracelet/bubbles/progress"
)

type Results struct {
	Numero      int
	Nome        string
	Votos       int
	Porcentagem float64
	Progress    progress.Model
}

func (r Results) String() string {
	s := printer.Sprintf("\nCandidato: %s\nVotos: %d\n", r.Nome, r.Votos)
	return s
}
