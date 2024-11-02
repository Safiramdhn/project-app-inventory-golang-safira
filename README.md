# Inventory Management System
A command-line inventory management system built with Go that handles items, transactions, categories, and locations.

## Features
User authentication (Login/Logout)
Item management (CRUD operations)
Transaction management
Category management
Location management
Role-based access control (Admin and Staff Warehouse)
## Prerequisites
Go 1.x
PostgreSQL
github.com/lib/pq
## Installation
**Clone the repository**
```git clone https://github.com/Safiramdhn/project-app-inventory-golang-safira.git```
**Set up the database**
```psql -U postgres -d inventory < inventory.sql```
**Run the application**
```go run main.go```

## API Endpoints & Examples
### Authentication
1. Login
   
```
// body.json
{
    "username": "admin",
    "password": "adminpassword"
}
```
Response:
```
{
  "statusCode": 200,
  "message": "Logged in successfully",
  "data": {
    "id": 1,
    "username": "admin",
    "role": "admin"
  }
}
```
3. Logout
Response:
```
{
  "statusCode": 200,
  "message": "Logged out successfully"
}
```
### Items
1. Add Item
```
// body.json
{
    "name": "New Laptop",
    "category": {
        "id": 1
    },
    "location": {
        "id": 1
    },
    "quantity": 10,
    "price": 1000000
}
```
Response:
```{
  "statusCode": 201,
  "message": "Item added successfully"
}
```
2. Update Item
```// body.json
{
    "id": 1,
    "name": "Updated Laptop",
    "price": 1500000
}
```
Response:
```
{
  "statusCode": 200,
  "message": "Item updated successfully"
}
```
3. Delete Item
```
// body.json
{
    "id": 1
}`
Response:
`{
  "statusCode": 200,
  "message": "Item deleted successfully"
}
```
4. Get Items (with pagination)
```
{
    "page": 1,
    "per_page": 10
}
```
Response:
```{
  "statusCode": 200,
  "message": "Items retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Laptop",
      "category": {
        "id": 1,
        "name": "Electronics"
      },
      "location": {
        "id": 1,
        "name": "Warehouse A"
      },
      "quantity": 15,
      "price": 1000000,
      "pagination": {
        "page": 1,
        "per_page": 10,
        "countData": 50
      }
    }
    // ... more items
  ]
}
```
### Categories
Create Category
```
// body.json
{
    "name": "New Category"
}
```
Response:
```
{
  "statusCode": 200,
  "message": "Category created successfully"
}
```
Locations
```
// body.json
{
    "name": "New Warehouse",
    "address": "123 Business Street"
}
```
Response:
```
{
  "statusCode": 200,
  "message": "Location created successfully"
}
```
### Transactions
1. Create Transaction
```
// body.json
{
    "item": {
        "id": 1
    },
    "quantity": 5,
    "type": "OUT",
    "description": "Sold to customer"
}
```
Response:
```
{
  "statusCode": 200,
  "message": "Transaction created successfully"
}
```
2. Get Transactions (with filter)
```
{
    "type": "OUT",
    "item": {
        "name": "Laptop"
    },
    "pagination": {
        "page": 1,
        "per_page": 10
    }
}
```
Response:
```
{
  "statusCode": 200,
  "message": "Transactions retrieved successfully",
  "data": [
    {
      "id": 1,
      "item": {
        "id": 1,
        "name": "Laptop",
        "price": 1000000
      },
      "quantity": 5,
      "total_price": 5000000,
      "timestamp": "2023-07-20T10:00:00Z",
      "type": "OUT",
      "description": "Sold to customer",
      "pagination": {
        "page": 1,
        "per_page": 10,
        "countData": 25
      }
    }
    // ... more transactions
  ]
}
```
## Response Status Codes
Status  |  Code	Description
200	    |  Success
201     |  Created successfully
400     |  Bad Request / Validation Error
401     |  Unauthorized
500     |  Internal Server Error
Error Responses
```
{
  "statusCode": 400,
  "message": "Error message here"
}
```
## Role Access
### Admin
- Manage items (CRUD)
- Manage categories
- Manage locations
- View transactions
### Staff Warehouse
- Create transactions
- View transactions
- View items
## Database Schema
The database schema is available in inventory.sql
![inventory](https://github.com/user-attachments/assets/1875d54b-b6bc-4795-9f02-a597c79a3758)
