package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func printval(slicepntr *[]int) (err error) {
	var slicearray = *slicepntr
	for i := range slicearray {
		fmt.Printf("Value in slicepntr[%d] is %v\n", i, slicearray[i])
	}
	fmt.Printf("The address of pointer slicepointer is %p\n", slicepntr)
	return nil
}

func main() {
	var (
		intarray     = []int{}
		intarrayptr  = &intarray
		arraysize    = 10
		val          int
		arraysizestr = "10"
		atoierr      error
	)
	fmt.Printf("Please enter desired size of array (default 10): ")
	if _, err := fmt.Scanf("%s", &arraysizestr); err != nil {
		fmt.Println(err)
	}
	if arraysize, atoierr = strconv.Atoi(arraysizestr); atoierr != nil {
		arraysize = 10
		fmt.Printf("Error converting string to integer. Defaulting to array size of 10.\n")
	}

	for i := 0; i < arraysize; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		val = r.Intn(1000)
		intarray = append(intarray, val)
		fmt.Println("intarray", i, "is", intarray[i])
	}
	fmt.Printf("Address for intarray is %p\n", intarrayptr)
	printval(&intarray)
}
