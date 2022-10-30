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

func pleitoPresidente() Pleito {
	return Pleito{
		Name:   prName,
		local:  "br",
		cargo:  presidente,
		codigo: pres2T,
	}
}

func pleitoEstado(uf Local) Pleito {
	return Pleito{
		Name:   govName,
		local:  uf,
		cargo:  governador,
		codigo: gov2T,
	}
}

/* STYLING */

func ProgressBar(color string) progress.Model {
	if color == "" {
		return progress.New()
	}
	return progress.New(progress.WithSolidFill(color))
}

// Help Style
var helpStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241"))

// Bold Style
var boldStyle = lipgloss.NewStyle().
	Bold(true)

// Progress bars
var emptyProg = []rune(" ")[0]
var emptyShowPercentage = false

type Partido int

var nomePartidos = map[Partido]string{
	15: "MDB",
	44: "União Brasil",
	13: "PT",
	40: "PSB",
	22: "PL",
	28: "PRTB",
	45: "PSDB",
	77: "Solidariedade",
	55: "PSD",
	10: "Republicanos",
}

var coresPartidos = map[Partido]string{
	13: "#DE0000",
	22: "#2A3591",
	10: "#FFDA05",
	45: "#004997",
	44: "#33BDF2",
	15: "#118C26",
}
