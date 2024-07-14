package wb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"wb-bot/logger"
)

func GetProductIDs(query string) ([]int, error) {
	client := &http.Client{}
	searchUrl := fmt.Sprintf("https://search.wb.ru/exactmatch/ru/common/v5/search?ab_testing=false&appType=1&dest=-1257786&query=%s&resultset=catalog&sort=popular&spp=30&suppressSpellcheck=false", url.QueryEscape(query))
	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,ru-RU;q=0.8,ru;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "https://www.wildberries.ru")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log.Error(err.Error())
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	var ids []int
	if len(response.Data.Products) > 0 {
		ids = make([]int, len(response.Data.Products))
		for i, product := range response.Data.Products {
			ids[i] = product.Id
		}
	}
	return ids, nil
}

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	Products []Product `json:"products"`
}

type Response struct {
	Data Data `json:"data"`
}
