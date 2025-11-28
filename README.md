# Author Service

## Starting the Service

```bash
docker compose up restapi
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