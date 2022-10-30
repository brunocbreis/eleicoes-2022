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

// UFs com segundo turno
const (
	br Local = "br"
	al Local = "al"
	am Local = "am"
	ba Local = "ba"
	es Local = "es"
	ms Local = "ms"
	pb Local = "pb"
	pe Local = "pe"
	rs Local = "rs"
	rr Local = "rr"
	sc Local = "sc"
	se Local = "se"
	sp Local = "sp"
)

var nomeUF = map[Local]string{
	br: "Brasil",
	al: "Alagoas",
	am: "Amazonas",
	ba: "Bahia",
	es: "Espírito Santo",
	ms: "Mato Grosso do Sul",
	pb: "Paraíba",
	pe: "Pernambuco",
	rs: "Rio Grande do Sul",
	rr: "Roraima",
	sc: "Santa Catarina",
	se: "Sergipe",
	sp: "São Paulo",
}

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
