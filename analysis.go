package main

import "fmt"

func mtfCost(config []int, requests []int) int {

	list := make([]int, len(config))
	copy(list, config)

	total := 0

	for _, req := range requests {

		pos := 0

		for i, v := range list {
			if v == req {
				pos = i + 1
				break
			}
		}

		total += pos

		list = moveToFront(list, req)
	}

	return total
}

func bestAndWorstCase() {

	config := []int{0, 1, 2, 3, 4}

	best := make([]int, 20)
	for i := range best {
		best[i] = 0
	}

	worst := []int{
		4, 3, 2, 1, 0,
		4, 3, 2, 1, 0,
		4, 3, 2, 1, 0,
		4, 3, 2, 1, 0,
	}

	fmt.Println("\n===== PREGUNTA 3 =====")

	fmt.Println("Secuencia:")
	fmt.Println(best)

	fmt.Println("Costo:")
	fmt.Println(mtfCost(config, best))

	fmt.Println("\n===== PREGUNTA 4 =====")

	fmt.Println("Secuencia:")
	fmt.Println(worst)

	fmt.Println("Costo:")
	fmt.Println(mtfCost(config, worst))
}