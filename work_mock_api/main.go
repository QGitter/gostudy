package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const KEY = "dev!@#$%^168"

var client = &http.Client{
	Timeout: 10 * time.Second,
}
var (
	request  *http.Request
	err      error
	response *http.Response
	data     []byte
	parse    *url.URL
)

func main() {

}

func Post(url string, params map[string]string) string {
	request, err = http.NewRequest(http.MethodPost, url, strings.NewReader(GetSortParams(params)))
	if err != nil {
		fmt.Println("request is error")
		return ""
	}
	CommHeader(request, params)
	response, err = client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("response is error")
		return ""
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("response is error")
		return ""
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("body is error")
		return ""
	}
	return string(data)
}

func Get(uri string) string {
	params := make(map[string]string)
	parse, err = url.Parse(uri)
	if err != nil {
		fmt.Println("parse is error")
		return ""
	}
	query := parse.RawQuery
	if query != "" {
		split := strings.Split(query, "&")
		for _, str := range split {
			child := strings.Split(str, "=")
			params[child[0]] = child[1]
		}
	}
	request, err = http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		fmt.Println("get is error")
		return ""
	}
	CommHeader(request, params)
	response, err = client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("response is error")
		return ""
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("response is error")
		return ""
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("body is error")
		return ""
	}
	return string(data)
}

func CommHeader(req *http.Request, params map[string]string) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")
	req.Header.Set("Game-Meios-Version-Name", "4.1.2")
	req.Header.Set("Game-Android-Version", "23")
	req.Header.Set("Game-Client-Id", "1189857322")
	req.Header.Set("Game-Android-Id", "9cceee1466ced851")
	req.Header.Set("Game-Model", "MP1503")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; MP1503 Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/44.0.2403.119 Mobile Safari/537.36")
	req.Header.Set("Game-Version", "1.4.0")
	req.Header.Set("Game-Access-Token", "_v2NWQxMmYzMDYjMTUzNjEyMTAxMSMwIzYxIzYxY2MyM2Y0YWI1M2NjNDM3Yzk5YzM5MzRlMjQ4YTFjODkjI0JKX1NIIzViMThiMWIz")
	req.Header.Set("client_version_code", "1004000")
	req.Header.Set("Game-Imei", "860273030007340")
	req.Header.Set("Game-Sign", GetSign(params))
	req.Header.Set("Game-Network", "5")
	req.Header.Set("Game-Language", "zh-Hans")
}

func GetSign(params map[string]string) string {
	return GetBase64Encode(GetSha1(GetSortParams(params), KEY))
}

func GetSortParams(params map[string]string) string {
	var paramsString string
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, value := range keys {
		paramsString = paramsString + value + "=" + params[value] + "&"
	}
	paramsString = paramsString[0 : len(paramsString)-1]
	return paramsString
}

func GetSha1(src, key string) string {
	m := hmac.New(sha1.New, []byte(key))
	m.Write([]byte(src))
	return string(m.Sum(nil))
}

func GetBase64Encode(str string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(str)))
}
