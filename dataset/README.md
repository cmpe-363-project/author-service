# Load Quotes Dataset

## Usage

> **Make sure the mysql services are running and databases are created.**

Test parsing first (recommended):
```bash
cd dataset
./load_quotes.sh --dry-run
```


Load data into databases:
```bash
./load_quotes.sh
```

## Requirements

- Python 3
- Docker containers `author-service-mysql` and `quote-service-mysql` must be running
- Databases `author_service` and `quote_service` must exist with tables `authors` and `quotes`
