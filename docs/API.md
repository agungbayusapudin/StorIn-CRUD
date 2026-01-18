# API Documentation

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

Semua endpoint (kecuali `/users/register` dan `/users/login`) memerlukan JWT token di header:

```
Authorization: Bearer <token>
```

## Response Format

### Success Response

```json
{
  "status": "success",
  "code": 200,
  "message": "Operation successful",
  "data": {
    // Response data
  }
}
```

### Error Response

```json
{
  "status": "error",
  "code": 400,
  "message": "Error message",
  "errors": [
    {
      "field": "email",
      "message": "Email is required"
    }
  ]
}
```

## User Endpoints

### Register User

**Endpoint:** `POST /users/register`

**Description:** Mendaftarkan user baru

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}
```

**Response (201 Created):**
```json
{
  "status": "success",
  "code": 201,
  "message": "User registered successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "name": "John Doe",
    "created_at": "2026-01-07T14:00:00Z"
  }
}
```

**Error (400 Bad Request):**
```json
{
  "status": "error",
  "code": 400,
  "message": "Validation error",
  "errors": [
    {
      "field": "email",
      "message": "Email already exists"
    }
  ]
}
```

### Login

**Endpoint:** `POST /users/login`

**Description:** Login user dan mendapatkan JWT token

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@example.com",
      "name": "John Doe"
    }
  }
}
```

**Error (401 Unauthorized):**
```json
{
  "status": "error",
  "code": 401,
  "message": "Invalid email or password"
}
```

### Get User Profile

**Endpoint:** `GET /users/profile`

**Description:** Mendapatkan profil user yang sedang login

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "User profile retrieved",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "name": "John Doe",
    "phone": "+62812345678",
    "created_at": "2026-01-07T14:00:00Z",
    "updated_at": "2026-01-07T14:00:00Z"
  }
}
```

### Update User Profile

**Endpoint:** `PUT /users/profile`

**Description:** Update profil user

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Jane Doe",
  "phone": "+62812345678"
}
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "User profile updated",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "name": "Jane Doe",
    "phone": "+62812345678",
    "updated_at": "2026-01-07T14:30:00Z"
  }
}
```

## Product Endpoints

### Get All Products

**Endpoint:** `GET /products`

**Description:** Mendapatkan daftar semua produk dengan pagination

**Query Parameters:**
- `page` (optional, default: 1) - Halaman
- `limit` (optional, default: 10) - Jumlah item per halaman
- `category` (optional) - Filter berdasarkan kategori
- `search` (optional) - Search berdasarkan nama produk

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Products retrieved",
  "data": {
    "items": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440001",
        "name": "Product 1",
        "description": "Product description",
        "price": 99.99,
        "category": "electronics",
        "stock": 50,
        "created_at": "2026-01-07T14:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 100,
      "total_pages": 10
    }
  }
}
```

### Get Product by ID

**Endpoint:** `GET /products/:id`

**Description:** Mendapatkan detail produk berdasarkan ID

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product retrieved",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "Product 1",
    "description": "Product description",
    "price": 99.99,
    "category": "electronics",
    "stock": 50,
    "created_at": "2026-01-07T14:00:00Z",
    "updated_at": "2026-01-07T14:00:00Z"
  }
}
```

**Error (404 Not Found):**
```json
{
  "status": "error",
  "code": 404,
  "message": "Product not found"
}
```

### Create Product

**Endpoint:** `POST /products`

**Description:** Membuat produk baru

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "New Product",
  "description": "Product description",
  "price": 99.99,
  "category": "electronics",
  "stock": 100
}
```

**Response (201 Created):**
```json
{
  "status": "success",
  "code": 201,
  "message": "Product created",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440002",
    "name": "New Product",
    "description": "Product description",
    "price": 99.99,
    "category": "electronics",
    "stock": 100,
    "created_at": "2026-01-07T14:30:00Z"
  }
}
```

### Update Product

**Endpoint:** `PUT /products/:id`

**Description:** Update produk

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Updated Product",
  "price": 149.99,
  "stock": 80
}
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product updated",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440002",
    "name": "Updated Product",
    "description": "Product description",
    "price": 149.99,
    "category": "electronics",
    "stock": 80,
    "updated_at": "2026-01-07T14:45:00Z"
  }
}
```

### Delete Product

**Endpoint:** `DELETE /products/:id`

**Description:** Menghapus produk

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Product deleted"
}
```

## Billing Endpoints

### Get All Invoices

**Endpoint:** `GET /billing/invoices`

**Description:** Mendapatkan daftar invoice user

**Query Parameters:**
- `page` (optional, default: 1) - Halaman
- `limit` (optional, default: 10) - Jumlah item per halaman
- `status` (optional) - Filter berdasarkan status (pending, paid, cancelled)

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Invoices retrieved",
  "data": {
    "items": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440010",
        "user_id": "550e8400-e29b-41d4-a716-446655440000",
        "total_amount": 199.98,
        "status": "pending",
        "created_at": "2026-01-07T14:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 50,
      "total_pages": 5
    }
  }
}
```

### Get Invoice by ID

**Endpoint:** `GET /billing/invoices/:id`

**Description:** Mendapatkan detail invoice

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Invoice retrieved",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440010",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "total_amount": 199.98,
    "status": "pending",
    "items": [
      {
        "product_id": "550e8400-e29b-41d4-a716-446655440001",
        "product_name": "Product 1",
        "quantity": 2,
        "price": 99.99,
        "subtotal": 199.98
      }
    ],
    "created_at": "2026-01-07T14:00:00Z",
    "updated_at": "2026-01-07T14:00:00Z"
  }
}
```

### Create Invoice

**Endpoint:** `POST /billing/invoices`

**Description:** Membuat invoice baru

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "items": [
    {
      "product_id": "550e8400-e29b-41d4-a716-446655440001",
      "quantity": 2,
      "price": 99.99
    }
  ]
}
```

**Response (201 Created):**
```json
{
  "status": "success",
  "code": 201,
  "message": "Invoice created",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440011",
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "total_amount": 199.98,
    "status": "pending",
    "created_at": "2026-01-07T14:30:00Z"
  }
}
```

### Pay Invoice

**Endpoint:** `POST /billing/invoices/:id/pay`

**Description:** Membayar invoice

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "payment_method": "credit_card",
  "amount": 199.98
}
```

**Response (200 OK):**
```json
{
  "status": "success",
  "code": 200,
  "message": "Invoice paid",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440011",
    "status": "paid",
    "paid_at": "2026-01-07T14:35:00Z"
  }
}
```

## Health Check

**Endpoint:** `GET /health`

**Description:** Check aplikasi health status

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2026-01-07T14:00:00Z",
  "database": "connected"
}
```

## Error Codes

| Code | Message | Description |
|------|---------|-------------|
| 200 | OK | Request successful |
| 201 | Created | Resource created successfully |
| 400 | Bad Request | Invalid request parameters |
| 401 | Unauthorized | Missing or invalid authentication |
| 403 | Forbidden | User doesn't have permission |
| 404 | Not Found | Resource not found |
| 409 | Conflict | Resource already exists |
| 500 | Internal Server Error | Server error |
| 503 | Service Unavailable | Service temporarily unavailable |

## Rate Limiting

API memiliki rate limiting:
- 100 requests per minute per IP
- 1000 requests per hour per user

Headers response:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1641564000
```

## Pagination

Untuk endpoint yang mengembalikan list, gunakan query parameters:

```
GET /products?page=1&limit=10
```

Response pagination:
```json
{
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

---

**Last Updated:** January 7, 2026
