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
	Pleito     Pleito
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

func (m *model) TogglePleito() {
	m.Pleito = 1 - m.Pleito
}

func (m model) Init() tea.Cmd {
	return Load
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// the Load cmd returns new results
	case []Results:
		m.Results = msg[:2]

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
		case "tab":
			m.TogglePleito()
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

	var pleito string
	if m.Pleito == federal {
		pleito = "Presidente"
	} else {
		pleito = "Governador"
	}

	var s string
	s += "\n"
	s += boldStyle.Render(printer.Sprint(title, " – ", pleito))
	s += printer.Sprintf("\n\n%s", m.ToString())

	s += printer.Sprintf("Total de votos apurados: %d", SumVotes(m.Results))

	s += "\n\n"
	s += helpStyle.Render(printer.Sprintf("Última atualização: %v", m.LastUpdate.Format("02/01/2006 15:04:05")))
	s += helpStyle.Render("\n'r': atualizar\t\t'q': sair")

	return s
}

func main() {
	p := tea.NewProgram(model{Pleito: federal})
	err := p.Start()
	if err != nil {
		panic(err)
	}
}
