1、实现将payload从文件中读取

2、实现HTTP请求功能（路径以及HTTP Header上的FUZZ）

3、实现将结果存入CSV文件中

4、使用消费者模型实现并发


其他细节：

1、特殊（lines = append(lines, "X-Original-URL:"+urlDir, "X-Rewrite-URL:"+urlDir)）追加到payload切片中

2、正则匹配payload，将合适的payload传入合适的HTTP request中

3、payloads需要自己制作

4、路径上的fuzz包括左边、右边以及左右两边

5、user-agent使用PC浏览器user-agent

6、需要传入的参数有（method、url、urlDir、filepath（payloads文件），并发大小）

7、执行完成后会在当前目录生成一个result.csv


使用指南：

  -u string
  
        Target URL,example:https://example.com
        
  -d string
  
        target directory,example:path
        
  -m string
  
        HTTP request method,Default is GET (default "GET")
        
  -f string
  
        payload file path
        
  -g int
  
        goroutine number,Default is 10. (default 10)
