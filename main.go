package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	cityCode := "130010"

	v := url.Values{}
	v.Set("city", cityCode)
	resp, err := http.Get("http://weather.livedoor.com/forecast/webservice/json/v1" + "?" + v.Encode())

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	jsonBytes := ([]byte)(byteArray)
}
