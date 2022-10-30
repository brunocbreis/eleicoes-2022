package main

import (
	"fmt"
	"time"
)

type model struct {
	url        string
	Results    []Results
	TotalVotes int
	LastUpdate time.Time
	Loading    bool
	Quitting   bool
	Pleito     Pleito
}

func NewModel() model {
	m := model{Pleito: defaultPleito}
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
		m.Pleito.cargo = governador
		m.Pleito.codigo = gov2T
		m.Pleito.local = sp
		m.Pleito.Name = govName

	case govName:
		m.Pleito.cargo = presidente
		m.Pleito.codigo = pres2T
		m.Pleito.local = br
		m.Pleito.Name = prName
	}
	m.UpdateURL()
}
