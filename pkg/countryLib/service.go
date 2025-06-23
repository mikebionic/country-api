package countryLib

import (
	"context"
	db "countryApi/database"
	"countryApi/pkg/utils"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCountryList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	offset := (page - 1) * perPage

	rows, err := db.DB.Query(
		context.Background(),
		GetCountriesPaginatedQuery,
		perPage,
		offset,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}
	defer rows.Close()

	var countries []Country
	var totalCount int

	for rows.Next() {
		var country Country
		err := rows.Scan(
			&country.ID,
			&country.Name,
			&country.ISO3,
			&country.NumericCode,
			&country.ISO2,
			&country.Phonecode,
			&country.Capital,
			&country.Currency,
			&country.CurrencyName,
			&country.CurrencySymbol,
			&country.TLD,
			&country.Native,
			&country.Region,
			&country.RegionID,
			&country.Subregion,
			&country.SubregionID,
			&country.Nationality,
			&country.Timezones,
			&country.Translations,
			&country.Latitude,
			&country.Longitude,
			&country.Emoji,
			&country.EmojiU,
			&country.CreatedAt,
			&country.UpdatedAt,
			&country.Flag,
			&country.WikiDataID,
			&totalCount,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Scan error", err.Error()))
			return
		}
		countries = append(countries, country)
	}

	response := utils.PaginatedResponse{
		Total:   totalCount,
		Page:    page,
		PerPage: perPage,
		Data:    countries,
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("Country list", response))
}

func GetCountry(ctx *gin.Context) {
	id := ctx.Param("id") // Get the country ID from the request URL

	var countryDetail CountryDetail
	var citiesJSON []byte

	err := db.DB.QueryRow(
		context.Background(),
		GetCountryWithCitiesQuery, // This is the updated query
		id,
	).Scan(
		&countryDetail.ID,
		&countryDetail.Name,
		&countryDetail.ISO3,
		&countryDetail.NumericCode,
		&countryDetail.ISO2,
		&countryDetail.Phonecode,
		&countryDetail.Capital,
		&countryDetail.Currency,
		&countryDetail.CurrencyName,
		&countryDetail.CurrencySymbol,
		&countryDetail.TLD,
		&countryDetail.Native,
		&countryDetail.Region,
		&countryDetail.RegionID,
		&countryDetail.Subregion,
		&countryDetail.SubregionID,
		&countryDetail.Nationality,
		&countryDetail.Timezones,
		&countryDetail.Translations,
		&countryDetail.Latitude,
		&countryDetail.Longitude,
		&countryDetail.Emoji,
		&countryDetail.EmojiU,
		&countryDetail.CreatedAt,
		&countryDetail.UpdatedAt,
		&countryDetail.Flag,
		&countryDetail.WikiDataID,
		&citiesJSON, // Store the cities as JSON
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.FormatErrorResponse("Country not found", err.Error()))
		return
	}

	err = json.Unmarshal(citiesJSON, &countryDetail.Cities)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Failed to unmarshal cities", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("Country details", countryDetail))
}

func GetCityList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	offset := (page - 1) * perPage

	rows, err := db.DB.Query(
		context.Background(),
		GetCitiesPaginatedQuery,
		perPage,
		offset,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}
	defer rows.Close()

	var cities []City
	var totalCount int

	for rows.Next() {
		var city City
		err := rows.Scan(
			&city.ID, &city.Name, &city.StateID, &city.StateCode,
			&city.CountryID, &city.CountryCode, &city.Latitude,
			&city.Longitude, &city.CreatedAt, &city.UpdatedAt,
			&city.Flag, &city.WikiDataID, &totalCount,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Scan error", err.Error()))
			return
		}
		cities = append(cities, city)
	}

	response := utils.PaginatedResponse{
		Total:   totalCount,
		Page:    page,
		PerPage: perPage,
		Data:    cities,
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("City list", response))
}

func SearchCountries(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	search := ctx.Query("q")
	if search == "" {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Search query is required", "Please provide a search term using the 'q' parameter"))
		return
	}

	offset := (page - 1) * perPage

	rows, err := db.DB.Query(
		context.Background(),
		SearchCountriesQuery,
		perPage,
		offset,
		search,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}
	defer rows.Close()

	var countries []Country
	var totalCount int

	for rows.Next() {
		var country Country
		err := rows.Scan(
			&country.ID,
			&country.Name,
			&country.ISO3,
			&country.NumericCode,
			&country.ISO2,
			&country.Phonecode,
			&country.Capital,
			&country.Currency,
			&country.CurrencyName,
			&country.CurrencySymbol,
			&country.TLD,
			&country.Native,
			&country.Region,
			&country.RegionID,
			&country.Subregion,
			&country.SubregionID,
			&country.Nationality,
			&country.Timezones,
			&country.Translations,
			&country.Latitude,
			&country.Longitude,
			&country.Emoji,
			&country.EmojiU,
			&country.CreatedAt,
			&country.UpdatedAt,
			&country.Flag,
			&country.WikiDataID,
			&totalCount,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Scan error", err.Error()))
			return
		}
		countries = append(countries, country)
	}

	response := utils.PaginatedResponse{
		Total:   totalCount,
		Page:    page,
		PerPage: perPage,
		Data:    countries,
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("Country search results", response))
}

func SearchCities(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	search := ctx.Query("q")
	if search == "" {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Search query is required", "Please provide a search term using the 'q' parameter"))
		return
	}

	offset := (page - 1) * perPage

	rows, err := db.DB.Query(
		context.Background(),
		SearchCitiesQuery,
		perPage,
		offset,
		search,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}
	defer rows.Close()

	var cities []CitySearchResult
	var totalCount int

	for rows.Next() {
		var city CitySearchResult
		err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.StateID,
			&city.StateCode,
			&city.CountryID,
			&city.CountryCode,
			&city.CountryName,
			&city.Latitude,
			&city.Longitude,
			&city.CreatedAt,
			&city.UpdatedAt,
			&city.Flag,
			&city.WikiDataID,
			&totalCount,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Scan error", err.Error()))
			return
		}
		cities = append(cities, city)
	}

	response := utils.PaginatedResponse{
		Total:   totalCount,
		Page:    page,
		PerPage: perPage,
		Data:    cities,
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("City search results", response))
}

func generateID() int64 {
	return time.Now().UnixNano() + rand.Int63n(1000)
}

func CreateCountry(ctx *gin.Context) {
	var req CreateCountryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid request body", err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Validation error", "Name is required"))
		return
	}

	now := time.Now()
	countryID := generateID() // Generate ID explicitly
	var returnedID int64

	err := db.DB.QueryRow(
		context.Background(),
		CreateCountryQuery,
		countryID, // Add ID as first parameter
		req.Name,
		req.ISO3,
		req.NumericCode,
		req.ISO2,
		req.Phonecode,
		req.Capital,
		req.Currency,
		req.CurrencyName,
		req.CurrencySymbol,
		req.TLD,
		req.Native,
		req.Region,
		req.RegionID,
		req.Subregion,
		req.SubregionID,
		req.Nationality,
		req.Timezones,
		req.Translations,
		req.Latitude,
		req.Longitude,
		req.Emoji,
		req.EmojiU,
		now,
		now,
		1, // flag default to 1 (active)
		req.WikiDataID,
	).Scan(&returnedID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}

	response := CreateCountryResponse{
		ID:        returnedID,
		Name:      req.Name,
		CreatedAt: now,
	}

	ctx.JSON(http.StatusCreated, utils.FormatResponse("Country created successfully", response))
}

func CreateCity(ctx *gin.Context) {
	var req CreateCityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid request body", err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Validation error", "Name is required"))
		return
	}
	if req.CountryID == 0 {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Validation error", "Country ID is required"))
		return
	}

	now := time.Now()
	cityID := generateID() // Generate ID explicitly
	var returnedID int64

	err := db.DB.QueryRow(
		context.Background(),
		CreateCityQuery,
		cityID, // Add ID as first parameter
		req.Name,
		req.StateID,
		req.StateCode,
		req.CountryID,
		req.CountryCode,
		req.Latitude,
		req.Longitude,
		now,
		now,
		1, // flag default to 1 (active)
		req.WikiDataID,
	).Scan(&returnedID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}

	response := CreateCityResponse{
		ID:        returnedID,
		Name:      req.Name,
		CountryID: req.CountryID,
		CreatedAt: now,
	}

	ctx.JSON(http.StatusCreated, utils.FormatResponse("City created successfully", response))
}

func UpdateCountry(ctx *gin.Context) {
	id := ctx.Param("id")
	countryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid country ID", err.Error()))
		return
	}

	var req UpdateCountryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid request body", err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Validation error", "Name is required"))
		return
	}

	now := time.Now()
	result, err := db.DB.Exec(
		context.Background(),
		UpdateCountryQuery,
		req.Name,
		req.ISO3,
		req.NumericCode,
		req.ISO2,
		req.Phonecode,
		req.Capital,
		req.Currency,
		req.CurrencyName,
		req.CurrencySymbol,
		req.TLD,
		req.Native,
		req.Region,
		req.RegionID,
		req.Subregion,
		req.SubregionID,
		req.Nationality,
		req.Timezones,
		req.Translations,
		req.Latitude,
		req.Longitude,
		req.Emoji,
		req.EmojiU,
		now,
		req.WikiDataID,
		countryID,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, utils.FormatErrorResponse("Country not found", "No country found with the given ID"))
		return
	}

	response := UpdateCountryResponse{
		ID:        countryID,
		Name:      req.Name,
		UpdatedAt: now,
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("Country updated successfully", response))
}

func DeleteCountry(ctx *gin.Context) {
	id := ctx.Param("id")
	countryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid country ID", err.Error()))
		return
	}

	result, err := db.DB.Exec(
		context.Background(),
		DeleteCountryQuery,
		time.Now(),
		countryID,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, utils.FormatErrorResponse("Country not found", "No country found with the given ID"))
		return
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("Country deleted successfully", nil))
}

func UpdateCity(ctx *gin.Context) {
	id := ctx.Param("id")
	cityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid city ID", err.Error()))
		return
	}

	var req UpdateCityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid request body", err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Validation error", "Name is required"))
		return
	}
	if req.CountryID == 0 {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Validation error", "Country ID is required"))
		return
	}

	now := time.Now()
	result, err := db.DB.Exec(
		context.Background(),
		UpdateCityQuery,
		req.Name,
		req.StateID,
		req.StateCode,
		req.CountryID,
		req.CountryCode,
		req.Latitude,
		req.Longitude,
		now,
		req.WikiDataID,
		cityID,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, utils.FormatErrorResponse("City not found", "No city found with the given ID"))
		return
	}

	response := UpdateCityResponse{
		ID:        cityID,
		Name:      req.Name,
		CountryID: req.CountryID,
		UpdatedAt: now,
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("City updated successfully", response))
}

func DeleteCity(ctx *gin.Context) {
	id := ctx.Param("id")
	cityID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.FormatErrorResponse("Invalid city ID", err.Error()))
		return
	}

	result, err := db.DB.Exec(
		context.Background(),
		DeleteCityQuery,
		time.Now(),
		cityID,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.FormatErrorResponse("Database error", err.Error()))
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, utils.FormatErrorResponse("City not found", "No city found with the given ID"))
		return
	}

	ctx.JSON(http.StatusOK, utils.FormatResponse("City deleted successfully", nil))
}
