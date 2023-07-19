package sum

import "fmt"

func main() {
	fmt.Println(sum())           //
	fmt.Println(sum(3))          //
	fmt.Println(sum(1, 2, 3, 4)) //
}
func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total = total + v
	}
	return total
}
