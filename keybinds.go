package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	TogglePleito key.Binding
	Refresh      key.Binding
	Quit         key.Binding
}

var DefaultKeyMap = KeyMap{
	TogglePleito: key.NewBinding(
		key.WithKeys("tab"),           // actual keybindings
		key.WithHelp("tab", "pleito"), // corresponding help text
	),
	Refresh: key.NewBinding(
		key.WithKeys("r", "F5"),
		key.WithHelp("r", "atualizar"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "sair"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.TogglePleito, k.Refresh, k.Quit}
}

func helpView() string {
	return help.NewModel().ShortHelpView(DefaultKeyMap.ShortHelp())
}
