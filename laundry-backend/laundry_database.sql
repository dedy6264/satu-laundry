-- Script untuk inisialisasi database PostgreSQL untuk Laundry Management System

-- Buat database (jalankan sebagai superuser)
-- CREATE DATABASE laundry_db;

-- Connect ke database
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

-- Tabel Pelanggan
CREATE TABLE IF NOT EXISTS pelanggan (
    id_pelanggan SERIAL PRIMARY KEY,
    id_outlet INTEGER NOT NULL,
    nama VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    telepon VARCHAR(20),
    alamat TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet)
);

-- Tabel Kategori Layanan
CREATE TABLE IF NOT EXISTS kategori_layanan (
    id_kategori SERIAL PRIMARY KEY,
    nama_kategori VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Layanan
CREATE TABLE IF NOT EXISTS layanan (
    id_layanan SERIAL PRIMARY KEY,
    id_kategori INTEGER NOT NULL,
    nama_layanan VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    harga DECIMAL(10, 2) NOT NULL,
    satuan VARCHAR(50) NOT NULL, -- kg, pcs, etc.
    estimasi_waktu INTERVAL, -- estimasi waktu pengerjaan
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_kategori) REFERENCES kategori_layanan(id_kategori)
);

-- Tabel Transaksi
CREATE TABLE IF NOT EXISTS transaksi (
    id_transaksi SERIAL PRIMARY KEY,
    id_pelanggan INTEGER NOT NULL,
    id_outlet INTEGER NOT NULL,
    nomor_invoice VARCHAR(50) UNIQUE NOT NULL,
    tanggal_masuk TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tanggal_selesai TIMESTAMP,
    tanggal_diambil TIMESTAMP,
    status VARCHAR(50) DEFAULT 'baru', -- baru, diproses, selesai, diambil
    total_biaya DECIMAL(12, 2) DEFAULT 0,
    dibayar DECIMAL(12, 2) DEFAULT 0,
    kembalian DECIMAL(12, 2) DEFAULT 0,
    catatan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan),
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet)
);

-- Tabel Detail Transaksi
CREATE TABLE IF NOT EXISTS detail_transaksi (
    id_detail SERIAL PRIMARY KEY,
    id_transaksi INTEGER NOT NULL,
    id_layanan INTEGER NOT NULL,
    jumlah DECIMAL(10, 2) NOT NULL,
    harga DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(12, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi) ON DELETE CASCADE,
    FOREIGN KEY (id_layanan) REFERENCES layanan(id_layanan)
);

-- Tabel Pembayaran
CREATE TABLE IF NOT EXISTS pembayaran (
    id_pembayaran SERIAL PRIMARY KEY,
    id_transaksi INTEGER NOT NULL,
    tanggal_bayar TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    jumlah_bayar DECIMAL(12, 2) NOT NULL,
    metode_bayar VARCHAR(50), -- tunai, transfer, dll
    catatan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi) ON DELETE CASCADE
);

-- Index untuk optimasi query
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_cabang_brand ON cabang(id_brand);
CREATE INDEX IF NOT EXISTS idx_outlet_cabang ON outlet(id_cabang);
CREATE INDEX IF NOT EXISTS idx_pelanggan_outlet ON pelanggan(id_outlet);
CREATE INDEX IF NOT EXISTS idx_layanan_kategori ON layanan(id_kategori);
CREATE INDEX IF NOT EXISTS idx_transaksi_pelanggan ON transaksi(id_pelanggan);
CREATE INDEX IF NOT EXISTS idx_transaksi_outlet ON transaksi(id_outlet);
CREATE INDEX IF NOT EXISTS idx_transaksi_status ON transaksi(status);
CREATE INDEX IF NOT EXISTS idx_transaksi_invoice ON transaksi(nomor_invoice);
CREATE INDEX IF NOT EXISTS idx_detail_transaksi ON detail_transaksi(id_transaksi);
CREATE INDEX IF NOT EXISTS idx_pembayaran_transaksi ON pembayaran(id_transaksi);

-- Data contoh

-- Data user admin contoh (password: admin123)
-- INSERT INTO users (email, password, name, role) VALUES 
-- ('admin@laundry.com', '$2a$10$5.A.xcF3M9B/Pg9P7x4.3O3pF8K6N7.8.9.0.1.2.3.4.5.6.7.8.9.0.1.2', 'Admin Laundry', 'admin');

-- Data brand contoh
-- INSERT INTO brand (nama_brand, deskripsi, pic_nama, pic_email, pic_telepon) VALUES
-- ('Satu Laundry', 'Jasa laundry profesional dengan layanan antar-jemput', 'John Doe', 'john@satu-laundry.com', '08123456789');

-- Data cabang contoh
-- INSERT INTO cabang (id_brand, nama_cabang, alamat, kota, provinsi, telepon, email, pic_nama, pic_email, pic_telepon) VALUES
-- (1, 'Cabang Surabaya', 'Jl. Raya Jemursari No. 123', 'Surabaya', 'Jawa Timur', '031123456', 'surabaya@satu-laundry.com', 'Jane Smith', 'jane@satu-laundry.com', '08234567890');

-- Data outlet contoh
-- INSERT INTO outlet (id_cabang, nama_outlet, alamat, kota, provinsi, telepon, email, latitude, longitude, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon) VALUES
-- (1, 'Outlet Jemursari', 'Jl. Raya Jemursari No. 123', 'Surabaya', 'Jawa Timur', '031123456', 'jemursari@satu-laundry.com', -7.32456789, 112.74567890, '08:00:00', '20:00:00', 'Robert Johnson', 'robert@satu-laundry.com', '08345678901');

-- Data kategori layanan contoh
-- INSERT INTO kategori_layanan (nama_kategori, deskripsi) VALUES
-- ('Cuci Kering', 'Layanan pencucian pakaian hingga kering'),
-- ('Cuci Setrika', 'Layanan pencucian pakaian hingga disetrika'),
-- ('Dry Cleaning', 'Layanan dry cleaning untuk pakaian khusus');

-- Data layanan contoh
-- INSERT INTO layanan (id_kategori, nama_layanan, deskripsi, harga, satuan, estimasi_waktu) VALUES
-- (1, 'Cuci Kering Reguler', 'Cuci kering reguler untuk pakaian biasa', 7000, 'kg', '1 day'),
-- (1, 'Cuci Kering Express', 'Cuci kering express selesai dalam 3 jam', 12000, 'kg', '3 hours'),
-- (2, 'Cuci Setrika Reguler', 'Cuci dan setrika reguler', 10000, 'kg', '1 day'),
-- (2, 'Cuci Setrika Express', 'Cuci dan setrika express selesai dalam 3 jam', 15000, 'kg', '3 hours'),
-- (3, 'Dry Cleaning Jaket', 'Dry cleaning untuk jaket kulit', 50000, 'pcs', '3 days'),
-- (3, 'Dry Cleaning Gaun', 'Dry cleaning untuk gaun panjang', 75000, 'pcs', '3 days');