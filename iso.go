package iso3166

import (
	"errors"
	"strings"
)

const (
	alpha2CodeLength = 2
	alpha3CodeLength = 3
)

var (
	ErrCountryNotFound = errors.New("iso3166: country not found")
)

type Country struct {
	NumericCode int
	Name        string
	Alpha2Code  string
	Alpha3Code  string
}

func Alpha2CodeExists(code string) bool {
	if len(code) != alpha2CodeLength {
		return false
	}

	code = strings.ToUpper(code)

	for _, c := range countries {
		if c.Alpha2Code == code {
			return true
		}
	}

	return false
}

func Alpha3CodeExists(code string) bool {
	if len(code) != alpha3CodeLength {
		return false
	}

	code = strings.ToUpper(code)

	for _, c := range countries {
		if c.Alpha3Code == code {
			return true
		}
	}

	return false
}

func NumericCodeExists(code int) bool {
	for _, c := range countries {
		if c.NumericCode == code {
			return true
		}
	}

	return false
}

func CountryByAlphaCode(code string) (*Country, error) {
	code = strings.ToUpper(code)
	l := len(code)

	if l < alpha2CodeLength || l > alpha3CodeLength {
		return nil, ErrCountryNotFound
	}

	for _, c := range countries {
		if c.Alpha3Code == code || c.Alpha2Code == code {
			return c, nil
		}
	}

	return nil, ErrCountryNotFound
}

func CountryByNumericCode(code int) (*Country, error) {
	if code < 0 {
		return nil, ErrCountryNotFound
	}

	for _, c := range countries {
		if c.NumericCode == code {
			return c, nil
		}
	}

	return nil, ErrCountryNotFound
}
