package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("Helo World")
	fmt.Print("Enter string length: ")
	var length int
	fmt.Scanln(&length)
	if length == 0 {
		panic("Invalid length")
	}

	fmt.Println("Enter the number of random strings to generate: ")
	var times int
	fmt.Scanln(&times)

	if times == 0 {
		times = 1
	}

	fmt.Print("Enter the string parameters for number,lower case, upper case as a 3 digit number 0 for 'No' 1 for 'Yes': ")
	var typeq string
	fmt.Scanln(&typeq)

	generateString(length, times, typeq)
}

func generateString(length, times int, typeq string) {
	if len(typeq) == 0 {
		panic("Invalid Input. Enter 0's or 1's")
	}

	typearr := strings.Split(typeq, "")
	charSet := ""

	if typearr[0] == "1" {
		charSet += "0123456789"
	}
	if typearr[1] == "1" {
		charSet += "qwertyuiopasdfghjklzxcvbnm"
	}
	if typearr[2] == "1" {
		charSet += "QWERTYUIOPASDFGHJKLZXCVBNM"
	}

	if len(charSet) == 0 {
		panic("Invalid Input. Enter 0's or 1's")
	}
	generateStrings(length, times, charSet)
}

func generateStrings(length, times int, charSet string) {
	for i := 0; i < times; i++ {
		str := make([]string, length)
		charSetArr := strings.Split(charSet, "")

		for i := range str {
			idx := rand.Intn(len(charSetArr))
			str[i] = charSetArr[idx]
		}
		fmt.Println(strings.Join(str, ""))
	}
}
