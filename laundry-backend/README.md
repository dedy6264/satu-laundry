# Laundry Backend API

Backend API untuk sistem manajemen laundry menggunakan Golang dengan framework Echo dan arsitektur Clean Architecture.

## Fitur

- Authentication (Login)
- Manajemen Brand
- Manajemen Cabang
- Manajemen Outlet
- Logging request/response terintegrasi

## Teknologi

- Golang
- Echo Framework
- PostgreSQL
- Clean Architecture
- JWT Authentication
- Logrus untuk logging

## Struktur Direktori

```
laundry-backend/
├── cmd/
│   └── main.go
├── configs/
├── internal/
│   ├── entities/
│   ├── repositories/
│   ├── usecases/
│   ├── delivery/
│   ├── middleware/
│   └── utils/
├── pkg/
└── docs/
```

## Instalasi

1. Clone repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Setup environment variables (lihat .env.example)
4. Jalankan aplikasi:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Authentication
- `POST /api/v1/login` - Login user

### Brand
- `POST /api/v1/brands` - Create brand
- `GET /api/v1/brands/:id` - Get brand by ID
- `GET /api/v1/brands` - Get all brands
- `PUT /api/v1/brands/:id` - Update brand
- `DELETE /api/v1/brands/:id` - Delete brand

### Cabang
- `POST /api/v1/cabangs` - Create cabang
- `GET /api/v1/cabangs/:id` - Get cabang by ID
- `GET /api/v1/cabangs/brand/:brand_id` - Get cabangs by brand ID
- `GET /api/v1/cabangs` - Get all cabangs
- `PUT /api/v1/cabangs/:id` - Update cabang
- `DELETE /api/v1/cabangs/:id` - Delete cabang

### Outlet
- `POST /api/v1/outlets` - Create outlet
- `GET /api/v1/outlets/:id` - Get outlet by ID
- `GET /api/v1/outlets/cabang/:cabang_id` - Get outlets by cabang ID
- `GET /api/v1/outlets` - Get all outlets
- `PUT /api/v1/outlets/:id` - Update outlet
- `DELETE /api/v1/outlets/:id` - Delete outlet

## Logging

Setiap request dan response akan di-log secara otomatis dengan informasi:
- Timestamp
- HTTP method
- URL
- Status code
- Durasi eksekusi
- Request body
- Response body
- User agent
- Client IP

## Environment Variables

Lihat file `.env` untuk konfigurasi environment variables.