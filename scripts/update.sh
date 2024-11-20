#!/bin/bash
cd ~/country_backend/country-api || { echo "Directory not found"; exit 1; }
git pull origin main
sudo systemctl stop countryApi.service
make build
make db
cp ~/country_backend/country-api/bin/countryApi ~/country_backend/app/countryApi
sudo systemctl start countryApi.service
echo "Update completed successfully!"


