package main

import "fmt"

func main() {

	/*
		Conditional at Golang
	*/
	var currentYear = 2021

	var scores = 10

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	if age := currentYear - 1998; age < 18 {
		fmt.Println("")
	}

	/* switch scores {
	case 10:
		fmt.Println("Strike goals")
	case 9:
		fmt.Println("Nice Try")
	case 8:
		fmt.Println("goks")
	case 7:
		fmt.Println("perceft")
	default:
		fmt.Println("not bad")
	} */

	switch {
	case scores == 10:
		fmt.Println("Strike goals")
	case (scores < 10) && (scores > 3):
		fmt.Println("Not Bad")
	default:
		{
			fmt.Println("Try Again")
			fmt.Println("Needed learn")
		}
	}
}
