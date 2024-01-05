package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type config struct {
	Next     string
	Previous string
}

type Area struct {
	Name string
	Url  string
}

type Generation struct {
	Name string
	Url  string
}

type GameIndice struct {
	GameIndex  int
	Generation Generation
}

type Language struct {
	Name string
	Url  string
}

type PokeLocationName struct {
	Name     string
	Language Language
}

type PokeLocationRegion struct {
	Name string
	Url  string
}

type PokeLocation struct {
	Id          int
	Name        string
	Region      PokeLocationRegion
	Names       []PokeLocationName
	GameIndices []GameIndice
	Areas       []Area
}

func Xd() error {
	fmt.Println("XD XD XD")
	return nil
}

func getJson(r *http.Response, target interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

//func GetLocations(c *config) ([]*string, error) {
//	var locations []*string
//
//	for i := 0; i < 20; i++ {
//		locationObj := new(pokeLocation)
//		res, err := http.Get(c.Next)
//		if err != nil {
//			fmt.Println(err)
//		}
//		err = getJson(res, &locationObj)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		locations = append(locations, locationObj.name)
//
//	}
//	fmt.Println("locations: ")
//	fmt.Println(locations)
//
//	return locations, nil
//}

func GetLocation(locationId int) (*PokeLocation, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location/%v/", locationId)
	fmt.Println(fmt.Sprintf("url: %v", url))

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//fmt.Println(fmt.Sprintf("res: %v", res))

	location := new(PokeLocation)

	defer res.Body.Close()

	fmt.Println(fmt.Sprintf("res.StatusCode: %v", res.StatusCode))
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(fmt.Sprintf("bodyBytes: %v", bodyBytes))
		fmt.Println(fmt.Sprintf("bodyString: %v", bodyString))
		err = json.Unmarshal(bodyBytes, &location)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("location: %v", location))

	}

	return location, nil
}
