package logaritmo

import (
	"errors"
	"testing"
)

func TestVerificarSliceNxN(t *testing.T) {
	result := verificarSliceNxN("123456", 6)

	if result != true {
		t.Errorf("el resultado de verificar el slice de tamaño NxN no coincide con el esperado, obtuve: %t, esperaba: %t", result, true)
	}
}

func TestCantidadSecuencia(t *testing.T) {
	testCases := []struct {
		nombre   string
		adn      string
		esperado int
	}{
		{
			nombre:   "Ninguna coincidencia",
			adn:      "ATTTCGGA",
			esperado: 0,
		},
		{
			nombre:   "Una coincidencia",
			adn:      "ATGGAAAAC",
			esperado: 1,
		},
		{
			nombre:   "Tres coincidencias separadas",
			adn:      "AAAATGTTGCCCCATTTT",
			esperado: 3,
		},
		{
			nombre:   "Dos coincidencias juntas",
			adn:      "ATTGCGGGGGGGGAAT",
			esperado: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nombre, func(t *testing.T) {
			resultado := cantidadSecuencia(tc.adn)

			if resultado != tc.esperado {
				t.Errorf("el resultado no coincide con el esperado, resultado: %d, esperado: %d", resultado, tc.esperado)
			}
		})
	}
}

func TestEsMutante(t *testing.T) {
	testCases := []struct {
		nombre            string
		adn               []string
		esperadoResultado bool
		esperadoError     error
	}{
		{
			nombre:            "Error por no ser un slice NxN, de forma horizontal",
			adn:               []string{"ATCG", "TTTC", "TTTT", "TTT"},
			esperadoResultado: false,
			esperadoError:     sliceNxNError,
		},
		{
			nombre:            "Error por no ser un sline NxN, de forma vertical",
			adn:               []string{"ATCG", "TTTC", "TTTT", "AAAA", "GGAT"},
			esperadoResultado: false,
			esperadoError:     sliceNxNError,
		},
		{
			nombre:            "Error por contener letras incorrectas",
			adn:               []string{"ATTG", "AJGT", "LLLM", "AAAA"},
			esperadoResultado: false,
			esperadoError:     letrasPermitidasError,
		},
		{
			nombre:            "No es mutante",
			adn:               []string{"ATCG", "ACTT", "TAAA", "GGGA"},
			esperadoResultado: false,
			esperadoError:     nil,
		},
		{
			nombre:            "Es mutante, coincidencia vertical y horizontal",
			adn:               []string{"ATCGA", "AGTTA", "AAGAA", "AGGGA", "TTTTT"},
			esperadoResultado: true,
			esperadoError:     nil,
		},
		{
			nombre:            "Es mutante, coincidencia diagonal",
			adn:               []string{"AACGA", "ATATA", "TATAA", "AGGTA", "TGGTT"},
			esperadoResultado: true,
			esperadoError:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nombre, func(t *testing.T) {
			resultado, resultadoError := EsMutante(tc.adn)
			if resultado != tc.esperadoResultado || !errors.Is(resultadoError, tc.esperadoError) {
				t.Errorf("el resultado o el error no coincide con el esperado, resultado: %t, esperado: %t, resultado error: %s, error esperado: %s", resultado, tc.esperadoResultado, resultadoError.Error(), tc.esperadoError.Error())
			}
		})
	}
}
