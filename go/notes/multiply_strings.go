package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func MultiplyStrings(num1, num2 string) string {
	op1, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal(err)
	}

	op2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal(err)
	}

	prod := strconv.Itoa(op1 * op2)
	return prod
}

func main() {

	n1, n2 := "2", "3"
	ans := MultiplyStrings(n1, n2)
	fmt.Fprintf(os.Stderr, "DEBUG[3]: multiply_strings.go:29: ans=%+v\n", ans)

	num1, num2 := "123", "456" //"56088"
	ans2 := MultiplyStrings(num1, num2)
	fmt.Fprintf(os.Stderr, "DEBUG[3]: multiply_strings.go:29: ans=%+v\n", ans2)
}
