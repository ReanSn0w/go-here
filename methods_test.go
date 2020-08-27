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

	items, err := api.Discover(Position{Lat: 42.36346, Lng: -71.05444}, 1, "restaurant", "countryCode:USA")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	log.Println(items[0].Title)
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

	items, err := api.Autosuggest(Position{Lat: 52.93175,Lng: 12.77165}, 5, "res")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	log.Println(items[0].Title)
}

func TestHereAPI_Browse(t *testing.T) {
	api := getAPI()

	items, err := api.Browse(Position{-23.000813,-43.351629}, 2, []string{"100-1100", "200-2000-0011", "100-1000"})
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}

	log.Println(items[0].Title)
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