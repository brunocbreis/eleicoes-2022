package main

import (
	"os"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return Load(m.url)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// the Load cmd returns new results
	case []Results:
		m.Results = msg
		m.UpdatePercentage()

		m.Loading = false
		m.LastUpdate = time.Now()

		return m, nil

	// key presses
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			m.Quitting = true

			return m, tea.Quit

		case key.Matches(msg, DefaultKeyMap.Refresh):
			m.Loading = true
			return m, Load(m.url)

		case key.Matches(msg, DefaultKeyMap.TogglePleito):
			m.TogglePleito()
			m.Loading = true

			return m, Load(m.url)
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

	if len(m.Results) == 0 {
		return "Carregando..."
	}

	// Give feedback if it takes too long to load
	var loading string
	if m.Loading {
		loading = "Carregando..."
	} else {
		loading = printer.Sprintf("Última atualização: %v", m.LastUpdate.Format("02/01/2006 15:04:05"))
	}

	// Render output
	var s string
	s += "\n"

	s += boldStyle.Render(printer.Sprint(title, " – ", m.Pleito.Name))
	s += printer.Sprintf(" (%s)", nomeUF[m.Pleito.local])
	s += printer.Sprintf("\n\n%s", m.ToString())

	s += printer.Sprintf("Total de votos apurados: %d", m.TotalVotes)

	s += "\n\n"
	s += helpStyle.Render(loading)
	s += "\n"
	s += helpView()

	return s
}

func main() {
	var estado Local

	if len(os.Args) > 1 {
		estado = Local(os.Args[1])
	}

	if _, ok := nomeUF[estado]; !ok {
		estado = "sp"
	}

	p := tea.NewProgram(NewModel(estado), tea.WithAltScreen())
	err := p.Start()
	if err != nil {
		panic(err)
	}
}
