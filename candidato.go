package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Candidato struct {
	Numero int
	Nome   string
}

type Results struct {
	Candidato Candidato
	Votos     int
}

func (r Results) String() string {
	p := message.NewPrinter(language.BrazilianPortuguese)
	s := p.Sprintf("\nCandidato: %s\nVotos: %d\n", r.Candidato.Nome, r.Votos)
	return s
}
