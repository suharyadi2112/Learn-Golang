package main

import (
	"log"
	"fmt"
	"encoding/json"
)

func main(){
	// Define a struct type with tags
	type DukcapilVerificationVM struct {
	    FamilyCardNumber        string  `json:"family_card_number"`
	    Nik                     string  `json:"nik"`
	    NikPercentage           float64 `json:"nik_percentage"`
	    Name                    string  `json:"name"`
	    NamePercentage          float64 `json:"name_percentage"`
	    StatusInFamily          string  `json:"status_in_family"`
	    CityName                string  `json:"city_name"`
	    DistrictName            string  `json:"district_name"`
	    Occupation              string  `json:"occupation"`
	    Rw                      string  `json:"rw"`
	    Rt                      string  `json:"rt"`
	    Address                 string  `json:"address"`
	    AddressPercentage       float64 `json:"address_percentage"`
	    HighestEducationLevel   string  `json:"highest_education_level"`
	    BirthPlace              string  `json:"birth_place"`
	    BirthPlacePercentage    float64 `json:"birth_place_percentage"`
	    MaritalStatus           string  `json:"marital_status"`
	    MaritalStatusPercentage float64 `json:"marital_status_percentage"`
	    MotherName              string  `json:"mother_name"`
	    MotherNamePercentage    float64 `json:"mother_name_percentage"`
	    ProvinceName            string  `json:"province_name"`
	    SubDistrictName         string  `json:"sub_district_name"`
	    Gender                  string  `json:"gender"`
	    GenderPercentage        float64 `json:"gender_percentage"`
	    BirthDate               string  `json:"birth_date"`
	    BirthDatePercentage     float64 `json:"birth_date_percentage"`
	    BaName                  string  `json:"ba_name"`
	    BaNamePercentage        float64 `json:"ba_name_percentage"`
	}

	// Decode the existing JSON data into a Person value
	var p DukcapilVerificationVM
	err := json.Unmarshal([]byte(`{"family_card_number":"Tidak Sesuai","name":"Sesuai","name_percentage":100,"status_in_family":"Tidak Sesuai","city_name":"Sesuai","district_name":"Sesuai","occupation":"Tidak Sesuai","rw":"Sesuai","rt":"Sesuai","address":"Sesuai","address_percentage":100,"highest_education_level":"Tidak Sesuai","birth_place":"Sesuai","birth_place_percentage":100,"marital_status":"Tidak Sesuai","mother_name":"Sesuai","mother_name_percentage":100,"province_name":"Sesuai","sub_district_name":"Sesuai","gender":"Sesuai","birth_date":"Sesuai", "ba_name" : "suharyadi", "ba_name_percentage": 100}`), &p)


	if err != nil {
		log.Fatal(err)
	}

	p.Name = "Tidak Sesuai"

	// Encode the updated Person value into a new JSON data
	b, err := json.Marshal(p)
	if err != nil {
	log.Fatal(err)
	}
	fmt.Println(string(b))
}