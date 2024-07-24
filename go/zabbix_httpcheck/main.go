// Description: HTTP requests check for Zabbix
// Author: Dong Guo

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Options 结构体定义了命令行参数
type Options struct {
	url        string  // URL 地址
	timeout    float64 // 超时时间，单位为秒
	content    string  // 预期的内容字符串
	auth       string  // 基本认证的用户名和密码
	value      bool    // 返回实际值而不是 0 和 1
	payload    string  // URL 编码的 HTTP POST 数据
	timeoutSet bool    // 是否由用户设置了超时时间
}

// parseOpts 解析命令行参数并返回一个 Options 结构体
func parseOpts() Options {
	var opts Options

	flag.StringVar(&opts.url, "u", "", "URL to GET or POST (default http://)")
	flag.Float64Var(&opts.timeout, "t", 10, "seconds before connection times out")
	flag.StringVar(&opts.content, "c", "", "string to expect in the content")
	flag.StringVar(&opts.auth, "a", "", "username:password on sites with basic authentication")
	flag.BoolVar(&opts.value, "V", false, "return actual value instead of 0 and 1")
	flag.StringVar(&opts.payload, "p", "", "URL encoded http POST data")

	// 示例模板
	exampleTemplate := `
	examples:
	  {prog} -u idc1-web1/health
	  {prog} -u http://idc1-web1/health
	  {prog} -u http://idc1-web1/health -c ok
	  {prog} -u http://idc1-web1/health -c ok -V
	  {prog} -u http://idc1-web1/health -c ok -t 2 -V
	  {prog} -u http://idc1-web2:3000
	  {prog} -u http://idc1-web3/login.php?page=redirect_string -a username:password -V
	  {prog} -u https://idc2-web1.yourdomain.com -V
	`
	// 替换示例模板中的占位符
	exampleText := strings.ReplaceAll(exampleTemplate, "{prog}", os.Args[0])

	// 去除前导制表符
	trimmedExampleText := trimLeadingTabs(exampleText)

	// 自定义用法信息
	flag.Usage = func() {
		fmt.Printf("usage: %s [-h] -u url [-t timeout] [-c content] [-a auth] [-V] [-p payload]\n", os.Args[0])
		fmt.Print(trimmedExampleText)
		fmt.Println("\noptions:")
		flag.PrintDefaults()
	}

	// 检查是否提供了命令行参数
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	flag.Parse()

	// 检查用户是否设置了超时时间
	timeoutSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "t" {
			timeoutSet = true
		}
	})

	// 检查是否提供了 URL
	if opts.url == "" {
		flag.Usage()
		os.Exit(2)
	}

	// 如果 URL 没有 http:// 或 https:// 前缀，则添加 http://
	if !strings.HasPrefix(opts.url, "http://") && !strings.HasPrefix(opts.url, "https://") {
		opts.url = "http://" + opts.url
	}

	opts.timeoutSet = timeoutSet
	return opts
}

// trimLeadingTabs 去除输入字符串中每行的前导制表符
func trimLeadingTabs(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, "\t")
	}
	return strings.Join(lines, "\n")
}

// getResults 根据提供的选项执行 HTTP 请求
func getResults(opts Options) int {
	client := &http.Client{
		Timeout: time.Duration(opts.timeout * float64(time.Second)),
	}

	var req *http.Request
	var err error

	// 创建 HTTP 请求
	if opts.payload != "" {
		req, err = http.NewRequest("POST", opts.url, strings.NewReader(opts.payload))
	} else {
		req, err = http.NewRequest("GET", opts.url, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return 2
	}

	// 设置基本认证
	if opts.auth != "" {
		authParts := strings.Split(opts.auth, ":")
		if len(authParts) != 2 {
			fmt.Println("Invalid auth format. Expected username:password")
			return 2
		}
		username, password := authParts[0], authParts[1]
		req.SetBasicAuth(username, password)
	}

	startTime := time.Now()
	resp, err := client.Do(req) // 发送 HTTP 请求
	responseSecs := time.Since(startTime).Seconds()

	// 处理请求错误
	if err != nil {
		if opts.value {
			fmt.Println(err.Error())
		} else {
			fmt.Println(1)
		}
		return 1
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return 2
	}

	// 根据选项打印结果
	if opts.value {
		if opts.content != "" {
			fmt.Println(string(body))
		} else if opts.timeoutSet {
			fmt.Printf("%.3f\n", responseSecs)
		} else {
			fmt.Println(resp.StatusCode)
		}
	} else {
		if resp.StatusCode == http.StatusOK {
			if opts.content != "" {
				if strings.Contains(string(body), opts.content) {
					fmt.Println(0)
				} else {
					fmt.Println(1)
				}
			} else {
				fmt.Println(0)
			}
		} else {
			fmt.Println(1)
		}
	}
	return 0
}

// main 是程序的入口点
func main() {
	opts := parseOpts()
	exitCode := getResults(opts)
	os.Exit(exitCode)
}
