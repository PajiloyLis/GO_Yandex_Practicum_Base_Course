package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Рeader struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		Id         int    `json:"id"`
		Attributes struct {
			Email       string `json:"email"`
			Article_ids []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	var r Response
	err := json.Unmarshal([]byte(rawResp), &r)
	return r, err
}

func main() {
	raw := `{
		"header": {
			"code": 0,
			"message": ""
		},
		"data": [{
			"type": "user",
			"id": 100,
			"attributes": {
				"email": "bob@yandex.ru",
				"article_ids": [10, 11, 12]
			}
		}]
	}`
	result, err := ReadResponse(raw)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
