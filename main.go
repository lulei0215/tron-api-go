package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	ValidateAddress()

}

type Accounts struct {
	Succuss bool	`json:"success,omitempty"`
	Meta map[string]interface{}	`json:"meta,omitempty"`
	Data []map[string]interface{} `json:"data,omitempty"`
	StatusCode int	`json:statuscode",omitempty"`
	Error string	`json:error",omitempty"`
	//Success bool `json:"success"`
	//Meta map[string]int `json:"meta"`
	//Data interface{} `json:"data"`

}
//{"success":false,"statusCode":400,"error":"A valid account address is required."}

//获取账户信息
func account()  {

	resp, err := Get("https://api.trongrid.io/v1/accounts/TUD4YXYdj2t1gP5th3A7t97mx1AUmrrQRt",nil,nil)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	var mapResult Accounts
	errs := json.Unmarshal(body, &mapResult)
	fmt.Println(mapResult)
	if errs != nil {
		log.Fatalln("JsonToMapDemo err: ", errs)
	}
	//fmt.Println(mapResult.Error)
	if mapResult.Error == "" {
		fmt.Println("errorshikong")
		log.Fatalln(mapResult.Error)
	}
}

//获取账户历史交易信息
func transactions()  {
	url :=  "TUD4YXYdj2t1gP5th3A7t97mx1AUmrrQRt"
	parems := make(map[string]string)
	parems = map[string]string{"limit":"10"}
	resp, err := Get("https://api.trongrid.io/v1/accounts/"+url+"/transactions",parems,nil)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//if err != nil {
	//	fmt.Println(err)
	//}
	var mapResult Accounts
	errs := json.Unmarshal(body, &mapResult)
	fmt.Println(mapResult)
	if errs != nil {
		log.Fatalln("JsonToMapDemo err: ", errs)
	}
	//fmt.Println(mapResult.Error)
	if mapResult.Error == "" {
		fmt.Println("errorshikong")
		log.Fatalln(mapResult.Error)
	}
	//fmt.Println() Println
}
type Genera struct {
	PrivateKey string `json:"privatekey"`
	Address string `json:"address"`
	HexAddress string `json:"hexaddress"`
}
//生成随机私钥和相应的账户地址.
func GenerateAddress()  {
	url :=  "https://api.trongrid.io/wallet/generateaddress"

	resp, err := Get(url,nil,nil)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	//var generateaddress map[string]string

	var df Genera
	json.Unmarshal(body,&df)
	//fmt.Println(errs)
	fmt.Println(df)
}
type Validate struct {
	Result bool `json:"result"`
	Message string `json:"message"`
}
//检查地址是否格式正确
func ValidateAddress()  {
	url := "https://api.trongrid.io/wallet/validateaddress"
	pare := map[string]string{
			"address":"TJbmbC8HQBoWFdkPBRTH2KqgpAbmk5cfUb",
	}

	resp, err := Post(url,pare,nil,nil)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var mapResult Validate
	errs := json.Unmarshal(body, &mapResult)
	if errs != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	fmt.Println(mapResult)
}
//curl --request GET \
//--url https://api.trongrid.io/v1/accounts/TUD4YXYdj2t1gP5th3A7t97mx1AUmrrQRt/transactions

//Get http get method
func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go GET URL : %s \n", req.URL.String())
	return client.Do(req)

}

//Post http post method
func Post(url string, body map[string]string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())
	return client.Do(req)
}