package main
import ("fmt")
	

var string1 = "HELLO 你好"
func main() {
	for key, value := range string1 {
		fmt.Println(key, value)
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}
}	