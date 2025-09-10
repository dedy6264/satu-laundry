-- Script untuk inisialisasi database PostgreSQL

-- Buat database (jalankan sebagai superuser)
-- CREATE DATABASE laundry_db;

-- \c laundry_db;

-- Tabel Users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    role VARCHAR(50) DEFAULT 'admin',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Brand (mengacu pada struktur yang sudah dibuat sebelumnya)
CREATE TABLE IF NOT EXISTS brand (
    id_brand SERIAL PRIMARY KEY,
    nama_brand VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    tanggal_berdiri DATE,
    pic_nama VARCHAR(100),
    pic_email VARCHAR(100),
    pic_telepon VARCHAR(20),
    logo_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Cabang
CREATE TABLE IF NOT EXISTS cabang (
    id_cabang SERIAL PRIMARY KEY,
    id_brand INTEGER NOT NULL,
    nama_cabang VARCHAR(100) NOT NULL,
    alamat TEXT,
    kota VARCHAR(50),
    provinsi VARCHAR(50),
    kode_pos VARCHAR(10),
    telepon VARCHAR(20),
    email VARCHAR(100),
    pic_nama VARCHAR(100),
    pic_email VARCHAR(100),
    pic_telepon VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_brand) REFERENCES brand(id_brand)
);

-- Tabel Outlet
CREATE TABLE IF NOT EXISTS outlet (
    id_outlet SERIAL PRIMARY KEY,
    id_cabang INTEGER NOT NULL,
    nama_outlet VARCHAR(100) NOT NULL,
    alamat TEXT,
    kota VARCHAR(50),
    provinsi VARCHAR(50),
    kode_pos VARCHAR(10),
    telepon VARCHAR(20),
    email VARCHAR(100),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    jam_buka TIME,
    jam_tutup TIME,
    pic_nama VARCHAR(100),
    pic_email VARCHAR(100),
    pic_telepon VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_cabang) REFERENCES cabang(id_cabang)
);

-- Index untuk optimasi query
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_cabang_brand ON cabang(id_brand);
CREATE INDEX IF NOT EXISTS idx_outlet_cabang ON outlet(id_cabang);

-- Data user admin contoh (password: admin123)
-- INSERT INTO users (email, password, name, role) VALUES 
-- ('admin@laundry.com', '$2a$10$5.A.xcF3M9B/Pg9P7x4.3O3pF8K6N7.8.9.0.1.2.3.4.5.6.7.8.9.0.1.2', 'Admin Laundry', 'admin');