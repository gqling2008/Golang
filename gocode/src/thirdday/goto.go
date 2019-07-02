package main
import ("fmt")
func main() {
	var breakflag bool
//外循环
	for x := 0; x < 10; x++ {
	//内循环
		for y := 0;y < 5; y++ {
		//满足某个条件时，退出nei循环
			if y ==5 {
				breakflag = true
				break
		}
	}
			if breakflag {
				break
	}

}
	fmt.Println("done")

}
