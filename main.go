package main

import "fmt"

func main() {

	config := []int{0, 1, 2, 3, 4}

	// PREGUNTA 1

	fmt.Println("\n===== PREGUNTA 1 =====")

	case1 := []int{
		0, 1, 2, 3, 4,
		0, 1, 2, 3, 4,
		0, 1, 2, 3, 4,
		0, 1, 2, 3, 4,
	}

	mtf(config, case1)

	// PREGUNTA 2

	fmt.Println("\n===== PREGUNTA 2 =====")

	case2 := []int{
		4, 3, 2, 1, 0,
		1, 2, 3, 4,
		3, 2, 1, 0,
		1, 2, 3, 4,
	}

	mtf(config, case2)

	// PREGUNTA 3 Y 4

	bestAndWorstCase()

	// PREGUNTA 5

	fmt.Println("\n===== PREGUNTA 5 (20 veces 2) =====")

	case5a := make([]int, 20)

	for i := range case5a {
		case5a[i] = 2
	}

	mtf(config, case5a)

	fmt.Println("\n===== PREGUNTA 5 (20 veces 3) =====")

	case5b := make([]int, 20)

	for i := range case5b {
		case5b[i] = 3
	}

	mtf(config, case5b)

	// PREGUNTA 6

	fmt.Println("\n===== PREGUNTA 6 IMTF =====")

	bestCase := make([]int, 20)
	for i := range bestCase {
		bestCase[i] = 0
	}

	worstCase := []int{
		4, 3, 2, 1, 0,
		4, 3, 2, 1, 0,
		4, 3, 2, 1, 0,
		4, 3, 2, 1, 0,
	}

	fmt.Println("\nIMTF - Mejor Caso")
	imtf(config, bestCase)

	fmt.Println("\nIMTF - Peor Caso")
	imtf(config, worstCase)
}