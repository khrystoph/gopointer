package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printval(slicepntr *[]int) (err error) {
	var slicearray = *slicepntr
	for i := range slicearray {
		fmt.Printf("Value in slicepntr[%d] is %v\n", i, slicearray[i])
	}
	fmt.Printf("The address of pointer slicepointer is %p", slicepntr)
	return nil
}

func main() {
	var (
		intarray    = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		intarrayptr = &intarray
	)

	for i, val := range intarray {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		val = r.Intn(1000)
		intarray[i] = val
		fmt.Println("intarray", i, "is", intarray[i])
	}
	fmt.Printf("Address for intarray is %p\n", intarrayptr)
	printval(&intarray)
}
