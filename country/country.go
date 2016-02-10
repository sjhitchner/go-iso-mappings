// Package to help map ISO Country Codes
//
// ISO 3166-1
// https://en.wikipedia.org/wiki/ISO_3166-1
//
// Defines standardized 2-digit, 3-digit, and numeric codes this package
// provides a way to easily convert between them.
//
// All country data is contained in tab-delimited countries.txt.
// To update with new country information edit this file with
// the new information
//
// To compile this package you need to run:
//  go generate
//
// This will generate all the necessary go maps that hold the mapping Data
//
package iso

import (
	"fmt"
)

//go:generate csv2map -d "\t" -in country.txt -k 2 -v 1 -m alpha2ToCountry -p iso -h
// Map between Country Name and 2-digit alpha code
func GetCountryNameFromAlpha2(alpha2 string) (string, error) {
	country, ok := alpha3ToCountry[alpha2]
	if !ok {
		return "", fmt.Errorf("No country found for alpha3 '%s'", alpha2)
	}
	return country, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 3 -v 1 -m alpha3ToCountry -p iso -h
// Map between Country Name and 3-digit alpha code
func GetCountryNameFromAlpha3(alpha3 string) (string, error) {
	country, ok := alpha3ToCountry[alpha3]
	if !ok {
		return "", fmt.Errorf("No country found for alpha3 '%s'", alpha3)
	}
	return country, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 4 -v 1 -m numericToCountry -p iso -h
// Map between Country Name and 3-digit alpha code
func GetCountryNameFromNumeric(numeric string) (string, error) {
	country, ok := numericToCountry[numeric]
	if !ok {
		return "", fmt.Errorf("No country found for numeric '%s'", numeric)
	}
	return country, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 3 -v 2 -m alpha3ToAlpha2 -p iso -h
// Map between 3-digit alpha code and 2-digit alpha code
func GetAlpha2FromAlpha3(alpha3 string) (string, error) {
	alpha2, ok := alpha3ToAlpha2[alpha3]
	if !ok {
		return "", fmt.Errorf("No alpha2 code found for alpha2 '%s'", alpha3)
	}
	return alpha2, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 4 -v 2 -m numericToAlpha2 -p iso -h
// Map between 3-digit alpha code and 2-digit alpha code
func GetAlpha2FromNumeric(numeric string) (string, error) {
	alpha2, ok := numericToAlpha2[numeric]
	if !ok {
		return "", fmt.Errorf("No alpha2 code found for numeric '%s'", numeric)
	}
	return alpha2, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 2 -v 3 -m alpha2ToAlpha3 -p iso -h
// Map between 3-digit alpha code and 2-digit alpha code
func GetAlpha3FromAlpha2(alpha2 string) (string, error) {
	alpha3, ok := alpha2ToAlpha3[alpha2]
	if !ok {
		return "", fmt.Errorf("No alpha3 code found for alpha2 '%s'", alpha2)
	}
	return alpha3, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 4 -v 3 -m numericToAlpha3 -p iso -h
// Map between 3-digit alpha code and 2-digit alpha code
func GetAlpha3FromNumeric(numeric string) (string, error) {
	alpha3, ok := numericToAlpha3[numeric]
	if !ok {
		return "", fmt.Errorf("No alpha3 code found for numeric '%s'", numeric)
	}
	return alpha3, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 2 -v 4 -m alpha2ToNumeric -p iso -h
// Map between 2-digit alpha code and numeric code
func GetNumericFromAlpha2(alpha2 string) (string, error) {
	numeric, ok := alpha2ToNumeric[alpha2]
	if !ok {
		return "", fmt.Errorf("No numeric code found for alpha2 '%s'", alpha2)
	}
	return numeric, nil
}

//go:generate csv2map -d "\t" -in country.txt -k 3 -v 4 -m alpha3ToNumeric -p iso -h
// Map between 2-digit alpha code and numeric code
func GetNumericFromAlpha3(alpha3 string) (string, error) {
	numeric, ok := alpha3ToNumeric[alpha3]
	if !ok {
		return "", fmt.Errorf("No numeric code found for alpha3 '%s'", alpha3)
	}
	return numeric, nil
}
