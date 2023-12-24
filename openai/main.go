package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	baidu     = string("https://www.baidu.com/")
	openai    = string("https://api.openai.com")
	chat      = openai + "/v1/chat/completions"
	ctJson    = "application/json"
	openaiKey = ""
)

func main() {
	//get()
	chatCompletions()
}

func chatCompletions() {

	//proxyURL, _ := url.Parse("http://127.0.0.1:7890")
	// 创建一个HTTP客户端
	client := &http.Client{
		// 设置Transport字段为自定义Transport，其中包含代理信息
		Transport: &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",           // 代理服务器协议
				Host:   "127.0.0.1:7890", // 代理服务器地址
			}),
		},
	}
	json := `{
		"model": "gpt-3.5-turbo",
		"messages": [{
		"role": "system",
		"content": "You are a helpful assistant."
		},{
		"role": "user",
		"content": "Who won the world series in 2020?"
		}
		]
		}`
	b := strings.NewReader(json)
	request, err2 := http.NewRequest("post", chat, b)
	if err2 != nil {
		fmt.Println("new request")
		fmt.Println(err2)
		return
	}

	request.Header.Set("Content-Type", ctJson)
	request.Header.Set("Authorization", "Bearer "+openaiKey)

	// 打印请求头部信息
	log.Println("Request Headers:")
	for key, values := range request.Header {
		for _, value := range values {
			log.Printf("%s: %s\n", key, value)
		}
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("client error")
		fmt.Println(err)
		return
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read all")
		fmt.Println(err)
		return
	}
	fmt.Println("ccc")
	fmt.Println(string(all))

}

/**

curl https://api.openai.com/v1/chat/completions \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $OPENAI_API_KEY" \
-d '{
"model": "gpt-3.5-turbo",
"messages": [
{
"role": "system",
"content": "You are a helpful assistant."
},
{
"role": "user",
"content": "Who won the world series in 2020?"
},
{
"role": "assistant",
"content": "The Los Angeles Dodgers won the World Series in 2020."
},
{
"role": "user",
"content": "Where was it played?"
}
]
}'
*/

func get() {
	//k := "chatgpt"
	resp, err := http.Get(baidu)
	if err != nil {
		fmt.Println("failed")
		return
	}
	b := resp.Body
	//b.Read(bb)
	bb, err := ioutil.ReadAll(b)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(bb))
}
