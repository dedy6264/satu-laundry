-- Script to manage service categories

-- Create kategori_layanan table to store service categories
CREATE TABLE IF NOT EXISTS kategori_layanan (
    id_kategori SERIAL PRIMARY KEY,
    nama_kategori VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes for faster queries
CREATE INDEX IF NOT EXISTS idx_kategori_layanan_nama ON kategori_layanan(nama_kategori);

-- CRUD Operations for Service Categories

-- CREATE - Add new service category
-- INSERT INTO kategori_layanan (nama_kategori, deskripsi) 
-- VALUES ($1, $2);

-- READ - Get service category by ID
-- SELECT id_kategori, nama_kategori, deskripsi, created_at, updated_at
-- FROM kategori_layanan
-- WHERE id_kategori = $1;

-- READ - Get all service categories with pagination
-- SELECT id_kategori, nama_kategori, deskripsi, created_at, updated_at
-- FROM kategori_layanan
-- ORDER BY id_kategori
-- LIMIT $1 OFFSET $2;

-- UPDATE - Update service category
-- UPDATE kategori_layanan 
-- SET nama_kategori = $1, deskripsi = $2, updated_at = NOW()
-- WHERE id_kategori = $3;

-- DELETE - Remove service category
-- DELETE FROM kategori_layanan WHERE id_kategori = $1;