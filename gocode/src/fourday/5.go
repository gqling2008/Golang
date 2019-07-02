package main
import ("fmt")
func foo(a int, b string) {
	//return a, b
	fmt.Println(a, b)
}

func add(a , b int) {
	//return a + b
	fmt.Println(a+b)
}
func main() {
	foo(100, "hello")
	add(100, 300)
}