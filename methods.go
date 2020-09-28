package go_here

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// HereAPI - структура, методы которой позволяют взаимодействовать с HereMaps REST API
//
type HereAPI struct {
	apiKey string
	lang   string
}

// NewAPI - Создаст структуру для дальнейшего взаимодействия с API
func NewAPI(apiKey string) HereAPI {
	return HereAPI{apiKey: apiKey, lang: "ru-RU"}
}

// SetLanguage - метод для устанвки языка ответа на запросы
//
// На вход передается строка для указания необходимого языка составленная следующим образом:
// пример: "ru-RU", где
// ru -
// RU -
func (api *HereAPI) SetLanguage(lang string) error {
	// TODO: - сделать валидацию возможных языковых кодов
	api.lang = lang
	return nil
}

// Discover - утпрощенный поиск мест.
func (api HereAPI) Discover(position Position, limit int, query string, in string) ([]Item, error) {
	return api.getItems("https://discover.search.hereapi.com/v1/discover", func(v url.Values) {
		v.Add("at", fmt.Sprintf("%v,%v", position.Lat, position.Lng))
		v.Add("limit", fmt.Sprintf("%v", limit))
		v.Add("q", query)
		v.Add("in", in)
	})
}

// Geocode - геокодер
//
// Вернет "просящему" информацию о месте или местах, соответствующих введенному запросу
func (api HereAPI) Geocode(query string) ([]Item, error) {
	return api.getItems("https://geocode.search.hereapi.com/v1/geocode", func(v url.Values) {
		v.Add("q", query)
	})
}

// Autosuggest - автоматическое дополнение к адресу
//
// При вводе части адреса, данный метод выдаст срез элементов, которые мог иметь в виду пользователь.
// Данные адреса будут отсортированы в порядке удаления от пользователя
func (api HereAPI) Autosuggest(position Position, limit int, query string) ([]Item, error) {
	return api.getItems("https://autosuggest.search.hereapi.com/v1/autosuggest", func(v url.Values) {
		v.Add("at", fmt.Sprintf("%v,%v", position.Lat, position.Lng))
		v.Add("limit", fmt.Sprintf("%v", limit))
		v.Add("q", query)
	})
}

// Browse - исследование месности
//
// По локации пользователя будут подобраны места удовлетворяющие выбранным категориям
// Места будут отсортированы в порядке удаления от пользователя
func (api HereAPI) Browse(position Position, limit int) ([]Item, error) {
	return api.getItems("https://browse.search.hereapi.com/v1/browse", func(v url.Values) {
		v.Add("at", fmt.Sprintf("%v,%v", position.Lat, position.Lng))
		v.Add("limit", fmt.Sprintf("%v", limit))
	})
}

// Lookup - просмотр информации об элементе
//
// Данный метод вернет информацию об элементе на карте по его идентификатору в Here Maps
func (api HereAPI) Lookup(id string) (Item, error) {
	return api.getItem("https://lookup.search.hereapi.com/v1/lookup", func(v url.Values) {
		v.Add("id", id)
	})
}

// ReverseGeocode - расшифровка по позиции
//
// Данный метод вернет информацию об объектах в HereMaps, которые расположены в данном месте
func (api HereAPI) ReverseGeocode(position Position) ([]Item, error) {
	return api.getItems("https://revgeocode.search.hereapi.com/v1/revgeocode", func(v url.Values) {
		v.Add("at", fmt.Sprintf("%v,%v", position.Lat, position.Lng))
	})
}

func (api HereAPI) request(url string) ([]byte, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, 500, err
	}

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	return buffer.Bytes(), resp.StatusCode, err
}

func (api HereAPI) unmarshal(data []byte, code int, obj interface{}) error {
	// Обработка ошибок сервиса
	if code != http.StatusOK {
		hereError := HereError{}
		err := json.Unmarshal(data, &hereError)
		if err != nil {
			return err
		}

		return hereError
	}

	err := json.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	return nil
}

func (api HereAPI) category(items []string) string {
	res := ""

	for i, c := range items {
		res += c
		if i != len(items) - 1 {
			res += ","
		}
	}

	return res
}

func (api HereAPI) getItems(baseURL string, prepareURL func(url.Values)) ([]Item, error) {
	resp := Response{}

	data, code, err := api.request(api.generateURL(baseURL, prepareURL))
	if err != nil {
		log.Println(err)
	}

	err = api.unmarshal(data, code, &resp)
	return resp.Items, err
}

func (api HereAPI) getItem(baseURL string, prepareURL func(url.Values)) (Item, error) {
	item := Item{}

	data, code, err := api.request(api.generateURL(baseURL, prepareURL))
	if err != nil {
		return item, err
	}

	err = api.unmarshal(data, code, &item)
	return item, err
}

func (api HereAPI) generateURL(base string, prepare func(values url.Values)) string {
	values := url.Values{}

	prepare(values)

	values.Add("lang", api.lang)
	values.Add("apiKey", api.apiKey)

	return base + "?" + values.Encode()
}
