# Blurred

#### Running the applicationcwith hot reloading:
```
air
```

#### Templ hot reloading:
```
templ generate --watch
```

#### Tailwind hot reloading:
```
tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch
```

#### Running migrations using golang-mgrate:
```
migrate -path ./internal/db/migrations -database $DATABASE_URL up
```

#### Dropping the database schema and data using golang-migrate:
```
migrate -path ./internal/db/migrations -database $DATABASE_URL drop
```

#### Creating test data for local testing using postgresql:
```
psql $DATABASE_URL -f internal/db/scripts/init_data.sql
```
