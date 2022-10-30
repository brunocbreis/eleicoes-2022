package main

type Pleito struct {
	Name   string
	local  Local
	cargo  Cargo
	codigo Codigo
}

type Cargo int

const (
	_ Cargo = iota
	presidente
	_
	governador
)

type Local string

const (
	br Local = "br"
	sp Local = "sp"
)

// Nomes
const (
	prName  = "Presidente"
	govName = "Governador"
)

type Codigo int

const (
	pres1T Codigo = 544
	pres2T Codigo = 545
	gov1T  Codigo = 546
	gov2T  Codigo = 547
)
