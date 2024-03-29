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

func TestMultiplica(t *testing.T) { //ShouldMultiplyCorrect
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func TestMultiplica2(t *testing.T) { //ShouldMultiplyCorrect
	teste := multiplica(10, 10)
	resultado := 1000

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func TestSub(t *testing.T) {
	teste := subtracao(10, 5)
	resultado := -15

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func TestSub2(t *testing.T) {
	teste := subtracao(5, 10)
	resultado := 15

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func TestDiv(t *testing.T) {
	teste := divide(20)
	resultado := 20

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)

	}
}

func TestDiv2(t *testing.T) {
	teste := divide(20)
	resultado := 10

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)

	}
}
