package main

import (
	"fmt"
	"strings"
	"strconv"
)
var Final int
func main() {
	var num string
	var cuisine string
	fmt.Println("*****Welcome to XYZ Cafe*****")
	fmt.Println("What would you like to have? \n      Here is Our Menu! \n=>To give an order , you need to enter the amount and then the code of item")

	fmt.Println("C : coffee, 40rs")
	fmt.Println("D : dosa, 80 rs")
	fmt.Println("T : tomato soup, 20rs")
	fmt.Println("I : idli , 20rs")
	fmt.Println("V : vada, 25rs.")
	fmt.Println("B : bhature &chhole, 30rs")
	fmt.Println("P : paneer pakoda, 30rs")
	fmt.Println("M : manchurian, 90rs")
	fmt.Println("H : hakka noodle, 70rs.")
	fmt.Println("F : French fries, 30rs")
	fmt.Println("J : jalebi, 30rs")
	fmt.Println("L : lemonade, 15rs")
	fmt.Println("S : spring roll, 20rs")

	for {
		fmt.Scan("%s",&num)

		pieces, err := strconv.Atoi(num)
		if err != nil {
			temp := strings.ToUpper(num)
			if temp == "END" {
				break
			} else {
				fmt.Println(err)
				continue
			}
		}
		fmt.Scan("%s", &cuisine)

		cuisine = strings.ToUpper(cuisine)

		if cuisine == "END" {
			break
		}

		switch cuisine {
		case "C":
			Final += 40 * pieces
			break

		case "D":
			Final += 80 * pieces
			break

		case "T":
			Final += 20 * pieces
			break

		case "I":
			Final += 20 * pieces
			break

		case "V":
			Final += 25 * pieces
			break

		case "B":
			Final += 30 * pieces
			break
		
		case "P":
			Final += 30 * pieces
			break

		case "M":
			Final += 90 * pieces
			break

		case "H":
			Final += 70 * pieces
			break

		case "F":
			Final += 30 * pieces
			break

		case "J":
			Final += 30 * pieces
			break

		case "L":
			Final += 15 * pieces
			break

		case "S":
			Final += 20 * pieces
			break

		}	
	}
	fmt.Println("Total amount to be paid is =", Final)

}
