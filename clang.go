package main

/*
extern void Print(int i);

void foo(void) {
	int i;
	for (i=0;i<10;i++) {
		Print(i);
	}
}
*/
import "C"
import "fmt"

//export Print
func Print(i C.int) {
	fmt.Println(uint32(i))
}

func main() {
	C.foo()
}
