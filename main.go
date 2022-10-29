package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Results    []Results
	LastUpdate time.Time
	Quitting   bool
}

func (m model) Init() tea.Cmd {
	return Load
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case []Results:
		m.Results = msg
		m.LastUpdate = time.Now()
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Quitting = true
			return m, tea.Quit
		case "r":
			return m, Load
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.Quitting {
		return ""
	}
	s := fmt.Sprintf("\n\n%v", m.Results)
	s += "\n\n"
	s += fmt.Sprintf("Última atualização: %v", m.LastUpdate.Format("02/01/2006 15:04:05"))
	return s
}

func main() {
	p := tea.NewProgram(model{})
	err := p.Start()
	if err != nil {
		panic(err)
	}
}
