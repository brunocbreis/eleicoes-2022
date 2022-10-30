package main

import (
	"fmt"
	"time"
)

type model struct {
	url        string
	uf         Local
	Results    []Results
	TotalVotes int
	LastUpdate time.Time
	Loading    bool
	Quitting   bool
	Pleito     Pleito
}

func NewModel(uf Local) model {
	m := model{Pleito: pleitoPresidente(), uf: uf}
	m.UpdateURL()

	return m
}

func (m *model) UpdateURL() {
	url := fmt.Sprintf("https://resultados.tse.jus.br/oficial/ele2022/%d/dados-simplificados/%s/%s-c000%d-e000%d-r.json", m.Pleito.codigo, m.Pleito.local, m.Pleito.local, m.Pleito.cargo, m.Pleito.codigo)
	m.url = url
}

func (m model) ToString() string {
	var s string
	for i := 0; i < maxCandidatos; i++ {
		res := m.Results[i]

		s += boldStyle.Render(res.Nome)
		s += "\n"
		s += res.Progress.ViewAs(res.Porcentagem)
		s += printer.Sprintf("\n%d votos", res.Votos)
		s += printer.Sprintf(" ⁕ %.2f", res.Porcentagem*100) + "%\n\n"
	}
	return s
}

func (m *model) UpdatePercentage() {
	m.SumVotes()

	res := m.Results

	for i := 0; i < len(res); i++ {
		votes := res[i].Votos
		res[i].Porcentagem = float64(votes) / float64(m.TotalVotes)
	}

}

func (m *model) SumVotes() {
	var total int
	for _, res := range m.Results {
		votes := res.Votos
		total = total + votes
	}

	m.TotalVotes = total
}

func (m *model) TogglePleito() {
	switch m.Pleito.Name {
	case prName:
		m.Pleito = pleitoEstado(m.uf)

	case govName:
		m.Pleito = pleitoPresidente()
	}
	m.UpdateURL()
}
