package main

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var lang = language.BrazilianPortuguese
var printer = message.NewPrinter(lang)

const title = "Resultado Eleições 2022"

const maxCandidatos = 2

var defaultPleito = Pleito{
	Name:   prName,
	local:  br,
	cargo:  presidente,
	codigo: pres2T,
}

/* STYLING */
// Lula Style
var (
	colorLula = "#DE0000"
	progLula  = progress.New(progress.WithSolidFill(colorLula))
)

// Bolsonaro Style
var (
	colorBolsonaro = "#1D2FC7"
	progBolsonaro  = progress.New(progress.WithSolidFill(colorBolsonaro))
)

// Help Stype
var helpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241"))

// Bold Style
var boldStyle = lipgloss.NewStyle().
	Bold(true)

// Progress bars
var emptyProg = []rune("-")[0]
var emptyShowPercentage = false
