package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historio string `json:"historia"`
}

var Personalidades = []Personalidade{} // variavel publica que será usada para moca
