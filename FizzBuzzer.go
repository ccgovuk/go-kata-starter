package main

import "strconv"

func FizzBuzz(limit int) []string {

	results := make([]string, limit)
	for index := range results {
		number := index + 1
		results[index] = strconv.Itoa(number)
		if isFizz(number) && isBuzz(number){
			results[index] = "FizzBuzz"
		} else if isFizz(number) {
			results[index] = "Fizz"
		} else if isBuzz(number) {
			results[index] = "Buzz"
		}
	}
	return results
}

func isBuzz(i int) bool {
	return i%5 == 0
}

func isFizz(i int) bool {
	return i%3 == 0
}
