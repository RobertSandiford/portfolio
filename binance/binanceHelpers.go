package binance

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"strconv"
)

func binanceGet(url string, params []Param, task string) (body []byte) {


	queryString := makeQueryString(params)

	Dump(queryString)

	client := &http.Client{
		/*CheckRedirect: redirectPolicyFunc,*/
	}

	req, err := http.NewRequest("GET", url + queryString, nil)
	if err != nil {
		panic("Error making new request when " + task)
	}

	//fmt.Println("binanceApiKey", binanceApiKey)
	//fmt.Println("binanceApiSecret", binanceApiSecret)

	// set the headers
	req.Header.Add("X-MBX-APIKEY", binanceApiKey)

	res, err := client.Do(req)
	if err != nil {
		panic("Error doing request when " + task)
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error reading body after " + task)
	}

	fmt.Println("binanceGet response", string(body))

	return body

}

func binancePost(url string, params []Param, task string) (body []byte) {

	paramString := makeParameterString(params)

	client := &http.Client{
		/*CheckRedirect: redirectPolicyFunc,*/
	}

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(paramString))
	if err != nil {
		panic("Error making new request when " + task)
	}

	//fmt.Println("binanceApiKey", binanceApiKey)
	//fmt.Println("binanceApiSecret", binanceApiSecret)

	// set the headers
	req.Header.Add("X-MBX-APIKEY", binanceApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(paramString)))

	res, err := client.Do(req)
	if err != nil {
		panic("Error doing request when " + task)
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error reading body after " + task)
	}
	
	fmt.Println("binancePost response", string(body))

	return body

}

func signParams(params []Param, method string) (signedParams []Param) {

	var dataString string
	if method == "GET" {
		//dataString = makeQueryString(params)
		dataString = makeParameterString(params)
	} else if method == "POST" {
		dataString = makeParameterString(params)
	} else {
		fmt.Println("Error: Incorrect METHOD for signParams")
		return []Param{}
	}
	
	fmt.Println("Signing", "(" + method + "):", dataString)

	hmacHash := GetHmac(dataString, binanceApiSecret)
	//fmt.Println("hmacHash", hmacHash)
	params = append(params, Param{Name: "signature", Value: string(hmacHash)})

	return params
}