package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Results    []Results
	LastUpdate time.Time
	Loading    bool
	Quitting   bool
}

func (m model) ToString() string {
	var s string
	for _, res := range m.Results {
		s += res.Nome
		s += printer.Sprintf("\nVotos: %d", res.Votos)
		s += printer.Sprintf("\n%.2f", res.Porcentagem*100) + "%\n\n"
	}
	return s
}

func (m model) Init() tea.Cmd {
	return Load
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// the Load cmd returns new results
	case []Results:
		m.Results = msg[:2]
		m.Results = UpdatePercentage(m.Results)

		m.Loading = false
		m.LastUpdate = time.Now()

		return m, nil

	// key presses
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Quitting = true
			return m, tea.Quit

		case "r":
			m.Loading = true
			return m, Load
		}
	}

	// else
	return m, nil
}

func (m model) View() string {
	if m.Loading {
		return "Carregando..."
	}
	if m.Quitting {
		return ""
	}

	var s string
	s += "\n"
	s += title
	s += printer.Sprintf("\n\n%s", m.ToString())

	s += printer.Sprintf("Total de votos: %d", SumVotes(m.Results))

	s += "\n\n"
	s += helpStyle.Render(printer.Sprintf("Última atualização: %v", m.LastUpdate.Format("02/01/2006 15:04:05")))
	s += helpStyle.Render("\n'r': atualizar\t\t'q': sair")

	return s
}

func main() {
	p := tea.NewProgram(model{})
	err := p.Start()
	if err != nil {
		panic(err)
	}
}
