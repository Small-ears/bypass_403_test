package plug

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// FileRead实现文件读取
func FileRead(filePath string) ([]string, error) {
	var str []string
	file, err := os.Open(filePath) //filepath.Clean能一定的规范路径格式
	if err != nil {
		fmt.Println("payloads file read erres:", err)
		return nil, err
	}

	defer file.Close() //关闭文件句柄

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF { //必须先判断
			str = append(str, strings.TrimSpace(line)) //文件末尾时，在处理 io.EOF 时会忽略了最后一行的添加
			break
		}

		if err != nil {
			return nil, err
		}

		str = append(str, strings.TrimSpace(line)) //使用 strings.TrimSpace 去除每行的首尾空白字符，包括\r\n
	}
	return str, err
}
