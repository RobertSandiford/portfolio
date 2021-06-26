package binance

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func makeQueryString(params []Param) (queryString string) {
	queryString = "?" + makeParameterString(params)
	return queryString
}

func makeParameterString(params []Param) (parameterString string) {
	for i, param := range params {
		parameterString += param.Name + "=" + param.Value
		if i+1 < len(params) {
			parameterString += "&"
		}
	}
	return parameterString
}

func httpGet(url, queryString, task string) (body []byte) {
	res, err := http.Get(url + queryString)
	if err != nil {
		panic("Error " + task)
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error reading body after " + task)
	}
	fmt.Println( string(body) )

	return body
}
