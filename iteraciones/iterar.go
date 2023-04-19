package iteraciones

import "fmt"

func IterarInfinito() {
	for {
		fmt.Println("Hola")
	}
}

func IterarCondicion() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func IterarCOndicionSaltos() {
	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}

	fmt.Println("------------------------------------------")

	for i := 100; i > 100; i -= 5 {
		fmt.Println(i)
	}
}

func IterarCOndicionSaltosBreakContinue() {

	for i := 100; i > 1; i -= 5 {
		if i == 50 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("------------------------------------------")

	for i := 0; i < 100; i += 5 {
		if i == 45 {
			break
		}
		fmt.Println(i)
	}

}
