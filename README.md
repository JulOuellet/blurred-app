# Blurred

#### Running the applicationcwith hot reloading:
```bash
air
```

#### Templ hot reloading:
```bash
templ generate --watch
```

#### Tailwind hot reloading:
```bash
tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch
```

#### Running migrations using golang-migrate:
```bash
migrate -path ./internal/db/migrations -database $DATABASE_URL up
```

#### Dropping the database schema and data using golang-migrate:
```bash
migrate -path ./internal/db/migrations -database $DATABASE_URL drop
```

#### Creating test data for local testing using postgresql:
```bash
psql $DATABASE_URL -f internal/db/scripts/init_data.sql
```
