CREATE TABLE IF NOT EXISTS restaurants
(
    id          bigserial PRIMARY KEY,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title       text                        NOT NULL,
    description text                        NOT NULL,
    coordinates text                        NOT NULL,
    address     text                        NOT NULL,
    cousine     text                        NOT NULL
);

CREATE TABLE IF NOT EXISTS menus
(
    id              bigserial PRIMARY KEY,
    created_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title           text                        NOT NULL,
    description     text,
    nutrition_value int
);

CREATE TABLE IF NOT EXISTS restaurants_and_menus
(
    "id"         bigserial PRIMARY KEY,
    "created_at" timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    "restaurant" bigserial,
    "menu"       bigserial,
    FOREIGN KEY (restaurant)
        REFERENCES restaurants(id),
    FOREIGN KEY (menu)
        REFERENCES menus(id)
);
