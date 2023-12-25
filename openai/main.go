package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	baidu     = string("https://www.baidu.com/")
	openaiu   = string("https://api.openai.com/v1")
	chat      = openaiu + "/chat/completions"
	ctJson    = "application/json"
	openaiKey = "123"
)

func main() {
	//get()
	chatCompletions()
	//goole()
	//g6()
}

func goole() {
	proxyUrl, err := url.Parse("http://127.0.0.1:7890/")
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	res, err := client.Get("http://www.google.com/")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}

	println(string(dump))
}

func chatCompletions() {
	proxyUrl, err := url.Parse("http://127.0.0.1:7890/")
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
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
	fmt.Println(chat)
	request, err2 := http.NewRequestWithContext(context.Background(),http.MethodPost, chat, b)
	if err2 != nil {
		fmt.Println("new request")
		fmt.Println(err2)
		return
	}

	request.Header.Set("Content-Type", ctJson)
	request.Header.Set("Authorization", "Bearer "+openaiKey)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("client error")
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}

	println(string(dump))
}

func chatCompletions1() {

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

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("client error")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}

	println(string(dump))

}

func chatCompletions2() {

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
	defer resp.Body.Close()

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

func g6() {
	config := openai.DefaultConfig("you-token")
	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	//设置代理
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	client := openai.NewClientWithConfig(config)
	// client := openai.NewClient("")

	// 向 OpenAI API 发送请求，并获取对话模型的回复
	resp, err := client.CreateChatCompletion(
		context.Background(),

		// 用于发送给 OpenAI API 的请求
		// 并在其中指定了使用的模型和用户发出的消息。
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "今日头条是什么",
				},
			},
		},
	)

	// 发生错误，程序输出错误消息并退出
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	// 程序会打印出 GPT-3.5Turbo 模型回复的内容
	fmt.Println(resp.Choices[0].Message.Content)
}
