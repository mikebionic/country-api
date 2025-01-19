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
