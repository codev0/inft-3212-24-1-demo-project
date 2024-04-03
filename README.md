# Restaurants project

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