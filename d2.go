package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s generations real imag [offset x [offset y]]\n", filepath.Base(os.Args[0]))
		return
	}
	generations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	r, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(err)
	}
	i, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatal(err)
	}
	var offsetX, offsetY float64
	if len(os.Args) > 4 {
		offsetX, err = strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(os.Args) > 5 {
		offsetY, err = strconv.ParseFloat(os.Args[5], 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	z := complex(r, i)
	findLimit(generations, z, offsetX, offsetY)
}

var const1 complex128 = complex(1, 1)
var const2 complex128 = complex(1, -1)

func f1(z complex128) complex128 {
	return z * const1 / 2.0
}
func f2(z complex128) complex128 {
	return 1.0 - z*const2/2.0
}

func findLimit(ply int, x complex128, offsetX, offsetY float64) {
	if ply == 0 {
		fmt.Printf("%f\t%f\n", real(x)+offsetX, imag(x)+offsetY)
		return
	}
	x1 := f1(x)
	x2 := f2(x)
	findLimit(ply-1, x1, offsetX, offsetY)
	if x1 != x2 {
		findLimit(ply-1, x2, offsetX, offsetY)
	}
}
