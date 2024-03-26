package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	nome  string
	idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}
func (ps ParaNome) Less(i, j int) bool {
	return ps[i].nome < ps[j].nome
}
func (ps ParaNome) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {

	criancas := []Dados{
		{"Julia", 5},
		{"Pedro", 6},
		{"Ana", 4},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}
