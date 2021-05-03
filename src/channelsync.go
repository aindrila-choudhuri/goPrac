package channelsync

import (
	"fmt"
)

func ChannelSyncMain() {
	oddChannel := make(chan []int)

	numbers := []int{1, 2, 3, 4, 5, 6}

	go printOdd(numbers, oddChannel)

	go printEven(numbers, oddChannel)

	var arrangedNumbers []int

	for i := 0; i < 2; i++ {
		nums := <-oddChannel

		arrangedNumbers = append(arrangedNumbers, nums...)

	}
	fmt.Println(arrangedNumbers)

}

func printOdd(nums []int, oddChan chan<- []int) {
	var numbers []int

	for _, n := range nums {
		if n%2 != 0 {
			numbers = append(numbers, n)
		}

	}

	oddChan <- numbers
}

func printEven(nums []int, evenChan chan<- []int) {
	var numbers []int

	for _, n := range nums {
		if n%2 == 0 {
			numbers = append(numbers, n)
		}

	}

	evenChan <- numbers
}
