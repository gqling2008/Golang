package main
import ("image"
		"image/color"
		"image/png"
		"math"
		"os"
		"log"
)
func main() {
	//图片大小
	const size = 300
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//遍历每个像素
	for x:= 0; x < size; x++ {
		for y:= 0; y < size; y++ {
			//填充为白色
			pic.SetGray(x, y, color.Gray{255})

		}
	}
	for x:= 0; x < size; x++ {
		
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{0})
	}
	//创建文件
	file, err := os.Create("sbin.png")

	if err != nil {
		log.Fatal(err)
	}
	//使用png格式将数据写入文件
	png.Encode(file, pic) //将image文件写入文件中
	//关闭文件
	file.Close()
}
//
