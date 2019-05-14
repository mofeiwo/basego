package exlib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"errors"
)

//http封装总结

const (
	USERAGENT = "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36"
)

//GET请求,返回是二进制
func SendGetResponseBytes(client *http.Client, reqUrl string, requestHeaders map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("GET", reqUrl, nil)

	if requestHeaders != nil {
		for k, v := range requestHeaders {
			req.Header.Add(k, v)
		}
	} else {
		req.Header.Add("Content-type", "application/x-www-form-urlencoded")
		req.Header.Add("User-Agent",USERAGENT)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HttpStatusCode:%d ,Desc:%s", resp.StatusCode, string(bodyData)))
	}

	return bodyData, nil
}
