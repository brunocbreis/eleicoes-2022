package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var lang = language.BrazilianPortuguese
var printer = message.NewPrinter(lang)

const title = "Resultado Eleições 2022"

// PARA SELECIONAR O TURNO DAS ELEIÇÕES
var pleito = 544 // segundo turno = 545
var url = fmt.Sprintf("https://resultados.tse.jus.br/oficial/ele2022/%d/dados-simplificados/br/br-c0001-e000%d-r.json", pleito, pleito)

/* STYLING */
// Lula Style
var (
	colorLula = "#DE0000"
	progLula  = progress.New(progress.WithSolidFill(colorLula))
)

// Bolsonaro Style
var (
	colorBolsonaro = "#34923F"
	progBolsonaro  = progress.New(progress.WithSolidFill(colorBolsonaro))
)

// Help Stype
var helpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241"))
