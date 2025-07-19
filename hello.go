package main

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

func suma(a int, b int) int {
	return a + b
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	cociente := a / b
	//resto := float64(int(a) % int(b))

	return cociente, nil
}

func imprimirNombres(nombres ...string) {
	for nombres := range nombres {
		fmt.Println(nombres)
	}
}

// ejemplo de closure
func contador() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

func main() {
	// defer fmt.Println(quote.Go())
	fmt.Println("--------------")
	fmt.Println(quote.Glass())
	fmt.Println("--------------")
	fmt.Println(quote.Hello())
	fmt.Println("--------------")
	fmt.Println(quote.Opt())
	fmt.Println(suma(1, 2))
	cociente, err := dividir(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	}

	imprimirNombres("Julian", "Esteban", "Nico")
	fmt.Println(cociente)

	cont := contador()
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())
	fmt.Println("contador: ", cont())

	miRectangulo := Rectangulo{10, 10}
	fmt.Print("El area del rectangulo es: ", miRectangulo.Area())
}

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}
