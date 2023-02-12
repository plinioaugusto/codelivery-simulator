package main

import (
	"fmt"

	route "github.com/plinioaugusto/simulator-full-cycle-immersion-2023/application/route"
)



func main() {
	route := route.Route{
		ID: 			"1",
		ClientID: 		"1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}