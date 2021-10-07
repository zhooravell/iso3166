package iso3166

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlpha2CodeExists_True(t *testing.T) {
	for _, test := range countries {
		t.Run(test.Alpha2Code, func(t *testing.T) {
			assert.True(t, Alpha2CodeExists(test.Alpha2Code))
		})
	}
}

func TestAlpha2CodeExists_False(t *testing.T) {
	tests := []string{
		"AA",
		"22",
		"ATA",
		"PYF",
		"404",
		"Zambia",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			assert.False(t, Alpha2CodeExists(test))
		})
	}
}

func TestAlpha3CodeExists_True(t *testing.T) {
	for _, test := range countries {
		t.Run(test.Alpha3Code, func(t *testing.T) {
			assert.True(t, Alpha3CodeExists(test.Alpha3Code))
		})
	}
}

func TestAlpha3CodeExists_False(t *testing.T) {
	tests := []string{
		"AAA",
		"22",
		"AT",
		"PY",
		"404",
		"Zambia",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			assert.False(t, Alpha3CodeExists(test))
		})
	}
}

func TestNumericCodeExists_True(t *testing.T) {
	for _, test := range countries {
		t.Run(fmt.Sprintf("Num:%d", test.NumericCode), func(t *testing.T) {
			assert.True(t, NumericCodeExists(test.NumericCode))
		})
	}
}

func TestNumericCodeExists_False(t *testing.T) {
	tests := []struct {
		name string
		num  int
	}{
		{"1000", 1000},
		{"0", 0},
		{"5555", 5555},
		{"-25", 25},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.False(t, NumericCodeExists(test.num))
		})
	}
}

func TestCountryByAlphaCode_NotFound(t *testing.T) {
	tests := []struct {
		name string
		code string
	}{
		{"more than 3 chars", "test"},
		{"less than 2 chars", "A"},
		{"not exists", "AA"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := CountryByAlphaCode(test.code)

			assert.Nil(t, c)
			assert.ErrorIs(t, err, ErrCountryNotFound)
		})
	}
}

func TestCountryByAlphaCode_OK(t *testing.T) {
	var (
		c   *Country
		err error
	)

	c, err = CountryByAlphaCode("USA")

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, 840, c.NumericCode)
	assert.Equal(t, "USA", c.Alpha3Code)
	assert.Equal(t, "US", c.Alpha2Code)
	assert.Equal(t, "United States of America", c.Name)

	c, err = CountryByAlphaCode("UA")

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, 804, c.NumericCode)
	assert.Equal(t, "UKR", c.Alpha3Code)
	assert.Equal(t, "UA", c.Alpha2Code)
	assert.Equal(t, "Ukraine", c.Name)
}

func TestNumericCodeExists_NotFound(t *testing.T) {
	tests := []struct {
		name string
		code int
	}{
		{"negative", -1},
		{"not exists", 12345},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := CountryByNumericCode(test.code)

			assert.Nil(t, c)
			assert.ErrorIs(t, err, ErrCountryNotFound)
		})
	}
}

func TestNumericCodeExists_OK(t *testing.T) {
	var (
		c   *Country
		err error
	)

	c, err = CountryByNumericCode(484)

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, 484, c.NumericCode)
	assert.Equal(t, "MEX", c.Alpha3Code)
	assert.Equal(t, "MX", c.Alpha2Code)
	assert.Equal(t, "Mexico", c.Name)

	c, err = CountryByNumericCode(246)

	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, 246, c.NumericCode)
	assert.Equal(t, "FIN", c.Alpha3Code)
	assert.Equal(t, "FI", c.Alpha2Code)
	assert.Equal(t, "Finland", c.Name)
}
