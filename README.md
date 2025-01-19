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

### Contributing

If you find an issue with the Postman collection or would like to enhance it, feel free to submit a pull request.
