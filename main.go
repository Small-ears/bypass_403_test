package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sync"

	"golang.com/golang.com/bypass_403/plug"
)

func main() {
	//接收用户传入参数
	var method, url, urlDir, filePath string
	var goroutineNum int
	flag.StringVar(&url, "u", "", "Target URL,example:https://example.com")
	flag.StringVar(&urlDir, "d", "", "target directory,example:path")
	flag.StringVar(&method, "m", "GET", "HTTP request method,Default is GET")
	flag.StringVar(&filePath, "f", "", "payload file path")
	flag.IntVar(&goroutineNum, "g", 10, "goroutine number,Default is 10.")
	flag.Parse()

	if url == "" || filePath == "" {
		flag.Usage()
		os.Exit(0) //状态码 0 通常表示程序正常退出，而非零的状态码表示程序异常退出或错误终止
	}

	// 读取文件中的字符串
	lines, err := plug.FileRead(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines = append(lines, "X-Original-URL:"+urlDir, "X-Rewrite-URL:"+urlDir)

	// 正则匹配合适的 payloads，匹配合适的 HTTP 请求方法
	reg := `^[a-zA-Z]+[-:]` // 匹配以字母开头的，包含-：的字符串
	regex, err := regexp.Compile(reg)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup                                      // 并发计数器，管理并发
	resultChan := make(chan *plug.ResponseResult, len(lines))  // 用于存储 ResponseResult
	summaryResultChan := make(chan *SummaryResult, len(lines)) // 用于存储 SummaryResult
	workerChan := make(chan string, goroutineNum)              // 控制并发大小

	// 启动 goroutine 池
	for i := 0; i < goroutineNum; i++ {
		wg.Add(1) // wg.Add(1) 应该在启动 goroutine 之前执行
		go worker(workerChan, resultChan, summaryResultChan, &wg, method, url, urlDir, regex)
	}

	// 向 workerChan 发送任务，将 payload 写入 workerChan 中，最多只能写 goroutineNum 个
	for _, line := range lines {
		workerChan <- line
	}

	close(workerChan)

	// 等待所有任务完成
	wg.Wait()
	// 关闭 resultChan
	close(resultChan)
	// 关闭 summaryResultChan
	close(summaryResultChan)

	// 收集结果
	var results []*plug.ResponseResult
	for r := range resultChan {
		results = append(results, r)
	}

	// 收集 SummaryResult，后两个函数执行的结果
	var summaryResults []*SummaryResult
	for sr := range summaryResultChan {
		summaryResults = append(summaryResults, sr)
	}

	//汇总
	for _, v := range summaryResults {
		results = append(results, v.ResultOne, v.ResultTwo)
	}

	// 将结果写入 CSV 中
	plug.ResultWrite(results)
}

type SummaryResult struct {
	ResultOne *plug.ResponseResult
	ResultTwo *plug.ResponseResult
}

func worker(workerChan <-chan string, resultChan chan<- *plug.ResponseResult, summaryResultChan chan<- *SummaryResult, wg *sync.WaitGroup, method, url, urlDir string, regex *regexp.Regexp) {
	defer wg.Done()

	for line := range workerChan {
		var result, resultOne, resultTwo *plug.ResponseResult
		var err error

		if regex.MatchString(line) {
			result, err = plug.SendHttpRequest_head(line, method, url, urlDir)
		} else {
			result, err = plug.SendHttpRequest_path(line, method, url, urlDir)
			//不能用result会覆盖前面的结果
			resultOne, _ = plug.SendHttpRequest_path_one(line, method, url, urlDir)
			resultTwo, _ = plug.SendHttpRequest_path_two(line, method, url, urlDir)
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		summaryResult := &SummaryResult{
			ResultOne: resultOne,
			ResultTwo: resultTwo,
		}

		resultChan <- result
		// // 只有在 ResultOne 或 ResultTwo 不为 nil 时才追加进去（部分payload带入会导致报错，因此返回nil）
		if summaryResult.ResultOne != nil || summaryResult.ResultTwo != nil {
			summaryResultChan <- summaryResult
		}
	}
}
