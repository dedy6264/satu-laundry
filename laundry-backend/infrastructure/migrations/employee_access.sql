-- Script to add employee access and login functionality

-- Create employee_access table to store login credentials and access permissions
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

-- Add indexes for faster queries
CREATE INDEX IF NOT EXISTS idx_employee_access_username ON employee_access(username);
CREATE INDEX IF NOT EXISTS idx_employee_access_pegawai ON employee_access(id_pegawai);
CREATE INDEX IF NOT EXISTS idx_employee_access_active ON employee_access(is_active);

-- Insert default access for existing employees (optional)
-- This would need to be run after adding all employees to the system
-- INSERT INTO employee_access (id_pegawai, username, password, role)
-- SELECT id_pegawai, email, '', 'karyawan'
-- FROM pegawai
-- WHERE email IS NOT NULL AND email != ''
-- ON CONFLICT (username) DO NOTHING;

-- CRUD Operations for Employee Access

-- CREATE - Add new employee access
-- INSERT INTO employee_access (id_pegawai, username, password, role) 
-- VALUES ($1, $2, $3, $4);

-- READ - Get employee access by ID
-- SELECT ea.id_access, ea.id_pegawai, ea.username, ea.role, ea.is_active, ea.last_login, 
--        ea.created_at, ea.updated_at,
--        p.nama_lengkap, p.email, p.id_outlet
-- FROM employee_access ea
-- JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
-- WHERE ea.id_access = $1;

-- READ - Get employee access by username
-- SELECT ea.id_access, ea.id_pegawai, ea.username, ea.password, ea.role, ea.is_active, 
--        ea.last_login, ea.created_at, ea.updated_at,
--        p.nama_lengkap, p.email, p.id_outlet
-- FROM employee_access ea
-- JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
-- WHERE ea.username = $1 AND ea.is_active = true;

-- READ - Get all employee access with pagination
-- SELECT ea.id_access, ea.id_pegawai, ea.username, ea.role, ea.is_active, ea.last_login, 
--        ea.created_at, ea.updated_at,
--        p.nama_lengkap, p.email, p.id_outlet
-- FROM employee_access ea
-- JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
-- ORDER BY ea.created_at DESC
-- LIMIT $1 OFFSET $2;

-- UPDATE - Update employee access
-- UPDATE employee_access 
-- SET username = $1, role = $2, is_active = $3, updated_at = NOW()
-- WHERE id_access = $4;

-- UPDATE - Update employee password
-- UPDATE employee_access 
-- SET password = $1, updated_at = NOW()
-- WHERE id_access = $2;

-- UPDATE - Update last login time
-- UPDATE employee_access 
-- SET last_login = NOW()
-- WHERE id_access = $1;

-- DELETE - Remove employee access
-- DELETE FROM employee_access WHERE id_access = $1;

-- Additional useful queries

-- Get employees by outlet with access information
-- SELECT ea.id_access, ea.username, ea.role, ea.is_active, ea.last_login,
--        p.id_pegawai, p.nama_lengkap, p.email, p.id_outlet
-- FROM employee_access ea
-- JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
-- WHERE p.id_outlet = $1
-- ORDER BY p.nama_lengkap;

-- Count active employees by outlet
-- SELECT p.id_outlet, COUNT(*) as active_employees
-- FROM employee_access ea
-- JOIN pegawai p ON ea.id_pegawai = p.id_pegawai
-- WHERE ea.is_active = true
-- GROUP BY p.id_outlet;