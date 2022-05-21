package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/byyam/DailyPrice"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	coinMarketKey = ""
	pushKey       = ""
)

func main() {
	var cm DailyPrice.CoinMarket
	if err := GetPrice(&cm); err != nil {
		log.Fatalln("get price error:", err)
	}
	if err := PushMsg(cm); err != nil {
		log.Fatalln("push message error:", err)
	}
}

func GetPrice(cm *DailyPrice.CoinMarket) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		return err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "2")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", coinMarketKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	log.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(respBody))

	_ = json.Unmarshal(respBody, cm)
	for _, d := range cm.Data {
		log.Println(d.Quote.USD.Price)
	}
	return nil
}

func PushMsg(cm DailyPrice.CoinMarket) error {
	client := &http.Client{}
	u := url.URL{Scheme: "https", Host: "qyapi.weixin.qq.com", Path: "/cgi-bin/webhook/send"}
	q := u.Query()
	q.Add("key", pushKey)
	u.RawQuery = q.Encode()
	log.Println("push url:", u.String())

	msg := DailyPrice.PushMsg{
		MsgType: "text",
		Text: DailyPrice.PushMsgText{
			Content: fmt.Sprintf("name: %s, price:%f", cm.Data[0].Name, cm.Data[0].Quote.USD.Price),
		},
	}
	data, err := json.Marshal(&msg)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	log.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(respBody))

	return nil
}
