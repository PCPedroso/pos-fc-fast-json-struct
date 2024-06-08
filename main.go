package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type Notas struct {
	Nota int
}

type Usuario struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Remoto bool   `json:"remoto"`
}

func main() {
	var parser fastjson.Parser
	jsonData := `{"usuario": {"id": 1, "nome": "José", "remoto":true}}`

	aluno, err := parser.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	usuarioJSON := aluno.Get("usuario").String()

	var usuario Usuario
	if err := json.Unmarshal([]byte(usuarioJSON), &usuario); err != nil {
		panic(err)
	}
	fmt.Println(usuario.Id, usuario.Nome)
}
