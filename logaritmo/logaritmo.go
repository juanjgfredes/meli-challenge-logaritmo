package logaritmo

import (
	"regexp"
)

var letrasPermitidasError = &mutanteError{msg: "el adn no puede contener letras que no sean: A, T, C, G"}
var sliceNxNError = &mutanteError{msg: "el adn tiene que ser de secuencias NxN"}

type mutanteError struct {
	msg string
}

func (e *mutanteError) Error() string {
	return e.msg
}

func EsMutante(adn []string) (bool, error) {
	if len(adn) < 4 {
		return false, nil
	}

	var fila int
	var columna = len(adn)
	var numCoincidencias int = 0
	var palabraVertical, palabraDiagonal string = "", ""

	patronLetrasPermitidas := "^[ATCG]*$"
	regexLetrasPermitadas := regexp.MustCompile(patronLetrasPermitidas) //regex que verifica que solo haya ATCG en las letras

	for iAdn := 0; iAdn < len(adn); iAdn++ {

		numCoincidencias += cantidadSecuencia(adn[iAdn])
		if numCoincidencias >= 2 {
			return true, nil
		}

		if iAdn == 0 {
			for iLetra := 0; iLetra < len(adn); iLetra++ {
				for recorrido := 0; recorrido < len(adn); recorrido++ {
					if iLetra == 0 { //Es para que verifique solo una vez y no varias veces innecesariamente
						fila = len(adn[recorrido])
						if fila != columna { //verifica si el adn ingresado es NxN, sino devuelve error
							return false, sliceNxNError
						}
						if !regexLetrasPermitadas.MatchString(adn[recorrido]) {
							return false, letrasPermitidasError
						}
					}

					palabraVertical += string(adn[recorrido][iLetra])
					iDiagonal := iLetra + recorrido
					if iDiagonal < len(adn) { //evita que ocurra un index out of range
						palabraDiagonal += string(adn[recorrido][(iDiagonal)])
					}
				}

				/*if !verificarSliceNxN(palabraVertical, numCaracteres) { //verifica si el adn ingresado es NxN, sino devuelve error
					return false, sliceNxNError
				}*/

				numCoincidencias += cantidadSecuencia(palabraVertical)
				numCoincidencias += cantidadSecuencia(palabraDiagonal)
				if numCoincidencias >= 2 {
					return true, nil
				}
				palabraDiagonal = ""
				palabraVertical = ""
			}
		} else {
			if (len(adn) - iAdn) >= 4 { //verifica que haya minimo 4 letras en diagonal posibles
				for recorrido := 0; recorrido < (len(adn) - iAdn); recorrido++ {
					palabraDiagonal += string(adn[iAdn:][recorrido][(recorrido)])
				}

				numCoincidencias += cantidadSecuencia(palabraDiagonal)
				if numCoincidencias >= 2 {
					return true, nil
				}
				palabraDiagonal = ""
			}
		}
	}
	return false, nil
}

func verificarSliceNxN(adn string, numCaracteres int) bool {
	if numCaracteres != len(adn) { //verifica si el adn ingresado es NxN, sino devuelve error
		return false
	} else {
		return true
	}
}

func cantidadSecuencia(adn string) int {
	patronLetrasContinuadas := "A{4}|T{4}|C{4}|G{4}"
	regexLetrasContinuadas := regexp.MustCompile(patronLetrasContinuadas) //regex que verifica si hay algunas de esas letras continuadas
	return len(regexLetrasContinuadas.FindAllStringIndex(adn, -1))
}
