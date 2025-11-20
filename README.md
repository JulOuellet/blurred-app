### Dropping the database schema and data using golang-migrate:
```
migrate -path ./internal/db/migrations -database $DATABASE_URL drop
```

### Creating test data for local testing using postgresql:
```
psql $DATABASE_URL -f internal/db/scripts/init_data.sql
```
