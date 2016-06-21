package main

/*
#include <stdio.h>
extern void myPrint(int i);

static inline void fibonacci(int m) {
	int i, n, f = 0 , s = 1;
	for (i=0;i<m;i++) {
		if ( i <= 1) {
			n = i;
		} else {
			n = f + s;
			f = s;
			s = n;
		}
		myPrint(n);

	}
}
*/
import "C"
import "fmt"

//export myPrint
func myPrint(i C.int) {
	fmt.Println(uint32(i))
}

func main() {
	C.fibonacci(10)
}
