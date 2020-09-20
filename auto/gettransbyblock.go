package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"trontool/api"
	"trontool/model"
	"trontool/tron"
)
type data struct {
	To string `json:"to"`
	From string `json:"from"`
	Block int `json:"block"`
	Value float64 `json:"value"`
}
func main()  {
	for i:=23248826;i < 23338826;i ++   {
		//fmt.Println(" i会死啊快开始：",i)
		getblock(i)
	}
}
func getblock(block int)  {
	redis := model.Redis()
	var re []api.TranscationResultData
	var ss api.TranscationResult
	url := fmt.Sprintf("https://api.trongrid.io/v1/blocks/%d/events",block)
	fmt.Println("url:",url)
	s := Get(url)
	// 将字符串反解析为结构体
	json.Unmarshal([]byte(s), &ss)
	re = ss.Data
	//var datas []data
	for _,res := range re{
		if res.ContractAddress == "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" && res.EventName == "Transfer"{
			for k, v := range res.Result {
				//fmt.Println(k,":::",v)
				var da data
				if k == "from"{
					da.From = v.(string)
				}
				if k == "to" {
					_addr, err := tron.EncodeHexAddress("41" + v.(string)[2:])
					if err != nil {
					}
					da.To = _addr
				}
				if k == "value"{
					da.Value , _ = strconv.ParseFloat(v.(string), 64)
					//fmt.Println("余额：",v)
				}
				is := redis.SIsMember("address",da.To).Val()
				if  is != false{
					da.Block = 123
					//datas = append(datas,da)
					redis.SAdd("cha919",da.To)
					fmt.Println("高度：",block,"地址：",da.To,"金额：",da.Value)
				}


			}

		}

	}
}
// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 50 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
