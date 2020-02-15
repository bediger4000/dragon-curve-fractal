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
	findLimit(generations, complex(0, 0), complex(1, 0))
}

var const1 complex64 = complex(1, 1)
var const2 complex64 = complex(1, -1)

func f1(z complex64) complex64 {
	return z * const1 / 2.0
}
func f2(z complex64) complex64 {
	return 1.0 - z*const2/2.0
}

func findLimit(ply int, x1, x2 complex64) {
	if ply == 0 {
		fmt.Printf("%f\t%f\n", real(x1), imag(x1))
		fmt.Printf("%f\t%f\n", real(x2), imag(x2))
		return
	}
	findLimit(ply-1, f1(x1), f2(x1))
	findLimit(ply-1, f1(x2), f2(x2))
}
