# Author Service

## Starting the Service

Start database:

```bash
docker compose up mysql -d
```

Create the database:

```bash
docker exec author-service-mysql mysql -uroot -proot -e "CREATE DATABASE author_service;"
```

Start the Author Service:

```bash
docker compose up restapi
```

Add sample authors: (TODO: this will be updated with the dummy quote databse being found)

```bash
docker exec author-service-mysql mysql -uroot -proot author_service -e "INSERT INTO authors (name) VALUES ('Author One'), ('Author Two'), ('Author Three');"
```


## API Endpoints

- `GET /api/version`
  - ```json
    {
      "version": "1.0.0"
    }
    ```

- `GET /api/authors/by-id?id=1,2,3`
  - ```json
    {
      "items": [
        {
          "id": 1,
          "name": "Author Name"
        },
        {
          "id": 2,
          "name": "Author One"
        },
        {
          "id": 3,
          "name": "Author Two"
        }
      ]
    }
    ```