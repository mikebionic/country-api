package countryLib

import (
	"countryApi/config"

	"github.com/gin-gonic/gin"
)

func CountryLib(router *gin.Engine) {
	group := router.Group(config.ENV.API_PREFIX)

	group.GET("/countries/", GetCountryList)
	group.GET("/countries/:id", GetCountry)
	group.POST("/countries", CreateCountry)
	group.PUT("/countries/:id", UpdateCountry)
	group.DELETE("/countries/:id", DeleteCountry)
	group.GET("/countries/search", SearchCountries)

	group.GET("/cities/", GetCityList)
	group.POST("/cities", CreateCity)
	group.PUT("/cities/:id", UpdateCity)
	group.DELETE("/cities/:id", DeleteCity)
	group.GET("/cities/search", SearchCities)
}
