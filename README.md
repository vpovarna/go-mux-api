### API SPECIFICATION

In concrete terms, our application should:

- Create a new product in response to a valid POST request at /product,
- Update a product in response to a valid PUT request at /product/{id},
- Delete a product in response to a valid DELETE request at /product/{id},
- Fetch a product in response to a valid GET request at /product/{id}, and
- Fetch a list of products in response to a valid GET request at /products.


### PostgreSQL
Starting docker container:
```
$ docker run -it -p 5432:5432 -d postgres
```
