package main

import "testing"

// Triple A - Arrange, Act, Assert
// Arrange - Preparar o ambiente para o teste
// Act - Executar o teste
// Assert - Verificar se o resultado é o esperado

func TestSoma(t *testing.T) { // convenção de nomes ShouldSumCorrect (assinatura do método)
	// arrange
	teste := soma(1, 2, 3)
	// act
	resultado := 6

	// assert
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func TestSoma2(t *testing.T) { //ShouldSumIncorret
	//arrange
	teste := soma(3, 2, 1)
	//act
	resultado := 7

	//assert
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
