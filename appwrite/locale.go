package appwrite

import (
)

// Locale service
type Locale struct {
	client Client
}

func NewLocale(clt Client) *Locale {
	return &Locale{
		client: clt,
	}
}

// Get get the current user location based on IP. Returns an object with user
// country code, country name, continent name, continent code, ip address and
// suggested currency. You can use the locale header to get the data in a
// supported language.
// 
// ([IP Geolocation by DB-IP](https://db-ip.com))
func (srv *Locale) Get() (*ClientResponse, error) {
	path := "/locale"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListContinents list of all continents. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListContinents() (*ClientResponse, error) {
	path := "/locale/continents"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListCountries list of all countries. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListCountries() (*ClientResponse, error) {
	path := "/locale/countries"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListCountriesEU list of all countries that are currently members of the EU.
// You can use the locale header to get the data in a supported language.
func (srv *Locale) ListCountriesEU() (*ClientResponse, error) {
	path := "/locale/countries/eu"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListCountriesPhones list of all countries phone codes. You can use the
// locale header to get the data in a supported language.
func (srv *Locale) ListCountriesPhones() (*ClientResponse, error) {
	path := "/locale/countries/phones"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListCurrencies list of all currencies, including currency symbol, name,
// plural, and decimal digits for all major and minor currencies. You can use
// the locale header to get the data in a supported language.
func (srv *Locale) ListCurrencies() (*ClientResponse, error) {
	path := "/locale/currencies"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListLanguages list of all languages classified by ISO 639-1 including
// 2-letter code, name in English, and name in the respective language.
func (srv *Locale) ListLanguages() (*ClientResponse, error) {
	path := "/locale/languages"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}
