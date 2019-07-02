package main
import ("fmt")

const (
	 // 定义每分钟的秒数
	 SecondsPerMinute = 60
	 // 定义每小时的秒数
	 SecondsPerHour = SecondsPerMinute * 60
	 // 定义每天的秒数
	 SecondsPerDay = SecondsPerHour * 24
)
func resolveTime( seconds int) (day int, hour int, minute int) {
	day = seconds/SecondsPerDay
	hour = seconds/SecondsPerHour
	minute = seconds/SecondsPerMinute
	return
}

func main() {
	fmt.Println(resolveTime(10000))
	_,hour,minute := resolveTime(1800)
	fmt.Println(hour,minute)
	day,_,minute := resolveTime(600)
	fmt.Println(day,minute)
	//day,hour,_ := resolveTime(100)
	//fmt.Println(day,hour)

}