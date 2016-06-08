package main

/*
extern void print(int i);

void foo(void) {
	int i;
	for (i=0;i<10;i++) {
		myprintf(i);
	}
}
*/
import "C"
import "fmt"

//export print
func print(i C.int) {
	fmt.Println(uint32(i))
}

func main() {
	C.foo()
}
