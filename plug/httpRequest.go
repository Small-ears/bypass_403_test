package plug

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 定义一个结构体存储请求结构
type ResponseResult struct {
	URL          string
	ReqHeader    []string
	StatusCode   int
	ResponseSize int
	ResponseBody string
}

// parameter_handling对传入的URL以及路径进行处理，保证符合预期格式
func parameter_handling(url, path string) (newUrl, newPath string) {

	//URL 参数处理
	if !strings.HasPrefix(url, "http") { //HasPrefix以http开头则返回true,区分大小写
		url = "https://" + url //没有http的添加http
	}
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}

	//Directory 参数处理
	// if !strings.HasPrefix(path, "/") {
	// 	path = "/" + path
	// }
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	return url, path
}

// SendHttpRequest发起请求并获取结果
func SendHttpRequest_head(payload, method, url, directory string) (*ResponseResult, error) {

	newUrl, newDirectory := parameter_handling(url, directory)
	url = newUrl + newDirectory //将处理后的URL和directory拼接传入

	//header处理,使用 strings.Split 分割字符串
	payload = strings.TrimSpace(payload) //去除字符串中的空格
	// 分割 payload
	tmpStr := strings.Split(payload, ":")
	if len(tmpStr) != 2 {
		return nil, fmt.Errorf("invalid payload format: %s", payload)
	}
	httpHeader1 := strings.TrimSpace(tmpStr[0])
	httpHeader2 := strings.TrimSpace(tmpStr[1])

	//HTTP请求创建等
	req, err := http.NewRequest(method, url, nil) //创建HTTP请求
	if err != nil {
		return nil, err //必须返回结构体类型的值而不能返回nil,*ResponseResult取值则可以
	}

	//设置自定义HTTP Header
	req.Header.Set(httpHeader1, httpHeader2) //有set和add方法，set要好些
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	client := &http.Client{
		Timeout: 3 * time.Second, //设置请求超时
		Transport: &http.Transport{ //设置忽略证书验证
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	//创建一个 HTTP 客户端
	resp, err := client.Do(req) //执行请求
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //释放相关资源和关闭底层连接

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	//提取info，组装结果
	tmpURL := req.URL.String() //将结构体中URL提取出来，转化为string
	var reqHeaders []string    //提取HTTP Header
	for k, v := range req.Header {
		headers := fmt.Sprintf("%s %s", k, v)
		reqHeaders = append(reqHeaders, headers)
	}

	tmpStatusCode := resp.StatusCode //获取状态码
	//tmpBodySize := resp.ContentLength //根据ContentLength获取响应主体大小
	tmpBodySize := len(body) //响应主体大小获取

	responseInfo := ResponseResult{
		URL:          tmpURL,
		ReqHeader:    reqHeaders,
		StatusCode:   tmpStatusCode,
		ResponseSize: tmpBodySize,
		ResponseBody: string(body), //强转类型
	}

	// 可以将多个HTTPResponseInfo对象存储到切片中
	return &responseInfo, nil
}

// SendHttpRequest发起请求并获取结果
func SendHttpRequest_path(payload, method, url, urlDir string) (*ResponseResult, error) {
	newURL, newDir := parameter_handling(url, urlDir)
	url = newURL + payload + newDir

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	tmpURL := req.URL.String()

	var reqHeaders []string
	for k, v := range req.Header {
		headers := fmt.Sprintf("%s %s", k, v)
		reqHeaders = append(reqHeaders, headers)
	}

	tmpStatusCode := resp.StatusCode
	tmpBodySize := len(body)

	responseInfo := ResponseResult{
		URL:          tmpURL,
		ReqHeader:    reqHeaders,
		StatusCode:   tmpStatusCode,
		ResponseSize: tmpBodySize,
		ResponseBody: string(body),
	}

	return &responseInfo, nil
}

func SendHttpRequest_path_one(payload, method, url, urlDir string) (*ResponseResult, error) {
	newURL, newDir := parameter_handling(url, urlDir)
	url = newURL + newDir + payload

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	tmpURL := req.URL.String()

	var reqHeaders []string
	for k, v := range req.Header {
		headers := fmt.Sprintf("%s %s", k, v)
		reqHeaders = append(reqHeaders, headers)
	}

	tmpStatusCode := resp.StatusCode
	tmpBodySize := len(body)

	responseInfo := ResponseResult{
		URL:          tmpURL,
		ReqHeader:    reqHeaders,
		StatusCode:   tmpStatusCode,
		ResponseSize: tmpBodySize,
		ResponseBody: string(body),
	}

	return &responseInfo, nil
}

func SendHttpRequest_path_two(payload, method, url, urlDir string) (*ResponseResult, error) {
	newURL, newDir := parameter_handling(url, urlDir)
	url = newURL + payload + newDir + payload

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	tmpURL := req.URL.String()

	var reqHeaders []string
	for k, v := range req.Header {
		headers := fmt.Sprintf("%s %s", k, v)
		reqHeaders = append(reqHeaders, headers)
	}

	tmpStatusCode := resp.StatusCode
	tmpBodySize := len(body)

	responseInfo := ResponseResult{
		URL:          tmpURL,
		ReqHeader:    reqHeaders,
		StatusCode:   tmpStatusCode,
		ResponseSize: tmpBodySize,
		ResponseBody: string(body),
	}

	return &responseInfo, nil
}
