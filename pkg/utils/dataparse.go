package utils

import (
	"encoding/json"
	"fmt"
)

type Country struct {
	Name      string
	Capital   string
	Continent string
}

func CountryData() {

	var country Country

	Data := []byte(`{
	"Name":"India",
	"Capital":"Delhi",
	"Continent":"Asia"
	}`)

	err := json.Unmarshal(Data, &country)

	if err != nil {
		fmt.Print("error", err)
	}

	fmt.Print("Structs is : ", country, "\n")
	fmt.Printf("%s's capital is %s and it is in %s\n", country.Name, country.Capital, country.Continent)
}
