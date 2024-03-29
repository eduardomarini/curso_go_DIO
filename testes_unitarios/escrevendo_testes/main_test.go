package main

import "testing"

func ShouldSumCorrect(t *testing.T) { // convenção de nomes ShouldSumCorrect (assinatura do método)
	teste := soma(1, 2, 3)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func ShouldSumIncorret(t *testing.T) { //ShouldSumIncorret
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func ShouldMultiplyCorrect(t *testing.T) { //ShouldMultiplyCorrect
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func ShouldMultiplyIncorrect(t *testing.T) { //ShouldMultiplyCorrect
	teste := multiplica(10, 10)
	resultado := 1000

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}
