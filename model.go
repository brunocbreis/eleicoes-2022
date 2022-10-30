package main

import "time"

type model struct {
	Results    []Results
	TotalVotes int
	LastUpdate time.Time
	Loading    bool
	Quitting   bool
	Pleito     Pleito
}

func NewModel() model {
	return model{}
}

func (m model) ToString() string {
	var s string
	for _, res := range m.Results {
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
	m.Pleito = 1 - m.Pleito
}
