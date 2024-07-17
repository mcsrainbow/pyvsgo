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

type Options struct {
	url        string
	timeout    float64
	content    string
	auth       string
	value      bool
	payload    string
	timeoutSet bool
}

// parseOpts parses the command line options and returns an Options struct.
func parseOpts() Options {
	var opts Options

	flag.StringVar(&opts.url, "u", "", "URL to GET or POST (default http://)")
	flag.Float64Var(&opts.timeout, "t", 10, "seconds before connection times out")
	flag.StringVar(&opts.content, "c", "", "string to expect in the content")
	flag.StringVar(&opts.auth, "a", "", "username:password on sites with basic authentication")
	flag.BoolVar(&opts.value, "V", false, "return actual value instead of 0 and 1")
	flag.StringVar(&opts.payload, "p", "", "URL encoded http POST data")

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
	exampleText := strings.ReplaceAll(exampleTemplate, "{prog}", os.Args[0])

	// Remove leading tabs
	trimmedExampleText := trimLeadingTabs(exampleText)

	flag.Usage = func() {
		fmt.Printf("usage: %s [-h] -u url [-t timeout] [-c content] [-a auth] [-V] [-p payload]\n", os.Args[0])
		fmt.Print(trimmedExampleText)
		fmt.Println("\noptions:")
		flag.PrintDefaults()
	}

	// Check if no arguments are provided
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	flag.Parse()

	// Check if timeout was set by the user
	timeoutSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "t" {
			timeoutSet = true
		}
	})

	if opts.url == "" {
		flag.Usage()
		os.Exit(2)
	}

	if !strings.HasPrefix(opts.url, "http://") && !strings.HasPrefix(opts.url, "https://") {
		opts.url = "http://" + opts.url
	}

	opts.timeoutSet = timeoutSet
	return opts
}

// trimLeadingTabs removes leading tab characters from each line of the input string.
func trimLeadingTabs(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, "\t")
	}
	return strings.Join(lines, "\n")
}

// getResults performs the HTTP request based on the provided options.
func getResults(opts Options) int {
	client := &http.Client{
		Timeout: time.Duration(opts.timeout * float64(time.Second)),
	}

	var req *http.Request
	var err error

	if opts.payload != "" {
		req, err = http.NewRequest("POST", opts.url, strings.NewReader(opts.payload))
	} else {
		req, err = http.NewRequest("GET", opts.url, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return 2
	}

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
	resp, err := client.Do(req)
	responseSecs := time.Since(startTime).Seconds()

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

// main is the entry point of the program.
func main() {
	opts := parseOpts()
	exitCode := getResults(opts)
	os.Exit(exitCode)
}

