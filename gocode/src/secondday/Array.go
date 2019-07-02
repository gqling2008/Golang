package main
import (
	"fmt"
)



func main()  {
	var HighRiseBuilding [30]int
	for i:=0; i < 30; i++ {
		HighRiseBuilding[i] = i+1
	}
	//fmt.Println(team)
	fmt.Println(HighRiseBuilding[10:15])
	fmt.Println(HighRiseBuilding[20:])
	fmt.Println(HighRiseBuilding[:2])
	var team [3]string
	team[0] = "nanjing"
	team[1] = "beijing"
	team[2] = "shanghai"
	for k, v := range team {
	fmt.Println(k, v)
	
}
	fmt.Println(team)
	a := make([]int, 2)
	b := make([]int, 2, 10)
	var c []string
	c = append(c, "qin", "hong","nanjing")
	var d []int
	d = append(d, 3)
	fmt.Println(a, b, c, d)
	fmt.Println(len(a), len(b))

	seq := []string{"1","3", "4","5", "2", "6"}
	index := 3
	fmt.Println(seq[:index], seq[index+1:])
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
	var test = 11
	if test > 10 {
		fmt.Println("test > 10")
	} else {
		fmt.Println("test <= 10")
	}
	step := 2
	for ; step >0; step-- {
		fmt.Println(step)
	}
}