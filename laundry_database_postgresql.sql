-- Database Sistem Laundry "Satu Laundry" - PostgreSQL Version
-- Struktur database untuk sistem multi-brand, multi-cabang, multi-outlet

-- Membuat database
-- CREATE DATABASE satu_laundry;
-- \c satu_laundry;

-- Extension untuk UUID (jika diperlukan)
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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

-- Tabel Brand
CREATE TABLE brand (
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
CREATE TABLE cabang (
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
CREATE TABLE outlet (
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
CREATE TABLE pelanggan (
    id_pelanggan SERIAL PRIMARY KEY,
    nomor_hp VARCHAR(20) UNIQUE NOT NULL,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    alamat TEXT,
    tanggal_lahir DATE,
    jenis_kelamin VARCHAR(10) CHECK (jenis_kelamin IN ('L', 'P')),
    poin_reward INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Pegawai
CREATE TABLE pegawai (
    id_pegawai SERIAL PRIMARY KEY,
    id_outlet INTEGER NOT NULL,
    nik VARCHAR(20) UNIQUE,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    telepon VARCHAR(20),
    alamat TEXT,
    tanggal_lahir DATE,
    jenis_kelamin VARCHAR(10) CHECK (jenis_kelamin IN ('L', 'P')),
    posisi VARCHAR(50),
    gaji DECIMAL(15, 2),
    tanggal_masuk DATE,
    status VARCHAR(20) DEFAULT 'aktif' CHECK (status IN ('aktif', 'tidak aktif')),
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


-- Tabel Paket Layanan
CREATE TABLE paket_layanan (
    id_layanan SERIAL PRIMARY KEY,
    nama_layanan VARCHAR(100) NOT NULL,
    id_kategori INTEGER NOT NULL,
    id_brand INTEGER NOT NULL,
    deskripsi TEXT,
    harga_per_kg DECIMAL(10, 2),
    durasi_pengerjaan INTEGER,
    satuan_durasi VARCHAR(10) DEFAULT 'hari' CHECK (satuan_durasi IN ('jam', 'hari')),
    kategori VARCHAR(20) DEFAULT 'kiloan' CHECK (kategori IN ('kiloan', 'satuan')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_brand) REFERENCES brand(id_brand),
    FOREIGN KEY (id_kategori) REFERENCES kategori_layanan(id_kategori)

);

-- Tabel Transaksi
CREATE TABLE transaksi (
    id_transaksi SERIAL PRIMARY KEY,
    id_pelanggan INTEGER NOT NULL,
    id_outlet INTEGER NOT NULL,
    id_pegawai INTEGER,
    nomor_invoice VARCHAR(50) UNIQUE,
    tanggal_masuk TIMESTAMP,
    tanggal_selesai TIMESTAMP,
    tanggal_diambil TIMESTAMP,
    berat_laundry DECIMAL(5, 2),
    total_harga DECIMAL(15, 2),
    uang_bayar DECIMAL(15, 2),
    uang_kembalian DECIMAL(15, 2),
    status_transaksi VARCHAR(20) DEFAULT 'diterima' CHECK (status_transaksi IN ('diterima', 'diproses', 'selesai', 'diambil')),
    status_pembayaran VARCHAR(20) DEFAULT 'belum lunas' CHECK (status_pembayaran IN ('lunas', 'belum lunas')),
    metode_pembayaran VARCHAR(20) DEFAULT 'tunai' CHECK (metode_pembayaran IN ('tunai', 'transfer', 'e-wallet')),
    catatan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan),
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet),
    FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai)
);

-- Tabel Detail Transaksi
CREATE TABLE detail_transaksi (
    id_detail SERIAL PRIMARY KEY,
    id_transaksi INTEGER NOT NULL,
    id_layanan INTEGER NOT NULL,
    kuantitas DECIMAL(5, 2),
    harga_satuan DECIMAL(10, 2),
    subtotal DECIMAL(15, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi) ON DELETE CASCADE,
    FOREIGN KEY (id_layanan) REFERENCES paket_layanan(id_layanan)
);

-- Tabel Inventaris
CREATE TABLE inventaris (
    id_inventaris SERIAL PRIMARY KEY,
    id_outlet INTEGER NOT NULL,
    nama_barang VARCHAR(100) NOT NULL,
    kategori VARCHAR(50),
    jumlah_stok INTEGER,
    satuan VARCHAR(20),
    harga_beli DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet)
);

-- Tabel Pengeluaran
CREATE TABLE pengeluaran (
    id_pengeluaran SERIAL PRIMARY KEY,
    id_outlet INTEGER NOT NULL,
    keterangan TEXT,
    jumlah DECIMAL(15, 2),
    tanggal_pengeluaran DATE,
    kategori VARCHAR(20) CHECK (kategori IN ('operasional', 'gaji', 'perlengkapan', 'lainnya')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet)
);

-- Tabel Pelanggan Outlet (untuk mencatat outlet tempat pelanggan terdaftar)
CREATE TABLE pelanggan_outlet (
    id_pelanggan_outlet SERIAL PRIMARY KEY,
    id_pelanggan INTEGER NOT NULL,
    id_outlet INTEGER NOT NULL,
    tanggal_terdaftar TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan) ON DELETE CASCADE,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet) ON DELETE CASCADE,
    UNIQUE (id_pelanggan, id_outlet)
);

-- Tabel Presensi Pegawai
CREATE TABLE presensi (
    id_presensi SERIAL PRIMARY KEY,
    id_pegawai INTEGER NOT NULL,
    id_outlet INTEGER NOT NULL,
    tanggal DATE NOT NULL,
    waktu_masuk TIME,
    waktu_keluar TIME,
    foto_masuk VARCHAR(255),
    foto_keluar VARCHAR(255),
    status_presensi VARCHAR(10) DEFAULT 'alfa' CHECK (status_presensi IN ('hadir', 'izin', 'sakit', 'alfa')),
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai) ON DELETE CASCADE,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet),
    UNIQUE (id_pegawai, tanggal)
);

-- Tabel History Status Transaksi
CREATE TABLE history_status_transaksi (
    id_history SERIAL PRIMARY KEY,
    id_transaksi INTEGER NOT NULL,
    status_lama VARCHAR(20) CHECK (status_lama IN ('diterima', 'diproses', 'selesai', 'diambil')),
    status_baru VARCHAR(20) NOT NULL CHECK (status_baru IN ('diterima', 'diproses', 'selesai', 'diambil')),
    waktu_perubahan TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi) ON DELETE CASCADE
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
CREATE TABLE IF NOT EXISTS employee_access (
    id_access SERIAL PRIMARY KEY,
    id_pegawai INTEGER NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'karyawan',
    is_active BOOLEAN DEFAULT true,
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai) ON DELETE CASCADE
);


-- Index untuk optimasi query
CREATE INDEX idx_transaksi_pelanggan ON transaksi(id_pelanggan);
CREATE INDEX idx_transaksi_outlet ON transaksi(id_outlet);
CREATE INDEX idx_transaksi_tanggal ON transaksi(tanggal_masuk);
CREATE INDEX idx_transaksi_status ON transaksi(status_transaksi);
CREATE INDEX idx_pegawai_outlet ON pegawai(id_outlet);
CREATE INDEX idx_outlet_cabang ON outlet(id_cabang);
CREATE INDEX idx_cabang_brand ON cabang(id_brand);
CREATE INDEX idx_paket_brand ON paket_layanan(id_brand);
CREATE INDEX idx_layanan_kategori ON paket_layanan(id_kategori);
CREATE INDEX idx_pembayaran_transaksi ON pembayaran(id_transaksi);
-- Add indexes for faster queries
CREATE INDEX IF NOT EXISTS idx_employee_access_username ON employee_access(username);
CREATE INDEX IF NOT EXISTS idx_employee_access_pegawai ON employee_access(id_pegawai);
CREATE INDEX IF NOT EXISTS idx_employee_access_active ON employee_access(is_active);


-- Function untuk mengenerate nomor invoice otomatis
CREATE OR REPLACE FUNCTION generate_invoice_number()
RETURNS TEXT AS $$
BEGIN
    RETURN 'INV' || TO_CHAR(NOW(), 'YYYYMMDD') || LPAD(NEXTVAL('invoice_seq')::TEXT, 6, '0');
END;
$$ LANGUAGE plpgsql;

-- Sequence untuk nomor invoice
CREATE SEQUENCE IF NOT EXISTS invoice_seq START 1;

-- Trigger function untuk mencatat history status transaksi
CREATE OR REPLACE FUNCTION trigger_history_status_transaksi()
RETURNS TRIGGER AS $$
BEGIN
    -- Mencatat history jika status transaksi berubah
    IF OLD.status_transaksi IS DISTINCT FROM NEW.status_transaksi THEN
        INSERT INTO history_status_transaksi (id_transaksi, status_lama, status_baru)
        VALUES (NEW.id_transaksi, OLD.status_transaksi, NEW.status_transaksi);
    END IF;
    
    -- Update status pembayaran jika uang bayar mencukupi total harga
    IF NEW.uang_bayar >= NEW.total_harga AND NEW.status_pembayaran = 'belum lunas' THEN
        NEW.status_pembayaran = 'lunas';
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger untuk transaksi
CREATE TRIGGER trigger_transaksi_update
    BEFORE UPDATE ON transaksi
    FOR EACH ROW
    EXECUTE FUNCTION trigger_history_status_transaksi();

-- Function untuk pendaftaran pelanggan baru
CREATE OR REPLACE FUNCTION sp_daftar_pelanggan_baru(
    p_nomor_hp VARCHAR(20),
    p_nama_lengkap VARCHAR(100),
    p_email VARCHAR(100),
    p_alamat TEXT,
    p_tanggal_lahir DATE,
    p_jenis_kelamin VARCHAR(10),
    p_id_outlet INTEGER
)
RETURNS TABLE(id_pelanggan INTEGER, status VARCHAR(50))
AS $$
DECLARE
    v_count INTEGER;
    v_id_pelanggan INTEGER;
    v_status VARCHAR(50);
BEGIN
    -- Cek apakah pelanggan sudah terdaftar berdasarkan nomor HP
    SELECT COUNT(*) INTO v_count FROM pelanggan WHERE nomor_hp = p_nomor_hp;
    
    IF v_count = 0 THEN
        -- Jika belum terdaftar, tambahkan sebagai pelanggan baru
        INSERT INTO pelanggan (nomor_hp, nama_lengkap, email, alamat, tanggal_lahir, jenis_kelamin)
        VALUES (p_nomor_hp, p_nama_lengkap, p_email, p_alamat, p_tanggal_lahir, p_jenis_kelamin)
        RETURNING pelanggan.id_pelanggan INTO v_id_pelanggan;
        
        v_status = 'baru';
    ELSE
        -- Jika sudah terdaftar, gunakan data yang sudah ada
        SELECT id_pelanggan INTO v_id_pelanggan FROM pelanggan WHERE nomor_hp = p_nomor_hp;
        v_status = 'exist';
    END IF;
    
    -- Cek apakah pelanggan sudah terdaftar di outlet ini
    SELECT COUNT(*) INTO v_count FROM pelanggan_outlet 
    WHERE id_pelanggan = v_id_pelanggan AND id_outlet = p_id_outlet;
    
    -- Jika belum terdaftar di outlet ini, tambahkan
    IF v_count = 0 THEN
        INSERT INTO pelanggan_outlet (id_pelanggan, id_outlet)
        VALUES (v_id_pelanggan, p_id_outlet);
    END IF;
    
    RETURN QUERY SELECT v_id_pelanggan, v_status;
END;
$$ LANGUAGE plpgsql;

-- Function untuk mengecek status pelanggan
CREATE OR REPLACE FUNCTION sp_cek_status_pelanggan(p_nomor_hp VARCHAR(20))
RETURNS TABLE(id_pelanggan INTEGER, status VARCHAR(50), nama_lengkap VARCHAR(100))
AS $$
DECLARE
    v_count INTEGER;
    v_id_pelanggan INTEGER;
    v_nama_lengkap VARCHAR(100);
    v_status VARCHAR(50);
BEGIN
    -- Cek apakah pelanggan terdaftar
    SELECT COUNT(*) INTO v_count FROM pelanggan WHERE nomor_hp = p_nomor_hp;
    
    IF v_count > 0 THEN
        SELECT pelanggan.id_pelanggan, pelanggan.nama_lengkap INTO v_id_pelanggan, v_nama_lengkap 
        FROM pelanggan WHERE nomor_hp = p_nomor_hp;
        v_status = 'terdaftar';
    ELSE
        v_id_pelanggan = 0;
        v_nama_lengkap = '';
        v_status = 'belum_terdaftar';
    END IF;
    
    RETURN QUERY SELECT v_id_pelanggan, v_status, v_nama_lengkap;
END;
$$ LANGUAGE plpgsql;

-- Function untuk mencatat presensi pegawai
CREATE OR REPLACE FUNCTION sp_catat_presensi(
    p_nik VARCHAR(20),
    p_id_outlet INTEGER,
    p_tanggal DATE,
    p_waktu_masuk TIME,
    p_foto_masuk VARCHAR(255),
    p_status_presensi VARCHAR(10),
    p_keterangan TEXT
)
RETURNS VARCHAR(50)
AS $$
DECLARE
    v_id_pegawai INTEGER;
    v_count INTEGER;
    v_status VARCHAR(50);
BEGIN
    -- Cek apakah pegawai dengan NIK tersebut ada
    SELECT COUNT(*) INTO v_count FROM pegawai WHERE nik = p_nik;
    
    IF v_count = 0 THEN
        v_status = 'pegawai_tidak_ditemukan';
    ELSE
        -- Dapatkan id_pegawai
        SELECT id_pegawai INTO v_id_pegawai FROM pegawai WHERE nik = p_nik;
        
        -- Cek apakah sudah ada presensi untuk tanggal tersebut
        SELECT COUNT(*) INTO v_count FROM presensi 
        WHERE id_pegawai = v_id_pegawai AND tanggal = p_tanggal;
        
        IF v_count = 0 THEN
            -- Jika belum ada, tambahkan presensi baru
            INSERT INTO presensi (id_pegawai, id_outlet, tanggal, waktu_masuk, foto_masuk, status_presensi, keterangan)
            VALUES (v_id_pegawai, p_id_outlet, p_tanggal, p_waktu_masuk, p_foto_masuk, p_status_presensi, p_keterangan);
            v_status = 'presensi_tercatat';
        ELSE
            -- Jika sudah ada, update presensi
            UPDATE presensi 
            SET waktu_masuk = p_waktu_masuk, foto_masuk = p_foto_masuk, status_presensi = p_status_presensi, keterangan = p_keterangan
            WHERE id_pegawai = v_id_pegawai AND tanggal = p_tanggal;
            v_status = 'presensi_diperbarui';
        END IF;
    END IF;
    
    RETURN v_status;
END;
$$ LANGUAGE plpgsql;

-- Function untuk mencatat waktu keluar pegawai
CREATE OR REPLACE FUNCTION sp_catat_keluar(
    p_nik VARCHAR(20),
    p_tanggal DATE,
    p_waktu_keluar TIME,
    p_foto_keluar VARCHAR(255)
)
RETURNS VARCHAR(50)
AS $$
DECLARE
    v_id_pegawai INTEGER;
    v_count INTEGER;
    v_status VARCHAR(50);
BEGIN
    -- Cek apakah pegawai dengan NIK tersebut ada
    SELECT COUNT(*) INTO v_count FROM pegawai WHERE nik = p_nik;
    
    IF v_count = 0 THEN
        v_status = 'pegawai_tidak_ditemukan';
    ELSE
        -- Dapatkan id_pegawai
        SELECT id_pegawai INTO v_id_pegawai FROM pegawai WHERE nik = p_nik;
        
        -- Cek apakah sudah ada presensi untuk tanggal tersebut
        SELECT COUNT(*) INTO v_count FROM presensi 
        WHERE id_pegawai = v_id_pegawai AND tanggal = p_tanggal;
        
        IF v_count = 0 THEN
            v_status = 'presensi_tidak_ditemukan';
        ELSE
            -- Update waktu keluar
            UPDATE presensi 
            SET waktu_keluar = p_waktu_keluar, foto_keluar = p_foto_keluar
            WHERE id_pegawai = v_id_pegawai AND tanggal = p_tanggal;
            v_status = 'waktu_keluar_tercatat';
        END IF;
    END IF;
    
    RETURN v_status;
END;
$$ LANGUAGE plpgsql;

-- View untuk laporan transaksi
CREATE VIEW v_laporan_transaksi AS
SELECT 
    t.id_transaksi,
    t.nomor_invoice,
    p.nomor_hp,
    p.nama_lengkap AS nama_pelanggan,
    b.nama_brand,
    c.nama_cabang,
    o.nama_outlet,
    o.pic_nama AS pic_outlet,
    pg.nama_lengkap AS nama_pegawai,
    t.tanggal_masuk,
    t.tanggal_selesai,
    t.tanggal_diambil,
    t.berat_laundry,
    t.total_harga,
    t.status_transaksi,
    t.status_pembayaran,
    t.metode_pembayaran
FROM transaksi t
JOIN pelanggan p ON t.id_pelanggan = p.id_pelanggan
JOIN outlet o ON t.id_outlet = o.id_outlet
JOIN cabang c ON o.id_cabang = c.id_cabang
JOIN brand b ON c.id_brand = b.id_brand
LEFT JOIN pegawai pg ON t.id_pegawai = pg.id_pegawai;

-- View untuk struktur brand-cabang-outlet
CREATE VIEW v_struktur_organisasi AS
SELECT 
    b.nama_brand,
    b.pic_nama AS brand_pic,
    b.pic_email AS brand_email,
    b.pic_telepon AS brand_telepon,
    c.nama_cabang,
    c.pic_nama AS cabang_pic,
    c.pic_email AS cabang_email,
    c.pic_telepon AS cabang_telepon,
    o.nama_outlet,
    o.pic_nama AS outlet_pic,
    o.pic_email AS outlet_email,
    o.pic_telepon AS outlet_telepon,
    o.alamat,
    o.telepon
FROM brand b
JOIN cabang c ON b.id_brand = c.id_brand
JOIN outlet o ON c.id_cabang = o.id_cabang;

-- View untuk informasi PIC
CREATE VIEW v_info_pic AS
SELECT 
    b.nama_brand,
    b.pic_nama AS brand_pic,
    b.pic_email AS brand_email,
    b.pic_telepon AS brand_telepon,
    c.nama_cabang,
    c.pic_nama AS cabang_pic,
    c.pic_email AS cabang_email,
    c.pic_telepon AS cabang_telepon,
    o.nama_outlet,
    o.pic_nama AS outlet_pic,
    o.pic_email AS outlet_email,
    o.pic_telepon AS outlet_telepon
FROM brand b
JOIN cabang c ON b.id_brand = c.id_brand
JOIN outlet o ON c.id_cabang = o.id_cabang;

-- View untuk outlet tempat pelanggan terdaftar
CREATE VIEW v_pelanggan_outlet AS
SELECT 
    p.id_pelanggan,
    p.nomor_hp,
    p.nama_lengkap AS nama_pelanggan,
    p.email AS email_pelanggan,
    b.nama_brand,
    c.nama_cabang,
    o.nama_outlet,
    po.tanggal_terdaftar
FROM pelanggan p
JOIN pelanggan_outlet po ON p.id_pelanggan = po.id_pelanggan
JOIN outlet o ON po.id_outlet = o.id_outlet
JOIN cabang c ON o.id_cabang = c.id_cabang
JOIN brand b ON c.id_brand = b.id_brand
ORDER BY p.nama_lengkap, po.tanggal_terdaftar;

-- View untuk riwayat transaksi pelanggan di semua outlet
CREATE VIEW v_riwayat_transaksi_pelanggan AS
SELECT 
    p.id_pelanggan,
    p.nomor_hp,
    p.nama_lengkap AS nama_pelanggan,
    b.nama_brand,
    c.nama_cabang,
    o.nama_outlet,
    t.nomor_invoice,
    t.tanggal_masuk,
    t.tanggal_selesai,
    t.tanggal_diambil,
    t.berat_laundry,
    t.total_harga,
    t.status_transaksi,
    t.status_pembayaran
FROM pelanggan p
JOIN transaksi t ON p.id_pelanggan = t.id_pelanggan
JOIN outlet o ON t.id_outlet = o.id_outlet
JOIN cabang c ON o.id_cabang = c.id_cabang
JOIN brand b ON c.id_brand = b.id_brand
ORDER BY p.nama_lengkap, t.tanggal_masuk DESC;

-- View untuk laporan presensi
CREATE VIEW v_laporan_presensi AS
SELECT 
    pr.id_presensi,
    pg.id_pegawai,
    pg.nik,
    pg.nama_lengkap AS nama_pegawai,
    b.nama_brand,
    c.nama_cabang,
    o.nama_outlet,
    pr.tanggal,
    pr.waktu_masuk,
    pr.waktu_keluar,
    pr.foto_masuk,
    pr.foto_keluar,
    pr.status_presensi,
    pr.keterangan,
    (pr.waktu_keluar - pr.waktu_masuk) AS durasi_kerja
FROM presensi pr
JOIN pegawai pg ON pr.id_pegawai = pg.id_pegawai
JOIN outlet o ON pr.id_outlet = o.id_outlet
JOIN cabang c ON o.id_cabang = c.id_cabang
JOIN brand b ON c.id_brand = b.id_brand
ORDER BY pr.tanggal DESC, pr.waktu_masuk DESC;

-- View untuk history status transaksi
CREATE VIEW v_history_status_transaksi AS
SELECT 
    h.id_history,
    t.nomor_invoice,
    p.nama_lengkap AS nama_pelanggan,
    h.status_lama,
    h.status_baru,
    h.waktu_perubahan,
    h.keterangan
FROM history_status_transaksi h
JOIN transaksi t ON h.id_transaksi = t.id_transaksi
JOIN pelanggan p ON t.id_pelanggan = p.id_pelanggan
ORDER BY h.waktu_perubahan DESC;