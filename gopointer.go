package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printval(slicepntr *[10]int) (err error) {
	for i := range slicepntr {
		fmt.Printf("Value in slicepntr[%d] is %v\n", i, slicepntr[i])
	}
	fmt.Printf("The address of pointer slicepointer is %v", &slicepntr)
	return nil
}

func main() {
	var intarray [10]int

	for i, val := range intarray {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		val = r.Intn(1000)
		intarray[i] = val
		fmt.Println("intarray", i, "is", intarray[i])
	}
	printval(&intarray)
}
