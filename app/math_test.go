package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(16, 16)

	if total != 32 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 32)
	}
}