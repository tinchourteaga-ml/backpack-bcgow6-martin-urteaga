package Ejercicio3

import "errors"

func Dividir(n1, denominador int) (int, error) {

	if denominador == 0 {
		return 0, errors.New("el denominador no puede ser 0")
	}

	return n1 / denominador, nil
}
