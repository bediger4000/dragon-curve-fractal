package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	generations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	findLimit(generations, complex(0, 0))
	findLimit(generations, complex(1, 0))
	fmt.Println()
	findLimit(generations, complex(1, -1))
}

var const1 complex64 = complex(1, 1)
var const2 complex64 = complex(1, -1)

func f1(z complex64) complex64 {
	return z * const1 / 2.0
}
func f2(z complex64) complex64 {
	return 1.0 - z*const1/2.0
}

func findLimit(ply int, x complex64) {
	if ply == 0 {
		fmt.Printf("%f\t%f\n", real(x), imag(x))
		return
	}
	x1 := f1(x)
	findLimit(ply-1, x1)
	x2 := f2(x)
	if x1 != x2 {
		findLimit(ply-1, x2)
	}
}
