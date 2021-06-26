package binance

import (
	"fmt"
	"github.com/shopspring/decimal"
	//"hash"
	"crypto/sha256"
	"crypto/hmac"
	"time"
	"encoding/hex"
	//"encoding/base64"
	"encoding/json"
	"strconv"
	"github.com/labstack/echo"
)

func CalculatePool() decimal.Decimal {
	return decimal.NewFromFloat32(0)
}

func GetHmac(message, key string) (hmacHash string) {

	//fmt.Println("hash m ", message )
	//fmt.Println("hash k ", key )

	hmacHasher := hmac.New( sha256.New, []byte(key) )
	hmacHasher.Write( []byte(message) )

	h := hmacHasher.Sum(nil)
	//fmt.Println("hash ", fmt.Sprintf("%x", h ) )

	return hex.EncodeToString( h )
	//return base64.StdEncoding.EncodeToString( hmacHasher.Sum(nil) )
}
/*
func GetHmac(message string, key string) (hmacHash []byte) {
	hmacHasher := hmac.New(sha256.New, []byte(key))
	hmacHasher.Write([]byte(message))
	return hmacHasher.Sum(nil)
}*/

func GetMiliTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
func MiliTimestampString() (timestamp string) {
	return strconv.FormatInt(GetMiliTimestamp(), 10)
}

func JsonEncode(data interface{}) (jsonString string, err error) {
	jsonBytes, err := json.Marshal(data)
	return string(jsonBytes), err
} 
func JsonEncodeBytes(data interface{}) (jsonBytes []byte, err error) {
	return json.Marshal(data)
} 

func JsonDecode(jsonString string, structure interface{}) (err error) {
	return json.Unmarshal( []byte(jsonString), structure )
}
func JsonDecodeBytes(jsonBytes []byte, structure interface{}) (err error) {
	return json.Unmarshal( jsonBytes, structure )
}

func jsonDecodePostBody(c echo.Context, data interface{}) error {
	return c.Bind(data)
}

func Dump(data interface{}) {
	fmt.Printf("%+v\n", data)
}