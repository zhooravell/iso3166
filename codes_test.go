package iso3166

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodes_Quantity(t *testing.T) {
	assert.Equal(t, countriesQuantity, len(countries))
}

func TestCodes_Uniq(t *testing.T) {
	for i, v := range countries {
		for j, c := range countries {
			if j == i {
				continue
			}

			if v.Name == c.Name {
				t.Fatal(fmt.Sprintf("duplicate name %s", v.Name))
			}

			if v.NumericCode == c.NumericCode {
				t.Fatal(fmt.Sprintf("duplicate num %d", v.NumericCode))
			}

			if v.Alpha2Code == c.Alpha2Code {
				t.Fatal(fmt.Sprintf("duplicate alpha2 %s", v.Alpha2Code))
			}

			if v.Alpha3Code == c.Alpha3Code {
				t.Fatal(fmt.Sprintf("duplicate alpha3 %s", v.Alpha3Code))
			}
		}
	}
}
