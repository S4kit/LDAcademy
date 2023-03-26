package main

import (
	"fmt"
	"os"      //used to control terminal variables
	"strconv" //manipulate enter vars to check if its int
)

func GCF(a, b int) int {
	if b == 0 { // GCF (a,0) or GCF (0,a) = a
		return a
	}

	return GCF(b, a%b) // trying to devid each time a%b to find common factors, till we reach thecondition of b=0 then GCF is a
}

func main() {

	if len(os.Args) < 3 { //checking if there's 2 vars, os.Args<3 because without vars len(os.args)=1
		fmt.Println("Please provide two integers as command-line arguments.") //we could use panic("Please provide two valid integers as command-line arguments.") in this situation but won't be using os.Exit()
		os.Exit(1)
	}

	a, err1 := strconv.Atoi(os.Args[1]) //saving first var in A and its Error return from str.Conv.Atoi
	b, err2 := strconv.Atoi(os.Args[2]) //saving first var in B and its Error return from str.Conv.Atoi

	if err1 != nil || err2 != nil { //if there is error in any of A or B
		fmt.Println("Please provide two valid integers as command-line arguments.") //we could use panic("Please provide two valid integers as command-line arguments.") in this situation but won't be using os.Exit()
		os.Exit(1)
	}

	fmt.Printf("The Greatest common factor of %d and %d is %d.\n", a, b, GCF(a, b)) //result
}
