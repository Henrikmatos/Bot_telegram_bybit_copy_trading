package get

import (
	// "bybit/bybit/bybit"
	"bybit/bybit/print"
	"bybit/bybit/sign"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequetJson(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panic(err)
	}
	return body, err
}

func GetPrice(symbol string) Price {
	var curr Price
	url := fmt.Sprint("https://api-testnet.bybit.com/v2/public/tickers?symbol=", symbol)
	body, err := GetRequetJson(url)
	if err != nil {
		log.Panic(err)
	}
	jsonErr := json.Unmarshal(body, &curr)
	if jsonErr != nil {
		log.Panic(jsonErr)
	}
	return curr
}

func GetWallet(api string, api_secret string) Wallet {
	var wall Wallet
	params := map[string]string{
		"api_key":   api,
		"coin":      "USDT",
		"timestamp": print.GetTimestamp(),
	}

	signature := sign.GetSigned(params, api_secret)
	url := fmt.Sprint(
		"https://api-testnet.bybit.com/v2/private/wallet/balance?api_key=",
		api,
		"&coin=USDT",
		"&timestamp=",
		params["timestamp"],
		"&sign=",
		signature,
	)
	body, err := GetRequetJson(url)
	if err != nil {
		log.Panic(err)
	}
	jsonErr := json.Unmarshal(body, &wall)
	if jsonErr != nil {
		log.Panic(jsonErr)
	}
	return (wall)
}
