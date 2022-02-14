package main

import (
	"fmt"
)

func main() {
	fmt.Println("Задание 2. Сортировка пузырьком")

	var num [6]int

	for i, _ := range num {
		fmt.Println("Введите число для массива")
		fmt.Scan(&num[i])
	}

	for j := 1; j != len(num); j++ {
		f := 0

		for i := 0; i != len(num)-j; i++ {
			if num[i] > num[i+1] {
				num[i], num[i+1] = num[i+1], num[i]
				f = 1
			}
		}
		if f == 0 {
			break
		}
	}
	fmt.Println(num)
}
