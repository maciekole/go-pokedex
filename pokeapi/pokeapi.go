package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func getLocationById(locationId int) (*PokeLocation, error) {
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

func getLocationByUrl(locationUrl *string) (*PokeLocation, error) {
	var url string

	if locationUrl == nil {
		url = "https://pokeapi.co/api/v2/location/1/"
	} else {
		url = *locationUrl
	}

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
		//bodyString := string(bodyBytes)
		//fmt.Println(fmt.Sprintf("bodyBytes: %v", bodyBytes))
		//fmt.Println(fmt.Sprintf("bodyString: %v", bodyString))
		err = json.Unmarshal(bodyBytes, &location)
		if err != nil {
			fmt.Println(err)
		}
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("no more locations to explore")
	}
	return location, nil
}

func GetLocationsForward(startingLocationUrl *string) ([]string, *string, *string, error) {
	var locations []string
	var nextStartingLocationBack *string

	for i := 0; i < 20; i++ {
		location, err := getLocationByUrl(startingLocationUrl)
		if err != nil {
			return nil, nil, nil, err
		}
		fmt.Println(fmt.Sprintf("location: %v", *location))
		nextLocation := fmt.Sprintf("https://pokeapi.co/api/v2/location/%v/", location.Id+1)

		nextStartingLocationBack = startingLocationUrl
		startingLocationUrl = &nextLocation
		locations = append(locations, location.Name)
	}

	return locations, startingLocationUrl, nextStartingLocationBack, nil
}

func GetLocationsBackward(startingLocationUrl *string) ([]string, *string, *string, error) {
	var locations []string
	var nextStartingLocation *string

	for i := 0; i < 20; i++ {
		location, err := getLocationByUrl(startingLocationUrl)
		if err != nil {
			return nil, nil, nil, err
		}
		fmt.Println(fmt.Sprintf("location: %v", *location))
		nextLocation := fmt.Sprintf("https://pokeapi.co/api/v2/location/%v/", location.Id-1)

		nextStartingLocation = startingLocationUrl
		startingLocationUrl = &nextLocation
		locations = append(locations, location.Name)
	}

	return locations, nextStartingLocation, startingLocationUrl, nil
}
