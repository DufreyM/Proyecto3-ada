package main

import "fmt"

func moveToFront(list []int, item int) []int {

	pos := -1

	for i, v := range list {
		if v == item {
			pos = i
			break
		}
	}

	if pos <= 0 {
		return list
	}

	list = append(list[:pos], list[pos+1:]...)
	list = append([]int{item}, list...)

	return list
}

func mtf(config []int, requests []int) int {

	list := make([]int, len(config))
	copy(list, config)

	totalCost := 0

	fmt.Println("Configuración inicial:", list)
	fmt.Println()

	for _, req := range requests {

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

		list = moveToFront(list, req)

		fmt.Printf("Nueva lista: %v\n", list)
		fmt.Println("--------------------------------")
	}

	fmt.Printf("Costo total = %d\n", totalCost)

	return totalCost
}