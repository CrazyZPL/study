package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	APIEndpoint  = "https://api.telegram.org/bot%s/%s"
	FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)

func main() {
	telegramToken := ""
	method := "getUpdates"
	myProxy := ""
	client := resty.New()
	client.SetProxy(myProxy)
	resp, err := client.R().SetHeader("Content-Type", "appplication/json").Post(fmt.Sprintf(APIEndpoint, telegramToken, method))
	if err != nil {
		fmt.Println("this request has error:", err)
		return
	}
	fmt.Print("resp is:", resp)
}
