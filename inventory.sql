CREATE TYPE "transaction_type_enum" AS ENUM (
  'IN',
  'OUT'
);

CREATE TYPE "status_enum" AS ENUM (
	'active',
	'deleted'
)

CREATE TABLE "admins" (
	id	serial PRIMARY KEY,
	username varchar,
	password varchar
)

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "name" varchar(255),
  "status" status_enum DEFAULT 'active',
  "created_at" timestamp DEFAULT (NOW()),
  "updated_at" timestamp
);

CREATE TABLE "locations" (
  "id" serial PRIMARY KEY,
  "name" varchar(255),
  "address" varchar,
  "status" status_enum DEFAULT 'active',
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp
);

CREATE TABLE "items" (
  "id" serial PRIMARY KEY,
  "name" varchar(255),
  "category_id" int,
  "location_id" int,
  "quantity" int,
  "status" status_enum DEFAULT 'active',
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp
);

CREATE TABLE "transactions" (
  "id" serial PRIMARY KEY,
  "item_id" int,
  "type" transaction_type_enum,
  "quantity" int,
  "description" varchar(255),
  "timestamp" timestamp,
  "status" status_enum DEFAULT 'active',
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp
);

ALTER TABLE "items" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

INSERT INTO admins (username, password) VALUES ('admin', 'adminpassword');

INSERT INTO categories (name) VALUES
('Electronics'),
('Raw Materials'),
('Furniture'),
('Stationery');

SELECT * FROM categories

INSERT INTO locations (name, address) VALUES
('Warehouse A', '123 Main St, Cityville'),
('Warehouse B', '456 Second St, Townsville'),
('Retail Store', '789 Third St, Villagetown');

SELECT * FROM locations

INSERT INTO items (name, category_id, location_id, quantity) VALUES
('Laptop', 1, 1, 15),
('Smartphone', 1, 1, 25),
('Steel Sheets', 2, 1, 50),
('Wooden Table', 3, 2, 10),
('Notebook', 4, 3, 100),
('Printer Paper', 4, 3, 5);  -- This item has low stock

SELECT * FROM items

INSERT INTO transactions (item_id, type, quantity, description, timestamp) VALUES
(1, 'IN', 10, 'Received new shipment of laptops', NOW()),
(2, 'IN', 20, 'Received new shipment of smartphones', NOW()),
(1, 'OUT', 5, 'Sold 5 laptops', NOW()),
(3, 'IN', 30, 'Received steel sheets for production', NOW()),
(4, 'OUT', 2, 'Shipped 2 wooden tables', NOW()),
(5, 'OUT', 20, 'Distributed notebooks to staff', NOW()),
(6, 'OUT', 3, 'Sold printer paper to customers', NOW());

SELECT * FROM transactions
