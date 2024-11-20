# countryApi

App is only dedicated to server country and city information to use in other applications, it uses separate db and standalone.

Put this repo folder in:
**~/country_backend/countryApi**

### BE CAREFUL, __make db__ drops the whole db, modify script beforehand
Update the app (need to have github access keys):
```bash
bash ~/country_backend/countryApi/scripts/update_country.sh
```

```bash
make db
make dev
```

To build application:

```bash
make build
```


