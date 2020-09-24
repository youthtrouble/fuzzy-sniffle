package main

import (
	"fmt"
	"os"
	"flag"
	"strconv"
)

// function to check power of two
func isPower(num int) bool {
	return (num != 0 && (num & (num - 1)) == 0)
}

// functon to check if given value n is a proth number or not
func isProthNumber(n int) bool {
	
	k := 1

	for (k < (n/k)) {

		if (n % k == 0) {
			if (isPower(n/k)) {
				return true
			}
		}
		k = k + 2
	}
	return false
}

// function to check if value is a prime number
func isPrime(value int) bool {

	for i := 2; i <= value/2; i++ {
		if value % i == 0 {
			return false
		}
	}
	return true 
}

func main() {
	flag.Parse()
	 args := flag.Args()

	// requesting for a number if none is specified as a flag
	if len(args) < 1 {
	fmt.Println("Please specify a number")  
	os.Exit(1)  
	}

	//converting the flag to an integer since args from the flag poackage are normally of type string
	n, err := strconv.Atoi(args[0])
    if err != nil {
        // handle error
        fmt.Println(err)
		os.Exit(2)
	}

	//checking if input is a proth number before checking if it is a proth prime
	if (isProthNumber(n - 1)) {
		if isPrime(n) {
			fmt.Println(n, "is a Proth Prime")
		} else {
			fmt.Println(n, "is a Proth Number but not a Proth Prime")
		}
	} else {
		fmt.Println("No, the value is not a Proth Number")
	}
}