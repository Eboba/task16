package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merger(massif [4]int, massifTwo [5]int) (massifThree [9]int) {

	for i, _ := range massifThree { // цикл объединяет два массива
		if i < len(massif) {
			massifThree[i] = massif[i]
		} else {
			massifThree[i] = massifTwo[i-4]
		}
	}

	return
}

func main() {
	fmt.Println("Задание 1. Слияние отсортированных массивов")

	var num [4]int
	var numTwo [5]int

	for i, _ := range num {
		fmt.Println("Введите число для первого массива")
		fmt.Scan(&num[i])

	}
	for i, _ := range numTwo {
		fmt.Println("Введите число для второго массива")
		fmt.Scan(&numTwo[i])

	}

	forward := 0
	back := len(num) - 1

	for { //Сортировка перемешиванием (англ. Cocktail sort).
		f := 0
		fmt.Println(num)
		for j := 0; j < len(num)-1; j++ {
			if num[j] > num[j+1] {
				f = 1
			}
		}
		if f == 0 {
			break
		}
		for i := forward; i < back; i++ {
			if num[i] > num[i+1] {
				num[i], num[i+1] = num[i+1], num[i]
				fmt.Println(num)
			}

		}
		forward++

		for i := back; i > forward-1; i-- {
			if num[i-1] > num[i] {
				num[i], num[i-1] = num[i-1], num[i]
				fmt.Println(num)
			}
		}
		back--
	}

	rand.Seed(time.Now().UnixNano())

	for { //  Bogosort (также глупая сортировка, stupid sort), кстать хотел задачу со скобками через рандом сделать, но не стал =)
		f := 0
		for j := 0; j < len(numTwo)-1; j++ {
			if numTwo[j] > numTwo[j+1] {
				f = 1
			}
		}
		if f == 0 {
			break
		}
		for i, _ := range num {
			z := rand.Intn(len(numTwo))
			numTwo[i], numTwo[z] = numTwo[z], numTwo[i]
		}
	}

	fmt.Println(merger(num, numTwo))
}
