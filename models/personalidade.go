package models

type Personalidade struct {
	Nome     string `json:"nome"`
	Historio string `json:"historia"`
}

var Personalidades = []Personalidade{} // variavel publica que será usada para moca
