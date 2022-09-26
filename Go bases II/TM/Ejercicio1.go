package main

import "fmt"

func calcularImpuesto(sueldo float64) float64 {
	switch {
	case sueldo > 150000:
		return sueldo * 0.27
	case sueldo > 50000:
		return sueldo * 0.17
	}
	return 0
}

func main() {
	fmt.Println("Impuesto a descontar: $", calcularImpuesto(160000))
	fmt.Println("Impuesto a descontar: $", calcularImpuesto(80000))
	fmt.Println("Impuesto a descontar: $", calcularImpuesto(25000))
}
