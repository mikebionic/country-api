# CountryApi Go

This application is a standalone service designed specifically to provide server country and city information. It uses a dedicated database and can be seamlessly integrated into other applications to supply geolocation data.

## Docker Image Deployment

```bash
# Ensure you've installed docker on your machine.
cd country-api/
docker compose up -d
```

Check **.env.docker** for more info
APP runs on port **7888**, configure your proxy manager accordingly

---

## Manual Deployment

> Please double check scripts in case if they work incorrectly, check folders, permissions and configurations!

Put this repo folder in:
**~/country_backend/countryApi**

### BE CAREFUL, **make db** drops the whole db, modify script beforehand

Update the application using script:

```bash
bash ~/country_backend/countryApi/scripts/update.sh
```

```bash
make db
make dev
```

Building the application:

```bash
make build
```

---

## Postman Collection

This repository includes a Postman collection to help you test and interact with the application's API endpoints. The collection contains pre-configured requests for retrieving server country and city information.

### How to Use the Collection

1. **Install Postman**  
   Make sure you have [Postman](https://www.postman.com/downloads/) installed on your system.

2. **Import the Collection**

   - Download the Postman collection JSON file from this repository.
   - Open Postman, click on the **Import** button, and upload the JSON file.

3. **Set Up Environment Variables** _(Optional but recommended)_

   - The collection supports environment variables to simplify testing.
   - Create a new environment in Postman and define the following variables:
     - `BASE_URL`: The base URL of your API (e.g., `http://localhost:7888` or your production URL).

4. **Run the Requests**
   - After importing, you can test each API endpoint directly from the collection.
   - Ensure the server is running, and make sure the `BASE_URL` is set correctly.


### Summary of Endpoints
Countries:
POST   /api/countries        - Create country
GET    /api/countries        - List countries (paginated)
GET    /api/countries/:id    - Get country details
PUT    /api/countries/:id    - Update country
DELETE /api/countries/:id    - Delete country (soft)
GET    /api/countries/search - Search countries

Cities:
POST   /api/cities           - Create city
GET    /api/cities           - List cities (paginated)
PUT    /api/cities/:id       - Update city
DELETE /api/cities/:id       - Delete city (soft)
GET    /api/cities/search    - Search cities

---

### API Endpoints Overview

Here are some key endpoints included in the Postman collection:

- **`GET /country`**  
  Retrieves the country information for the server.  
  **Example Response:**

  ```json
  {
    "country": "United States"
  }
  ```

- **`GET /city`**  
  Retrieves the city information for the server.  
  **Example Response:**

  ```json
  {
    "city": "San Francisco"
  }
  ```

- **`GET /location`**  
  Combines both country and city information.  
  **Example Response:**
  ```json
  {
    "country": "United States",
    "city": "San Francisco"
  }
  ```

### **Create Country (Full):**
```json
{
  "name": "Test Country",
  "iso3": "TST",
  "iso2": "TS",
  "capital": "Test Capital",
  "currency": "TST",
  "latitude": 20.593684,
  "longitude": 78.96288
}
```

### **Create City (Minimal):**
```json
{
  "name": "Test City",
  "country_id": 1
}
```

## **Expected Responses:**

### **Success Response:**
```json
{
  "message": "Country created successfully",
  "success": true,
  "data": {
    "id": 1,
    "name": "Test Country",
    "created_at": "2025-01-01T12:00:00Z"
  },
  "errorMsg": ""
}
```

### **Error Response:**
```json
{
  "message": "Validation error",
  "success": false,
  "data": null,
  "errorMsg": "Name is required"
}
```


### Contributing

If you find an issue with the Postman collection or would like to enhance it, feel free to submit a pull request.
