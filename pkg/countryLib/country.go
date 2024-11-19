package countryLib

import (
	"countryApi/config"

	"github.com/gin-gonic/gin"
)

func CountryLib(router *gin.Engine) {
	group := router.Group(config.ENV.API_PREFIX)

	group.GET("/countries/", GetCountryList)
	group.GET("/countries/:id", GetCountry)
	group.GET("/cities/", GetCityList)
	group.GET("/countries/search", SearchCountries)
	group.GET("/cities/search", SearchCities)
}
