package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merger(massif [4]int, massifTwo [5]int) (massifThree [9]int) {

	var a, b int

	for i, _ := range massifThree {
		if a == len(massif) { // первые два условия проверяют, чтоб не один из массивов не закончился
			massifThree[i] = massifTwo[b]
			b++
		} else if b == len(massifTwo) {
			massifThree[i] = massif[a]
			a++
		} else if massif[a] < massifTwo[b] { // следующие два условия проверяют, какой элемент положить в данный момент
			massifThree[i] = massif[a]
			a++
		} else {
			massifThree[i] = massifTwo[b]
			b++
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
			}
		}
		forward++
		for i := back; i > forward-1; i-- {
			if num[i-1] > num[i] {
				num[i], num[i-1] = num[i-1], num[i]
			}
		}
		back--
	}

	rand.Seed(time.Now().UnixNano())

	for { //  Bogosort (также глупая сортировка, stupid sort)
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
