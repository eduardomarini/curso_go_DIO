package main

import "testing"

// padrão triple A -> AAA
// Arrange -> Organizar
// Act -> Agir
// Assert -> Afirmar (verificar as asserções)

func ShouldSumCorrect(t *testing.T) { // convenção de nomes ShouldSumCorrect (assinatura do método)
	// arrange
	teste := soma(1, 2, 3)
	// act
	resultado := 6
	// assert
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func ShouldSumIncorret(t *testing.T) { //ShouldSumIncorret
	// arrange
	teste := soma(3, 2, 1)
	// act
	resultado := 7

	// assert
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
