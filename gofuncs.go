package main

import (
	"errors"
	"fmt"
)

func impuesto(salario int64) float64 {
	impuesto := 0.0
	if salario > 50000 {
		impuesto += 0.17
	}
	if salario > 150000 {
		impuesto += 0.10
	}

	return float64(salario) * impuesto
}

func promedio(notas ...uint) float64 {
	var sum uint = 0
	for _, nota := range notas {
		sum += nota
	}

	return float64(sum) / float64(len(notas))
}

func salario(minutos int, categoria string) float64 {
	horas := float64(minutos) / float64(60)
	switch categoria {
	case "A":
		return 3000 * horas * 1.50
	case "B":
		return 1500 * horas * 1.20
	case "C":
		return 1000 * horas
	}
	return 0
}

func operation(op string) (func(operandos ...int) float64, error) {
	switch op {
	case "minimum":
		return min, nil
	case "maximum":
		return max, nil
	case "average":
		return average, nil
	}

	return nil, errors.New("No such operator")
}

func min(operandos ...int) float64 {
	min := operandos[0]
	for _, value := range operandos {
		if value < min {
			min = value
		}
	}

	return float64(min)
}

func max(operandos ...int) float64 {
	max := operandos[0]
	for _, value := range operandos {
		if value > max {
			max = value
		}
	}

	return float64(max)
}

func average(notas ...int) float64 {
	sum := 0
	for _, nota := range notas {
		sum += nota
	}

	return float64(sum) / float64(len(notas))
}

var foodAmount = map[string]int{
	"dog":       10000,
	"cat":       5000,
	"hamster":   250,
	"tarantula": 150,
}

func animal(animal string) (func(amount int) int, error) {
	if _, ok := foodAmount[animal]; ok {
		return func(amount int) int { return amount * foodAmount[animal] }, nil
	}
	return nil, errors.New("Invalid animal")
}

func main() {
	fmt.Println("impuestos")
	fmt.Println(50001, impuesto(50001))
	fmt.Println(6912, impuesto(6912))
	fmt.Println(150001, impuesto(150001))

	fmt.Println("promedios")
	fmt.Println("10,8,2,5,7", promedio(10, 8, 2, 5, 7))

	fmt.Println("salarios")
	fmt.Println("3000, A", salario(3721, "A"))
	fmt.Println("3000, B", salario(3721, "B"))
	fmt.Println("3000, C", salario(3721, "C"))

}
