package main

import "testing"

func TestSoma(t *testing.T) { // convenção de nomes ShouldSumCorrect (assinatura do método)
	teste := soma(1, 2, 3)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func TestSoma2(t *testing.T) { //ShouldSumIncorret
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}
