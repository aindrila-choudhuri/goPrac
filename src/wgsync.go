package wgsync

import (
	"fmt"
	"sync"
)

var arragedNumbers []int

func WgSyncMain() {

	var wg sync.WaitGroup

	numbers := []int{1, 2, 3, 4, 5, 6}

	wg.Add(1)
	go printOdd(numbers, &wg)

	wg.Add(1)
	go printEven(numbers, &wg)

	wg.Wait()

	fmt.Println("arragedNumbers: ", arragedNumbers)

}

func printOdd(nums []int, wg *sync.WaitGroup) {
	for _, n := range nums {
		if n%2 != 0 {
			arragedNumbers = append(arragedNumbers, n)
		}

	}

	wg.Done()

}

func printEven(nums []int, wg *sync.WaitGroup) {
	for _, n := range nums {
		if n%2 == 0 {
			arragedNumbers = append(arragedNumbers, n)
		}

	}

	wg.Done()
}
