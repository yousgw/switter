package main

import "fmt"

func main(){
	var a,b int
	a = 3
	b = 4
	switch a {
	case 1:
		fmt.Println("failed.")
	case 3:
		if b != 5 {
			fmt.Println("b is not 5.")

		} else {
			goto nextCase
		}
		break//明示しないとnextCaseは実行される
	nextCase:
		fallthrough//条件無視して通過
	case 5:
		fmt.Println("succsess.")

	default:
		fmt.Println("failed.")

	}
}
