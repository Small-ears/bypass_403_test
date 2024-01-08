package plug

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ResultWrite(results []*ResponseResult) {
	//创建CSV文件
	file, err := os.Create("result.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//创建 CSV Writer
	write := csv.NewWriter(file)
	defer write.Flush()

	//写入CSV 头部标题
	header := []string{"URL", "StatusCode", "ReqHeader", "ResponseSize", "ResponseBody"}
	write.Write(header)

	// 遍历结构体切片，并将每个结构体写入 CSV 文件
	for _, result := range results {
		// 转换数据为字符串切片
		row := []string{
			result.URL,
			strconv.Itoa(result.StatusCode),
			// 将 ReqHeader 切片连接为一个字符串
			strings.Join(result.ReqHeader, ", "), //接受一个字符串切片和一个连接符，返回将切片元素连接起来的一个字符串
			strconv.Itoa(result.ResponseSize),    //使用 strconv.Itoa 来将其转换为字符串
			result.ResponseBody,
		}

		// 写入一行数据
		write.Write(row)
	}
	write.Flush()

	// 检查错误
	if err := write.Error(); err != nil {
		panic(err)
	}
}
