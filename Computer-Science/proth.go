package main

import (
	"fmt"
	"os"
	"flag"
	"strconv"
	"math/big"
)

func isPower(num int) bool {
	return (num != 0 && (num & (num - 1)) == 0)
}

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

func main() {
	flag.Parse()
	 args := flag.Args()

	if len(args) < 1 {
	fmt.Println("Please specify a number")  
	os.Exit(1)  
	}

	n, err := strconv.Atoi(args[0])
    if err != nil {
        // handle error
        fmt.Println(err)
		os.Exit(2)
	}

	if (isProthNumber(n - 1)) {
	if big.NewInt(n).ProbablyPrime(0) {
		fmt.Println(n, "is a Proth Prime")
	} else {
		fmt.Println(n, "is a Proth Number but not a Proth Prime")
	}
	} else {
		fmt.Println("No, the value is not a Proth Number")
	}
}