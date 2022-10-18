package main

import (
	"encoding/json"
	"fmt"
	http2 "gobase/net/http"
)

func main() {
	key := ""
	ip := "59.174.98.189"
	//ip := "2408:400a:11d:9d10::1003"
	ans := GaoDeIp2Attribution(ip,key)
	fmt.Println(ans)
}

type IpAttribution struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
}
func GaoDeIp2Attribution(ip ,key string) IpAttribution {
	resp := IpAttribution{}
	url := "https://restapi.amap.com/v5/ip?output=json&type=4&ip="+ip+"&key="+key
	data,err := http2.HttpGet(url)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &resp)
	if resp.Status != "1" {
		url = "https://restapi.amap.com/v5/ip?output=json&type=6&ip="+ip+"&key="+key
		data,err = http2.HttpGet(url)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(data, &resp)
	}
	return resp
}