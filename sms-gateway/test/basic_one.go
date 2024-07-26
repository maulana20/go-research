package basic_one

import "fmt"

type fn func (string) string

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func getFullName (fn fn) (string, string) {
	return "Maulana", fn("Saputra")
}

func sumAll (numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	name := "Maulana Saputra"
	fmt.Println(reverse(name))

	firstName, lastName := getFullName(reverse)
	fmt.Println(firstName, lastName, "Antoneli")

	numbers := []int{10, 20, 30}
	total := sumAll(numbers...)
	fmt.Println(total)
}