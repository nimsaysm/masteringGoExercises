package chapter02

/*
#include <stdio.h>

static int sumNumbers(int a, int b) {
    int result = a + b;
    printf("result: %d\n", result);
    return result;
}
*/
import "C"

import "fmt"

func Runner() {
	fmt.Println("*** Running Chapter 02 ***")

	RunCProgram()
}

// RunCProgram runs a C program using Go
func RunCProgram() {
	fmt.Println("*** Running RunCProgram ***")
	result := C.sumNumbers(1, 2)
	fmt.Printf("Result from C: %d\n", int(result))
}
