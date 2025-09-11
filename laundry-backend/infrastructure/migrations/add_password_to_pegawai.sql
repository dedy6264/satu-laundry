-- Script to add password column to pegawai table

-- Add password column to pegawai table
ALTER TABLE pegawai 
ADD COLUMN IF NOT EXISTS password VARCHAR(255);

-- Add index for faster login queries
CREATE INDEX IF NOT EXISTS idx_pegawai_email ON pegawai(email);
CREATE INDEX IF NOT EXISTS idx_pegawai_nik ON pegawai(nik);