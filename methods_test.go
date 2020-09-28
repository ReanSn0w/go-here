package go_here

import (
	"log"
	"os"
	"testing"
)

func getAPI() HereAPI {
	key := os.Getenv("Key")
	return NewAPI(key)
}

func TestHereAPI_Discover(t *testing.T) {
	api := getAPI()

	items, err := api.Discover(Position{Lat: 55.829169, Lng: 37.493320}, 10, "restaurant", "countryCode:USA")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	for _, item := range items {
		log.Println(item.Title)
	}
}

func TestHereAPI_Geocode(t *testing.T) {
	api := getAPI()

	items, err := api.Geocode("5 Rue Daunou, 75002 Paris, France")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	log.Println(items[0].Title)
}

func TestHereAPI_Autosuggest(t *testing.T) {
	api := getAPI()

	items, err := api.Autosuggest(Position{55.665781, 37.496703}, 20, "адм")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	for _, item := range items {
		log.Println(item.Title, item.Address.Label, item.Position)
	}
}

func TestHereAPI_Browse(t *testing.T) {
	api := getAPI()

	items, err := api.Browse(Position{55.665781, 37.496703}, 10)
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	for _, item := range items {
		log.Println(item.Title)
	}
}

func TestHereAPI_Lookup(t *testing.T) {
	api := getAPI()

	item, err := api.Lookup("here:pds:place:276u0vhj-b0bace6448ae4b0fbc1d5e323998a7d2")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	log.Println(item.Title)
}

func TestHereAPI_ReverseGeocode(t *testing.T) {
	api := getAPI()

	items, err := api.ReverseGeocode(Position{-23.000813,-43.351629})
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	log.Println(items[0].Title)
}