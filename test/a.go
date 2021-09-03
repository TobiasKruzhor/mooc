package main

import "fmt"

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score : %d", score))
	}
	return g
}
func main() {
	fmt.Println(
		grade(33),
		grade(0),
		grade(66),
		grade(88),
		grade(1001),
	)
}
