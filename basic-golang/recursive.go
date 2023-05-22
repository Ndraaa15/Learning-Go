package main

func recursive(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * recursive(num-1)
	}
}

func faktorial(num int) int {
	result := 1
	for i := num; i >= 1; i-- {
		result *= i
	}
	return result
}

// func main() {
// 	result := faktorial(10)
// 	fmt.Println(result)
// 	fmt.Println(recursive(10))
// }
