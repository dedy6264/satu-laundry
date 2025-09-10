-- Database Sistem Laundry "Satu Laundry"
-- Struktur database untuk sistem multi-brand, multi-cabang, multi-outlet

-- Membuat database
CREATE DATABASE IF NOT EXISTS satu_laundry;
USE satu_laundry;

-- Tabel Brand
CREATE TABLE brand (
    id_brand INT PRIMARY KEY AUTO_INCREMENT,
    nama_brand VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    tanggal_berdiri DATE,
    pic_nama VARCHAR(100),
    pic_email VARCHAR(100),
    pic_telepon VARCHAR(20),
    logo_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabel Cabang
CREATE TABLE cabang (
    id_cabang INT PRIMARY KEY AUTO_INCREMENT,
    id_brand INT NOT NULL,
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
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_brand) REFERENCES brand(id_brand)
);

-- Tabel Outlet
CREATE TABLE outlet (
    id_outlet INT PRIMARY KEY AUTO_INCREMENT,
    id_cabang INT NOT NULL,
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
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_cabang) REFERENCES cabang(id_cabang)
);

-- Tabel Pelanggan
CREATE TABLE pelanggan (
    id_pelanggan INT PRIMARY KEY AUTO_INCREMENT,
    nomor_hp VARCHAR(20) UNIQUE NOT NULL,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    alamat TEXT,
    tanggal_lahir DATE,
    jenis_kelamin ENUM('L', 'P'),
    poin_reward INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabel Pelanggan Outlet (untuk mencatat outlet tempat pelanggan terdaftar)
CREATE TABLE pelanggan_outlet (
    id_pelanggan_outlet INT PRIMARY KEY AUTO_INCREMENT,
    id_pelanggan INT NOT NULL,
    id_outlet INT NOT NULL,
    tanggal_terdaftar TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan) ON DELETE CASCADE,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet) ON DELETE CASCADE,
    UNIQUE KEY unique_pelanggan_outlet (id_pelanggan, id_outlet)
);

-- Tabel Pegawai
CREATE TABLE pegawai (
    id_pegawai INT PRIMARY KEY AUTO_INCREMENT,
    id_outlet INT NOT NULL,
    nik VARCHAR(20) UNIQUE,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    telepon VARCHAR(20),
    alamat TEXT,
    tanggal_lahir DATE,
    jenis_kelamin ENUM('L', 'P'),
    posisi VARCHAR(50),
    gaji DECIMAL(15, 2),
    tanggal_masuk DATE,
    status ENUM('aktif', 'tidak aktif') DEFAULT 'aktif',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet)
);

-- Tabel Paket Layanan
CREATE TABLE paket_layanan (
    id_paket INT PRIMARY KEY AUTO_INCREMENT,
    id_brand INT NOT NULL,
    nama_paket VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    harga_per_kg DECIMAL(10, 2),
    durasi_pengerjaan INT, -- dalam jam
    satuan_durasi ENUM('jam', 'hari') DEFAULT 'hari',
    kategori ENUM('kiloan', 'satuan') DEFAULT 'kiloan',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_brand) REFERENCES brand(id_brand)
);

-- Tabel Transaksi
CREATE TABLE transaksi (
    id_transaksi INT PRIMARY KEY AUTO_INCREMENT,
    id_pelanggan INT NOT NULL,
    id_outlet INT NOT NULL,
    id_pegawai INT,
    nomor_invoice VARCHAR(50) UNIQUE,
    tanggal_masuk DATETIME,
    tanggal_selesai DATETIME,
    tanggal_diambil DATETIME,
    berat_laundry DECIMAL(5, 2),
    total_harga DECIMAL(15, 2),
    uang_bayar DECIMAL(15, 2),
    uang_kembalian DECIMAL(15, 2),
    status_transaksi ENUM('diterima', 'diproses', 'selesai', 'diambil') DEFAULT 'diterima',
    status_pembayaran ENUM('lunas', 'belum lunas') DEFAULT 'belum lunas',
    metode_pembayaran ENUM('tunai', 'transfer', 'e-wallet') DEFAULT 'tunai',
    catatan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan),
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet),
    FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai)
);

-- Tabel Detail Transaksi
CREATE TABLE detail_transaksi (
    id_detail INT PRIMARY KEY AUTO_INCREMENT,
    id_transaksi INT NOT NULL,
    id_paket INT NOT NULL,
    kuantitas DECIMAL(5, 2),
    harga_satuan DECIMAL(10, 2),
    subtotal DECIMAL(15, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi) ON DELETE CASCADE,
    FOREIGN KEY (id_paket) REFERENCES paket_layanan(id_paket)
);

-- Tabel History Status Transaksi
CREATE TABLE history_status_transaksi (
    id_history INT PRIMARY KEY AUTO_INCREMENT,
    id_transaksi INT NOT NULL,
    status_lama ENUM('diterima', 'diproses', 'selesai', 'diambil'),
    status_baru ENUM('diterima', 'diproses', 'selesai', 'diambil') NOT NULL,
    waktu_perubahan TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi) ON DELETE CASCADE
);

-- Tabel Inventaris
CREATE TABLE inventaris (
    id_inventaris INT PRIMARY KEY AUTO_INCREMENT,
    id_outlet INT NOT NULL,
    nama_barang VARCHAR(100) NOT NULL,
    kategori VARCHAR(50),
    jumlah_stok INT,
    satuan VARCHAR(20),
    harga_beli DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet)
);

-- Tabel Pengeluaran
CREATE TABLE pengeluaran (
    id_pengeluaran INT PRIMARY KEY AUTO_INCREMENT,
    id_outlet INT NOT NULL,
    keterangan TEXT,
    jumlah DECIMAL(15, 2),
    tanggal_pengeluaran DATE,
    kategori ENUM('operasional', 'gaji', 'perlengkapan', 'lainnya'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet)
);

-- Tabel Presensi Pegawai
CREATE TABLE presensi (
    id_presensi INT PRIMARY KEY AUTO_INCREMENT,
    id_pegawai INT NOT NULL,
    id_outlet INT NOT NULL,
    tanggal DATE NOT NULL,
    waktu_masuk TIME,
    waktu_keluar TIME,
    foto_masuk VARCHAR(255),
    foto_keluar VARCHAR(255),
    status_presensi ENUM('hadir', 'izin', 'sakit', 'alfa') DEFAULT 'alfa',
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai) ON DELETE CASCADE,
    FOREIGN KEY (id_outlet) REFERENCES outlet(id_outlet),
    UNIQUE KEY unique_presensi (id_pegawai, tanggal)
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

-- Trigger untuk mengenerate nomor invoice otomatis
DELIMITER //
CREATE TRIGGER before_transaksi_insert
BEFORE INSERT ON transaksi
FOR EACH ROW
BEGIN
    IF NEW.nomor_invoice IS NULL OR NEW.nomor_invoice = '' THEN
        SET NEW.nomor_invoice = CONCAT('INV', DATE_FORMAT(NOW(), '%Y%m%d'), LPAD(LAST_INSERT_ID() + 1, 6, '0'));
    END IF;
END//

-- Trigger untuk mencatat history status transaksi
CREATE TRIGGER after_transaksi_update
AFTER UPDATE ON transaksi
FOR EACH ROW
BEGIN
    -- Mencatat history jika status transaksi berubah
    IF OLD.status_transaksi != NEW.status_transaksi THEN
        INSERT INTO history_status_transaksi (id_transaksi, status_lama, status_baru)
        VALUES (NEW.id_transaksi, OLD.status_transaksi, NEW.status_transaksi);
    END IF;
    
    -- Update status pembayaran jika uang bayar mencukupi total harga
    IF NEW.uang_bayar >= NEW.total_harga AND NEW.status_pembayaran = 'belum lunas' THEN
        UPDATE transaksi SET status_pembayaran = 'lunas' WHERE id_transaksi = NEW.id_transaksi;
    END IF;
END//
DELIMITER ;

-- View untuk laporan transaksi
CREATE VIEW v_laporan_transaksi AS
SELECT 
    t.id_transaksi,
    t.nomor_invoice,
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
    TIMEDIFF(pr.waktu_keluar, pr.waktu_masuk) AS durasi_kerja
FROM presensi pr
JOIN pegawai pg ON pr.id_pegawai = pg.id_pegawai
JOIN outlet o ON pr.id_outlet = o.id_outlet
JOIN cabang c ON o.id_cabang = c.id_cabang
JOIN brand b ON c.id_brand = b.id_brand
ORDER BY pr.tanggal DESC, pr.waktu_masuk DESC;

-- Stored Procedure untuk pendaftaran pelanggan otomatis
DELIMITER //
CREATE PROCEDURE sp_daftar_pelanggan_baru(
    IN p_nomor_hp VARCHAR(20),
    IN p_nama_lengkap VARCHAR(100),
    IN p_email VARCHAR(100),
    IN p_alamat TEXT,
    IN p_tanggal_lahir DATE,
    IN p_jenis_kelamin ENUM('L', 'P'),
    IN p_id_outlet INT,
    OUT p_id_pelanggan INT,
    OUT p_status VARCHAR(50)
)
BEGIN
    DECLARE v_count INT DEFAULT 0;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;
    
    START TRANSACTION;
    
    -- Cek apakah pelanggan sudah terdaftar berdasarkan nomor HP
    SELECT COUNT(*) INTO v_count FROM pelanggan WHERE nomor_hp = p_nomor_hp;
    
    IF v_count = 0 THEN
        -- Jika belum terdaftar, tambahkan sebagai pelanggan baru
        INSERT INTO pelanggan (nomor_hp, nama_lengkap, email, alamat, tanggal_lahir, jenis_kelamin)
        VALUES (p_nomor_hp, p_nama_lengkap, p_email, p_alamat, p_tanggal_lahir, p_jenis_kelamin);
        
        SET p_id_pelanggan = LAST_INSERT_ID();
        SET p_status = 'baru';
    ELSE
        -- Jika sudah terdaftar, gunakan data yang sudah ada
        SELECT id_pelanggan INTO p_id_pelanggan FROM pelanggan WHERE nomor_hp = p_nomor_hp;
        SET p_status = 'exist';
    END IF;
    
    -- Cek apakah pelanggan sudah terdaftar di outlet ini
    SELECT COUNT(*) INTO v_count FROM pelanggan_outlet 
    WHERE id_pelanggan = p_id_pelanggan AND id_outlet = p_id_outlet;
    
    -- Jika belum terdaftar di outlet ini, tambahkan
    IF v_count = 0 THEN
        INSERT INTO pelanggan_outlet (id_pelanggan, id_outlet)
        VALUES (p_id_pelanggan, p_id_outlet);
    END IF;
    
    COMMIT;
END //

-- Stored Procedure untuk mengecek status pelanggan
CREATE PROCEDURE sp_cek_status_pelanggan(
    IN p_nomor_hp VARCHAR(20),
    OUT p_id_pelanggan INT,
    OUT p_status VARCHAR(50),
    OUT p_nama_lengkap VARCHAR(100)
)
BEGIN
    DECLARE v_count INT DEFAULT 0;
    
    -- Cek apakah pelanggan terdaftar
    SELECT COUNT(*) INTO v_count FROM pelanggan WHERE nomor_hp = p_nomor_hp;
    
    IF v_count > 0 THEN
        SELECT id_pelanggan, nama_lengkap INTO p_id_pelanggan, p_nama_lengkap 
        FROM pelanggan WHERE nomor_hp = p_nomor_hp;
        SET p_status = 'terdaftar';
    ELSE
        SET p_id_pelanggan = 0;
        SET p_nama_lengkap = '';
        SET p_status = 'belum_terdaftar';
    END IF;
END //

-- Stored Procedure untuk mencatat presensi pegawai
CREATE PROCEDURE sp_catat_presensi(
    IN p_nik VARCHAR(20),
    IN p_id_outlet INT,
    IN p_tanggal DATE,
    IN p_waktu_masuk TIME,
    IN p_foto_masuk VARCHAR(255),
    IN p_status_presensi ENUM('hadir', 'izin', 'sakit', 'alfa'),
    IN p_keterangan TEXT,
    OUT p_status VARCHAR(50)
)
BEGIN
    DECLARE v_id_pegawai INT DEFAULT 0;
    DECLARE v_count INT DEFAULT 0;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;
    
    START TRANSACTION;
    
    -- Cek apakah pegawai dengan NIK tersebut ada
    SELECT COUNT(*) INTO v_count FROM pegawai WHERE nik = p_nik;
    
    IF v_count = 0 THEN
        SET p_status = 'pegawai_tidak_ditemukan';
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
            SET p_status = 'presensi_tercatat';
        ELSE
            -- Jika sudah ada, update presensi
            UPDATE presensi 
            SET waktu_masuk = p_waktu_masuk, foto_masuk = p_foto_masuk, status_presensi = p_status_presensi, keterangan = p_keterangan
            WHERE id_pegawai = v_id_pegawai AND tanggal = p_tanggal;
            SET p_status = 'presensi_diperbarui';
        END IF;
    END IF;
    
    COMMIT;
END //

-- Stored Procedure untuk mencatat waktu keluar pegawai
CREATE PROCEDURE sp_catat_keluar(
    IN p_nik VARCHAR(20),
    IN p_tanggal DATE,
    IN p_waktu_keluar TIME,
    IN p_foto_keluar VARCHAR(255),
    OUT p_status VARCHAR(50)
)
BEGIN
    DECLARE v_id_pegawai INT DEFAULT 0;
    DECLARE v_count INT DEFAULT 0;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;
    
    START TRANSACTION;
    
    -- Cek apakah pegawai dengan NIK tersebut ada
    SELECT COUNT(*) INTO v_count FROM pegawai WHERE nik = p_nik;
    
    IF v_count = 0 THEN
        SET p_status = 'pegawai_tidak_ditemukan';
    ELSE
        -- Dapatkan id_pegawai
        SELECT id_pegawai INTO v_id_pegawai FROM pegawai WHERE nik = p_nik;
        
        -- Cek apakah sudah ada presensi untuk tanggal tersebut
        SELECT COUNT(*) INTO v_count FROM presensi 
        WHERE id_pegawai = v_id_pegawai AND tanggal = p_tanggal;
        
        IF v_count = 0 THEN
            SET p_status = 'presensi_tidak_ditemukan';
        ELSE
            -- Update waktu keluar
            UPDATE presensi 
            SET waktu_keluar = p_waktu_keluar, foto_keluar = p_foto_keluar
            WHERE id_pegawai = v_id_pegawai AND tanggal = p_tanggal;
            SET p_status = 'waktu_keluar_tercatat';
        END IF;
    END IF;
    
    COMMIT;
END //

DELIMITER ;