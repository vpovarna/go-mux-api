### API SPECIFICATION

In concrete terms, our application should:

- Create a new product in response to a valid POST request at /product,
- Update a product in response to a valid PUT request at /product/{id},
- Delete a product in response to a valid DELETE request at /product/{id},
- Fetch a product in response to a valid GET request at /product/{id}, and
- Fetch a list of products in response to a valid GET request at /products.

### Export variables
These variables need to be exported in your terminal before starting the app:
```
export APP_DB_USERNAME=postgres
export APP_DB_PASSWORD=
export APP_DB_NAME=postgres
```

### PostgreSQL
Starting postgres docker container:
```
$ cd docker
$ docker-compose up -d
```

### Build server
```
$ go build -v -o product-api
```
 
### Running tests
Ensure the above env variables are exported. From the `cmd/server-run` folder, run:
```
$ go test -v
```