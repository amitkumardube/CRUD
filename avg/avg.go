package main

//import "fmt"

func Avg(nos ...int) int {
	sum := 0
	for _, num := range nos {
		sum += num
	}
	if sum == 0 {
		return 0
	}
	return sum / len(nos)
}

func main() {
//	nos := []int{1, 2, 3, 4, 5, 9}
//	fmt.Println(Avg(nos...))
//	fmt.Println(Avg(1,2,3,4,5))
}