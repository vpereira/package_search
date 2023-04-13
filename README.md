To run the docker environment:

```
docker-compose build && docker-compose up
```

To enter in the app docker environment:

```
docker-compose exec app bash
```

Then if you want to run the tests:

```
make test
```

If you want to run the tool:

```
make build && ./package_search
```

If you want to test it locally:

```
curl "http://localhost:8080/record?name=package1"
```