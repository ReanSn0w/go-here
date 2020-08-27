package go_here

import "fmt"

// HereError - структура описывает формат ошибок, который возвращает API
// - реальзует стандартный интерфейс Error
type HereError struct {
	Title string `json:"error"`
	Description string `json:"error_description"`
}

// Error метод для реализации стандартного интерфейса ошибки golang
func (error HereError) Error() string {
	return fmt.Sprintf("%v: %v", error.Title, error.Description)
}

type Response struct {
	Items []Item `json:"items"`
}

// Item - Структура описывает элемент на карте
type Item struct {
	Title           string        `json:"title"`
	ID              string        `json:"id"`
	OntologyID      string        `json:"ontologyId,omitempty"`
	ResultType      string        `json:"resultType"`
	HouseNumberType string        `json:"houseNumberType,omitemty"`
	Address         Address       `json:"address"`
	Position        Position      `json:"position"`
	Access          []Position    `json:"access"`
	Distance        int64         `json:"distance,omitempty"`
	MapView         MapView       `json:"mapView,omitempty"`
	Categories      []Category    `json:"categories,omitempty"`
	Contacts        []Contact     `json:"contacts,omitempty"`
	References      []Reference   `json:"references,omitempty"`
	FoodTypes       []FoodType    `json:"foodTypes,omitempty"`
	OpeningHours    []OpeningHour `json:"openingHours,omitempty"`
}

// Position - структура описывает местоположение объекта
type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Address - структура описывает детальное расшифровку адреса
type Address struct {
	Label       string `json:"label"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	State       string `json:"state"`
	County      string `json:"county"`
	City        string `json:"city"`
	District    string `json:"district"`
	Street      string `json:"street"`
	PostalCode  string `json:"postalCode"`
	HouseNumber string `json:"houseNumber"`
}

// MapView - положение элемента на карте
type MapView struct {
	West  float64 `json:"west"`
	South float64 `json:"south"`
	East  float64 `json:"east"`
	North float64 `json:"north"`
}

type Contact struct {
	Phone []ContactItem `json:"phone"`
	WWW   []ContactItem `json:"www"`
	Email []ContactItem `json:"email"`
}

type ContactItem struct {
	Value string `json:"value"`
}

type Category struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type FoodType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Primary *bool  `json:"primary,omitempty"`
}

type OpeningHour struct {
	Text       []string     `json:"text"`
	IsOpen     bool         `json:"isOpen"`
	Structured []Structured `json:"structured"`
}

type Structured struct {
	Start      string `json:"start"`
	Duration   string `json:"duration"`
	Recurrence string `json:"recurrence"`
}

type Reference struct {
	Supplier Supplier `json:"supplier"`
	ID       string   `json:"id"`
}

type Supplier struct {
	ID string `json:"id"`
}