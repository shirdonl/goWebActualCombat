package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	CopeHandle("GET","https://www.baidu.com","")
}

//http请求处理
func CopeHandle(method, urlVal,data string)  {
	client := &http.Client{}
	var req *http.Request

	if data == "" {
		urlArr := strings.Split(urlVal,"?")
		if len(urlArr)  == 2 {
			urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	}else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}

	cookie := &http.Cookie{Name: "X-Xsrftoken",Value: "abccadf41ba5fasfasjijalkjaqezgbea3ga", HttpOnly: true}
	req.AddCookie(cookie)

	//添加header
	req.Header.Add("X-Xsrftoken","aaab6d695bbdcd111e8b681002324e63af81")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

//将get请求的参数进行转义
func getParseParam(param string) string  {
	return url.PathEscape(param)
}

