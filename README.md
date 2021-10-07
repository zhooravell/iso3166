GoLang ISO-3166
===============
> 3166-1 alpha-2, 3166-1 alpha-3, 3166-1 numeric

[![License][license-image]][license-link]

# Examples

```go
package main

import (
	"fmt"
	"github.com/zhooravell/iso3166"
)

func main() {
	if iso3166.Alpha2CodeExists("GB") {
		fmt.Println("GB iso code exists!")
	}

	c, err := iso3166.CountryByAlphaCode("TUR")

	if err != nil {
		fmt.Printf("Num: %d; Alpha2: %s; Alpha3: %s; Name: %s\n", c.NumericCode, c.Alpha2Code, c.Alpha3Code, c.Name)
		// Output: Num: 792; Alpha2: TR; Alpha3: TUR; Name: Turkey
	}
}
```

# Sources

* [Wiki](https://en.wikipedia.org/wiki/ISO_3166)
* [iso.org](https://www.iso.org/ru/iso-3166-country-codes.html)

[license-link]: https://github.com/zhooravell/iso3166/blob/master/LICENSE

[license-image]: https://img.shields.io/dub/l/vibe-d.svg