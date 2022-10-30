package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return Load
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// the Load cmd returns new results
	case []Results:
		m.Results = msg[:2]
		m.UpdatePercentage()

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
	// Clear screen before quitting
	if m.Quitting {
		return ""
	}

	// Give feedback if it takes too long to load
	var loading string
	if m.Loading {
		loading = "Carregando..."
	} else {
		loading = printer.Sprintf("Última atualização: %v", m.LastUpdate.Format("02/01/2006 15:04:05"))
	}

	// Pleito
	var pleito string
	if m.Pleito == federal {
		pleito = "Presidente"
	} else {
		pleito = "Governador"
	}

	// Render output
	var s string
	s += "\n"

	s += boldStyle.Render(printer.Sprint(title, " – ", pleito))
	s += printer.Sprintf("\n\n%s", m.ToString())

	s += printer.Sprintf("Total de votos apurados: %d", m.TotalVotes)

	s += "\n\n"
	s += helpStyle.Render(loading)
	s += helpStyle.Render("\n'r': atualizar\t\t'q': sair")

	return s
}

func main() {
	p := tea.NewProgram(NewModel())
	err := p.Start()
	if err != nil {
		panic(err)
	}
}
