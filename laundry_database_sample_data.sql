-- Data contoh untuk testing database sistem laundry

-- Data Brand
INSERT INTO brand (nama_brand, deskripsi, tanggal_berdiri, pic_nama, pic_email, pic_telepon, logo_url) VALUES
('Satu Laundry', 'Laundry profesional dengan layanan berkualitas tinggi', '2020-01-01', 'Budi Santoso', 'budi@satulaundry.com', '08111111111', '/images/logo_satulaundry.png'),
('Bersih Cepat', 'Laundry express dengan layanan 6 jam selesai', '2021-05-15', 'Ani Wijaya', 'ani@bersihcepat.com', '08222222222', '/images/logo_bersihcepat.png'),
('Wangi Segar', 'Laundry dengan aroma harum tahan lama', '2019-11-20', 'Candra Putra', 'candra@wangisegar.com', '08333333333', '/images/logo_wangisegar.png');

-- Data Cabang
INSERT INTO cabang (id_brand, nama_cabang, alamat, kota, provinsi, kode_pos, telepon, email, pic_nama, pic_email, pic_telepon) VALUES
(1, 'Cabang Jakarta Pusat', 'Jl. Merdeka No. 123', 'Jakarta', 'DKI Jakarta', '10110', '021-1234567', 'jakpus@satulaundry.com', 'Dewi Lestari', 'dewi.jakpus@satulaundry.com', '08111111112'),
(1, 'Cabang Jakarta Selatan', 'Jl. Sudirman No. 456', 'Jakarta', 'DKI Jakarta', '12190', '021-2345678', 'jaksel@satulaundry.com', 'Eko Prasetyo', 'eko.jaksel@satulaundry.com', '08111111113'),
(2, 'Cabang Bandung', 'Jl. Braga No. 789', 'Bandung', 'Jawa Barat', '40111', '022-3456789', 'bandung@bersihcepat.com', 'Fina Setiawati', 'fina.bandung@bersihcepat.com', '08111111114');

-- Data Outlet
INSERT INTO outlet (id_cabang, nama_outlet, alamat, kota, provinsi, kode_pos, telepon, email, jam_buka, jam_tutup, pic_nama, pic_email, pic_telepon) VALUES
(1, 'Outlet Merdeka', 'Jl. Merdeka No. 10', 'Jakarta', 'DKI Jakarta', '10110', '021-1111111', 'merdeka@satulaundry.com', '07:00:00', '21:00:00', 'Agus Prasetyo', 'agus.merdeka@satulaundry.com', '08111111115'),
(1, 'Outlet Thamrin', 'Jl. Thamrin No. 25', 'Jakarta', 'DKI Jakarta', '10110', '021-2222222', 'thamrin@satulaundry.com', '07:00:00', '21:00:00', 'Bambang Susilo', 'bambang.thamrin@satulaundry.com', '08111111116'),
(2, 'Outlet Senayan', 'Jl. Senayan No. 50', 'Jakarta', 'DKI Jakarta', '12190', '021-3333333', 'senayan@satulaundry.com', '08:00:00', '20:00:00', 'Citra Dewi', 'citra.senayan@satulaundry.com', '08111111117'),
(3, 'Outlet Braga', 'Jl. Braga No. 100', 'Bandung', 'Jawa Barat', '40111', '022-4444444', 'braga@bersihcepat.com', '07:30:00', '20:30:00', 'Doni Wijaya', 'doni.braga@bersihcepat.com', '08111111118');

-- Data Pelanggan
INSERT INTO pelanggan (nomor_hp, nama_lengkap, email, alamat, tanggal_lahir, jenis_kelamin, poin_reward) VALUES
('08123456789', 'Budi Santoso', 'budi.santoso@email.com', 'Jl. Anggrek No. 10, Jakarta', '1990-05-15', 'L', 150),
('08234567890', 'Ani Wijaya', 'ani.wijaya@email.com', 'Jl. Melati No. 20, Jakarta', '1992-08-22', 'P', 200),
('08345678901', 'Candra Putra', 'candra.putra@email.com', 'Jl. Mawar No. 30, Bandung', '1988-12-10', 'L', 75),
('08456789012', 'Dewi Lestari', 'dewi.lestari@email.com', 'Jl. Kenanga No. 40, Bandung', '1995-03-30', 'P', 120);

-- Data Pelanggan Outlet (Outlet tempat pelanggan terdaftar)
INSERT INTO pelanggan_outlet (id_pelanggan, id_outlet, tanggal_terdaftar) VALUES
(1, 1, '2023-01-15 10:00:00'),
(1, 2, '2023-03-20 14:30:00'),
(2, 1, '2023-02-10 09:15:00'),
(3, 4, '2023-01-25 11:45:00'),
(4, 1, '2023-03-05 16:20:00'),
(4, 4, '2023-04-12 13:10:00');

-- Contoh penggunaan stored procedure untuk pendaftaran pelanggan baru
-- CALL sp_daftar_pelanggan_baru('08567890123', 'Eko Prasetyo', 'eko@email.com', 'Jl. Dahlia No. 50, Jakarta', '1991-07-20', 'L', 1, @id_pelanggan, @status);
-- SELECT @id_pelanggan, @status;

-- Contoh penggunaan stored procedure untuk mengecek status pelanggan
-- CALL sp_cek_status_pelanggan('08123456789', @id_pelanggan, @status, @nama);
-- SELECT @id_pelanggan, @status, @nama;

-- Data Presensi Pegawai
INSERT INTO presensi (id_pegawai, id_outlet, tanggal, waktu_masuk, waktu_keluar, foto_masuk, foto_keluar, status_presensi, keterangan) VALUES
(1, 1, '2023-06-01', '08:00:00', '17:00:00', '/images/presensi/1_20230601_masuk.jpg', '/images/presensi/1_20230601_keluar.jpg', 'hadir', NULL),
(2, 1, '2023-06-01', '08:15:00', '17:15:00', '/images/presensi/2_20230601_masuk.jpg', '/images/presensi/2_20230601_keluar.jpg', 'hadir', NULL),
(3, 1, '2023-06-01', '08:30:00', '17:30:00', '/images/presensi/3_20230601_masuk.jpg', '/images/presensi/3_20230601_keluar.jpg', 'hadir', NULL),
(1, 1, '2023-06-02', '07:50:00', '16:50:00', '/images/presensi/1_20230602_masuk.jpg', '/images/presensi/1_20230602_keluar.jpg', 'hadir', NULL),
(2, 1, '2023-06-02', '08:05:00', '17:05:00', '/images/presensi/2_20230602_masuk.jpg', '/images/presensi/2_20230602_keluar.jpg', 'hadir', NULL),
(4, 2, '2023-06-01', '08:00:00', '17:00:00', '/images/presensi/4_20230601_masuk.jpg', '/images/presensi/4_20230601_keluar.jpg', 'hadir', NULL),
(5, 4, '2023-06-01', '08:10:00', '17:10:00', '/images/presensi/5_20230601_masuk.jpg', '/images/presensi/5_20230601_keluar.jpg', 'hadir', NULL),
(1, 1, '2023-06-05', '08:20:00', NULL, '/images/presensi/1_20230605_masuk.jpg', NULL, 'izin', 'Acara keluarga'),
(3, 1, '2023-06-05', NULL, NULL, NULL, NULL, 'sakit', 'Sakit flu');

-- Contoh penggunaan stored procedure untuk mencatat presensi
-- CALL sp_catat_presensi('1122334455667788', 1, '2023-06-06', '08:00:00', '/images/presensi/1_20230606_masuk.jpg', 'hadir', NULL, @status);
-- SELECT @status;

-- Contoh penggunaan stored procedure untuk mencatat waktu keluar
-- CALL sp_catat_keluar('1122334455667788', '2023-06-06', '17:00:00', '/images/presensi/1_20230606_keluar.jpg', @status);
-- SELECT @status;

-- Data Pegawai
INSERT INTO pegawai (id_outlet, nik, nama_lengkap, email, telepon, alamat, tanggal_lahir, jenis_kelamin, posisi, gaji, tanggal_masuk) VALUES
(1, '1122334455667788', 'Agus Prasetyo', 'agus.prasetyo@satulaundry.com', '08111111111', 'Jl. Karyawan No. 1, Jakarta', '1985-07-10', 'L', 'Manager', 7500000, '2020-01-15'),
(1, '2233445566778899', 'Rina Setiawati', 'rina.setiawati@satulaundry.com', '08222222222', 'Jl. Karyawan No. 2, Jakarta', '1990-11-25', 'P', 'Kasir', 4500000, '2020-03-01'),
(1, '3344556677889900', 'Bambang Susilo', 'bambang.susilo@satulaundry.com', '08333333333', 'Jl. Karyawan No. 3, Jakarta', '1987-02-18', 'L', 'Operator', 4000000, '2020-02-10'),
(2, '4455667788990011', 'Siti Rahayu', 'siti.rahayu@satulaundry.com', '08444444444', 'Jl. Karyawan No. 4, Jakarta', '1993-09-12', 'P', 'Manager', 7500000, '2021-05-20'),
(4, '5566778899001122', 'Tono Wijaya', 'tono.wijaya@bersihcepat.com', '08555555555', 'Jl. Karyawan No. 5, Bandung', '1988-04-05', 'L', 'Manager', 7500000, '2021-06-01');

-- Data Paket Layanan untuk masing-masing brand
-- Paket untuk brand "Satu Laundry" (id_brand = 1)
INSERT INTO paket_layanan (id_brand, nama_paket, deskripsi, harga_per_kg, durasi_pengerjaan, satuan_durasi, kategori) VALUES
(1, 'Cuci Kering Reguler', 'Cuci dan jemur biasa', 7000, 2, 'hari', 'kiloan'),
(1, 'Cuci Kering Express', 'Cuci dan jemur cepat (6 jam)', 12000, 6, 'jam', 'kiloan'),
(1, 'Cuci Setrika Reguler', 'Cuci, setrika, dan lipat', 12000, 2, 'hari', 'kiloan'),
(1, 'Cuci Setrika Express', 'Cuci, setrika, dan lipat (6 jam)', 18000, 6, 'jam', 'kiloan'),
(1, 'Dry Cleaning', 'Pembersihan khusus untuk bahan tertentu', 25000, 3, 'hari', 'satuan'),
(1, 'Sepatu', 'Pembersihan sepatu', 25000, 1, 'hari', 'satuan');

-- Paket untuk brand "Bersih Cepat" (id_brand = 2)
INSERT INTO paket_layanan (id_brand, nama_paket, deskripsi, harga_per_kg, durasi_pengerjaan, satuan_durasi, kategori) VALUES
(2, 'Laundry Express 3 Jam', 'Laundry selesai dalam 3 jam', 20000, 3, 'jam', 'kiloan'),
(2, 'Laundry Express 6 Jam', 'Laundry selesai dalam 6 jam', 15000, 6, 'jam', 'kiloan'),
(2, 'Laundry Reguler', 'Laundry selesai dalam 1 hari', 8000, 1, 'hari', 'kiloan'),
(2, 'Setrika Saja', 'Hanya setrika', 6000, 6, 'jam', 'kiloan'),
(2, 'Karpet', 'Pembersihan karpet', 30000, 2, 'hari', 'satuan');

-- Paket untuk brand "Wangi Segar" (id_brand = 3)
INSERT INTO paket_layanan (id_brand, nama_paket, deskripsi, harga_per_kg, durasi_pengerjaan, satuan_durasi, kategori) VALUES
(3, 'Cuci Kering Premium', 'Cuci dan jemur dengan pewangi premium', 10000, 2, 'hari', 'kiloan'),
(3, 'Cuci Setrika Premium', 'Cuci, setrika dengan pewangi premium', 15000, 2, 'hari', 'kiloan'),
(3, 'Wangi Tahan Lama', 'Laundry dengan aroma tahan 2 minggu', 18000, 2, 'hari', 'kiloan'),
(3, 'Bedcover', 'Pembersihan bedcover', 35000, 3, 'hari', 'satuan'),
(3, 'Tas/Koper', 'Pembersihan tas atau koper', 40000, 2, 'hari', 'satuan');

-- Data Inventaris
INSERT INTO inventaris (id_outlet, nama_barang, kategori, jumlah_stok, satuan, harga_beli) VALUES
(1, 'Mesin Cuci 10kg', 'Peralatan', 5, 'unit', 5000000),
(1, 'Mesin Pengering', 'Peralatan', 3, 'unit', 7000000),
(1, 'Setrika Uap', 'Peralatan', 10, 'unit', 500000),
(1, 'Deterjen', 'Bahan Baku', 50, 'liter', 25000),
(1, 'Pewangi Laundry', 'Bahan Baku', 30, 'liter', 30000),
(4, 'Mesin Cuci 7kg', 'Peralatan', 3, 'unit', 4000000),
(4, 'Setrika Biasa', 'Peralatan', 5, 'unit', 300000);

-- Data Pengeluaran
INSERT INTO pengeluaran (id_outlet, keterangan, jumlah, tanggal_pengeluaran, kategori) VALUES
(1, 'Pembelian deterjen', 1250000, '2023-06-01', 'perlengkapan'),
(1, 'Gaji bulan Juni', 16000000, '2023-06-05', 'gaji'),
(1, 'Listrik dan air', 2500000, '2023-06-10', 'operasional'),
(4, 'Pembelian setrika', 1500000, '2023-06-02', 'perlengkapan'),
(4, 'Gaji bulan Juni', 5000000, '2023-06-05', 'gaji');

-- Data Transaksi
INSERT INTO transaksi (id_pelanggan, id_outlet, id_pegawai, tanggal_masuk, tanggal_selesai, tanggal_diambil, berat_laundry, total_harga, uang_bayar, uang_kembalian, status_transaksi, status_pembayaran, metode_pembayaran, catatan) VALUES
(1, 1, 2, '2023-06-01 09:00:00', '2023-06-03 15:00:00', '2023-06-04 10:00:00', 5.5, 66000, 70000, 4000, 'diambil', 'lunas', 'tunai', 'Pelanggan ramah'),
(2, 1, 2, '2023-06-02 14:30:00', '2023-06-02 20:30:00', '2023-06-03 14:00:00', 3.2, 57600, 60000, 2400, 'diambil', 'lunas', 'tunai', NULL),
(3, 4, 5, '2023-06-03 10:15:00', '2023-06-03 16:15:00', NULL, 4.0, 60000, 0, 0, 'selesai', 'belum lunas', 'e-wallet', 'Menunggu pembayaran'),
(4, 1, 2, '2023-06-05 16:45:00', '2023-06-07 10:00:00', NULL, 6.8, 81600, 0, 0, 'diproses', 'belum lunas', 'transfer', 'Urgent order');

-- Data Detail Transaksi
INSERT INTO detail_transaksi (id_transaksi, id_paket, kuantitas, harga_satuan, subtotal) VALUES
(1, 3, 5.5, 12000, 66000),
(2, 4, 3.2, 18000, 57600),
(3, 9, 4.0, 15000, 60000),
(4, 3, 6.8, 12000, 81600);

-- Data History Status Transaksi
INSERT INTO history_status_transaksi (id_transaksi, status_lama, status_baru, waktu_perubahan, keterangan) VALUES
(1, 'diterima', 'diproses', '2023-06-01 09:30:00', 'Mulai dicuci'),
(1, 'diproses', 'selesai', '2023-06-03 15:00:00', 'Selesai dicuci dan disetrika'),
(1, 'selesai', 'diambil', '2023-06-04 10:00:00', 'Pelanggan mengambil laundry'),
(2, 'diterima', 'diproses', '2023-06-02 15:00:00', 'Mulai dikerjakan'),
(2, 'diproses', 'selesai', '2023-06-02 20:30:00', 'Selesai dikerjakan'),
(2, 'selesai', 'diambil', '2023-06-03 14:00:00', 'Pelanggan mengambil laundry'),
(3, 'diterima', 'diproses', '2023-06-03 10:30:00', 'Mulai dikerjakan'),
(3, 'diproses', 'selesai', '2023-06-03 16:15:00', 'Selesai dikerjakan'),
(4, 'diterima', 'diproses', '2023-06-05 17:00:00', 'Mulai dikerjakan');