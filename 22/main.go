// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
// значение которых > 2^20.
package main

import (
	"fmt"
	"math/big" // Пакет "math/big" позволяет работать с огромными числами
)

func main() {
	// Инициализируем наши большие переменные
	a := big.NewInt(20)
	b := big.NewInt(2)

	// Выполняем арифмитические операции

	mulResult := new(big.Int).Mul(a, b)
	fmt.Println(mulResult)

	divResult := new(big.Int).Div(a, b)
	fmt.Println(divResult)

	sumResult := new(big.Int).Add(a, b)
	fmt.Println(sumResult)

	subResult := new(big.Int).Sub(a, b)
	fmt.Println(subResult)
}
