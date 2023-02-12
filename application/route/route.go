package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID 			string    	 `json:"routeId"`
	ClientID 	string		 `json:"clientId"`
	Positions 	[]Position 	 `json:"position"`
}

type Position struct {
	Latitude float64
	Longitude	float64	
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("route id not informed")
	}
	file, error := os.Open("destinations/" + route.ID + ".txt")
	if error != nil {
		return error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		latitude, error := strconv.ParseFloat(data[1], 64)
		if error != nil {
			return nil
		}
		longitude, error := strconv.ParseFloat(data[0], 64)
		if error != nil {
			return nil
		}
		route.Positions = append(route.Positions, Position{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}
	
	return nil
}

func (route *Route) ExportJsonPositions() ([]string, error) {
	var partialRoute PartialRoutePosition
	var result []string
	total := len(route.Positions)

	for key, value := range route.Positions {
		partialRoute.ID = route.ID
		partialRoute.ClientID = route.ClientID
		partialRoute.Position = []float64{value.Latitude, value.Longitude}
		partialRoute.Finished = false

		if total-1 == key {
			partialRoute.Finished = true
		}

		jsonRoute, error := json.Marshal(route)

		if error != nil {
			return nil, error
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}