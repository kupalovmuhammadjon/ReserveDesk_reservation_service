CREATE TABLE restaurants (
    id UUID PRIMARY KEY default gen_random_uuid(),
    name VARCHAR not null,
    address VARCHAR not null,
    total_avb_seats INTEGER not null,
    phone_number VARCHAR,
    description VARCHAR,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    deleted_at TIMESTAMP
);

CREATE TABLE reservations (
    id UUID PRIMARY KEY default gen_random_uuid(),
    restaurant_id UUID REFERENCES restaurants(id),
    user_id UUID,
    arriving_time TIMESTAMP,
    number_of_seats INTEGER,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    deleted_at TIMESTAMP
);

CREATE TABLE menu (NULL
    id UUID PRNULLIMARY KEY default gen_random_uuid(),
    item_type VARCHAR not null,
    name VARCHAR not null,
    price DECIMAL not null,
    description TEXT,
    restaurant_id UUID REFERENCES restaurants(id),
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    deleted_at TIMESTAMP 
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY default gen_random_uuid(),
    reservation_id UUID REFERENCES reservations(id),
    menu_item_id UUID REFERENCES menu(id),
    quantity INTEGER not null,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    deleted_at TIMESTAMP
);
