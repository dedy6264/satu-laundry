-- Script to add user access and login functionality

-- Create user_access table to store login credentials and access permissions
CREATE TABLE IF NOT EXISTS user_access (
    id_access SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'karyawan',
    is_active BOOLEAN DEFAULT true,
    last_login TIMESTAMP,
    reference_level VARCHAR(10),
    reference_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes for faster queries
CREATE INDEX IF NOT EXISTS idx_user_access_username ON user_access(username);
CREATE INDEX IF NOT EXISTS idx_user_access_active ON user_access(is_active);

-- CRUD Operations for User Access

-- CREATE - Add new user access
-- INSERT INTO user_access (username, password, role, reference_level, reference_id) 
-- VALUES ($1, $2, $3, $4, $5);

-- READ - Get user access by ID
-- SELECT id_access, username, role, is_active, last_login, 
--        reference_level, reference_id, created_at, updated_at
-- FROM user_access
-- WHERE id_access = $1;

-- READ - Get user access by username
-- SELECT id_access, username, password, role, is_active, 
--        last_login, reference_level, reference_id, created_at, updated_at
-- FROM user_access
-- WHERE username = $1 AND is_active = true;

-- READ - Get all user access with pagination
-- SELECT id_access, username, role, is_active, last_login, 
--        reference_level, reference_id, created_at, updated_at
-- FROM user_access
-- ORDER BY created_at DESC
-- LIMIT $1 OFFSET $2;

-- UPDATE - Update user access
-- UPDATE user_access 
-- SET username = $1, role = $2, is_active = $3, updated_at = NOW()
-- WHERE id_access = $4;

-- UPDATE - Update user password
-- UPDATE user_access 
-- SET password = $1, updated_at = NOW()
-- WHERE id_access = $2;

-- UPDATE - Update last login time
-- UPDATE user_access 
-- SET last_login = NOW()
-- WHERE id_access = $1;

-- DELETE - Remove user access
-- DELETE FROM user_access WHERE id_access = $1;