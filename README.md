# Restaurants project

## How to run app

### Using app golang directly on Terminal
Provide all needed correct values.
```shell
go run ./cmd/abr-plus \
-dsn="postgres://password:pa55word@localhost:5432/lecture6?sslmode=disable" \
-migrations=file://pkg/abr-plus/migrations \
-fill=true \
-env=development \
-port=8081
```

#### List of flags
`dsn` — postgress connection string with username, password, address, port, database name, and SSL mode. Default: `Value is not correct by security reasons`.

`migrations` — Path to folder with migration files. If not provided, migrations do not applied.

`fill` — Fill database with dummy data. Default: `false`.

`env` - App running mode. Default: `development`

`port` - App port. Default: `8081`

### Run with docker-compose

```shell
env POSTGRES_PASSWORD="STRONG_PASSWORD" APP_DSN="postgres://postgres:postgres@db:5432/example?sslmode=disable" docker-compose --env-file .env.example up --build
```

`env POSTGRES_PASSWORD="postgres"` this command adds envoirment variable then available in docker compose.

`--build` flag force docker compose to rebuild app. For example, if you have changed source code, you need this flag.

#### Add write menus permission to user example SQL

```sql
-- Add write menus permission to user with id 1
INSERT INTO users_permissions
SELECT 1, permissions.id
FROM permissions
WHERE permissions.code = 'menus:write';
```

## Restaurants REST API

```
# list of all menu items
GET /menus
POST /menus
GET /menus/:id
PUT /menus/:id
DELETE /menus/:id
```

## DB Structure

```
// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table restaurants {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  title text
  coordinates text
  address text
  cousine text
}

Table menu {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  title text
  description text
  nutrition_value text
}

// many-to-many
Table restaurants_and_menu {
  id bigserial [primary key]
  created_at timestamp
  updated_at timestamp
  restaurant bigserial
  menu bigserial
}

Ref: restaurants_and_menu.restaurant < restaurants.id
Ref: restaurants_and_menu.menu < menu.id

```
