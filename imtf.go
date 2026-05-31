package main

import "fmt"

func imtf(config []int, requests []int) int {

	list := make([]int, len(config))
	copy(list, config)

	totalCost := 0

	fmt.Println("Configuración inicial:", list)
	fmt.Println()

	for idx, req := range requests {

		cost := 0

		for i, v := range list {
			if v == req {
				cost = i + 1
				break
			}
		}

		totalCost += cost

		fmt.Printf("Lista: %v\n", list)
		fmt.Printf("Solicitud: %d\n", req)
		fmt.Printf("Costo: %d\n", cost)

		limit := idx + cost

		if limit > len(requests) {
			limit = len(requests)
		}

		move := false

		for _, next := range requests[idx+1 : limit] {
			if next == req {
				move = true
				break
			}
		}

		if move {
			list = moveToFront(list, req)
		}

		fmt.Printf("Nueva lista: %v\n", list)
		fmt.Println("--------------------------------")
	}

	fmt.Printf("Costo total IMTF = %d\n", totalCost)

	return totalCost
}