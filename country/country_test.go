package country

import (
	"testing"
)

const (
	TestAlpha2  = "CA"
	TestAlpha3  = "CAN"
	TestCountry = "Canada"
	TestNumeric = "124"
)

func TestGetCountryNameFromAlpha2(t *testing.T) {
	country, err := GetCountryNameFromAlpha2(TestAlpha2)
	if err != nil {
		t.Fatalf("No Country found for '%s'", TestAlpha2)
	}
	if country != TestCountry {
		t.Fatalf("Country should be '%s' got '%s' for '%s'", TestCountry, country, TestAlpha2)
	}
}

func TestGetCountryNameFromAlpha3(t *testing.T) {
	country, err := GetCountryNameFromAlpha3(TestAlpha3)
	if err != nil {
		t.Fatalf("No Country found for '%s'", TestAlpha3)
	}
	if country != TestCountry {
		t.Fatalf("Country should be '%s' got '%s' for '%s'", TestCountry, country, TestAlpha3)
	}
}

func TestGetCountryNameFromNumeric(t *testing.T) {
	country, err := GetCountryNameFromNumeric(TestNumeric)
	if err != nil {
		t.Fatalf("No Country found for '%s'", TestNumeric)
	}
	if country != TestCountry {
		t.Fatalf("Country should be '%s' got '%s' for '%s'", TestCountry, country, TestNumeric)
	}
}

func TestGetAlpha2FromAlpha3(t *testing.T) {
	alpha2, err := GetAlpha2FromAlpha3(TestAlpha3)
	if err != nil {
		t.Fatalf("No Alpha2 found for '%s'", TestAlpha3)
	}
	if alpha2 != TestAlpha2 {
		t.Fatalf("Alpha2 should be '%s' got '%s' for '%s'", TestAlpha2, alpha2, TestAlpha3)
	}
}

func TestGetAlpha2FromNumeric(t *testing.T) {
	alpha2, err := GetAlpha2FromNumeric(TestNumeric)
	if err != nil {
		t.Fatalf("No Alpha2 found for '%s'", TestNumeric)
	}
	if alpha2 != TestAlpha2 {
		t.Fatalf("Alpha2 should be '%s' got '%s' for '%s'", TestAlpha2, alpha2, TestNumeric)
	}
}

func TestGetAlpha3FromAlpha2(t *testing.T) {
	alpha3, err := GetAlpha3FromAlpha2(TestAlpha2)
	if err != nil {
		t.Fatalf("No Alpha3 found for '%s'", TestAlpha2)
	}
	if alpha3 != TestAlpha3 {
		t.Fatalf("Alpha3 should be '%s' got '%s' for '%s'", TestAlpha3, alpha3, TestAlpha2)
	}
}

func TestGetAlpha3FromNumeric(t *testing.T) {
	alpha3, err := GetAlpha3FromNumeric(TestNumeric)
	if err != nil {
		t.Fatalf("No Alpha3 found for '%s'", TestNumeric)
	}
	if alpha3 != TestAlpha3 {
		t.Fatalf("Alpha3 should be '%s' got '%s' for '%s'", TestAlpha3, alpha3, TestNumeric)
	}
}

func TestGetNumericFromAlpha2(t *testing.T) {
	numeric, err := GetNumericFromAlpha2(TestAlpha2)
	if err != nil {
		t.Fatalf("No numeric found for '%s'", TestAlpha2)
	}
	if numeric != TestNumeric {
		t.Fatalf("Numeric should be '%s' got '%s' for '%s'", TestNumeric, numeric, TestAlpha2)
	}
}

func TestGetNumericFromAlpha3(t *testing.T) {
	numeric, err := GetNumericFromAlpha3(TestAlpha3)
	if err != nil {
		t.Fatalf("No numeric found for '%s'", TestAlpha3)
	}
	if numeric != TestNumeric {
		t.Fatalf("Numeric should be '%s' got '%s' for '%s'", TestNumeric, numeric, TestAlpha3)
	}
}
