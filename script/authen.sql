-- Create the roles table with UUID primary key and the specified fields
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- UUID primary key
    name VARCHAR(255) NOT NULL,                    -- Role name (e.g., End User, Admin)
    icon VARCHAR(255),                             -- Icon representation of the role
    description TEXT,                              -- Description of the role
    ip_access TEXT,                                -- Custom field for IP-based access control
    enforce_tfa BOOLEAN DEFAULT FALSE,             -- Enforce two-factor authentication
    admin_access BOOLEAN DEFAULT FALSE,            -- Admin access flag
    app_access BOOLEAN DEFAULT FALSE,              -- App access flag
    public_id TEXT                                 -- Public identifier for roles
);

-- Create the users table with UUID primary key and the specified fields
CREATE TABLE public.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- UUID primary key
    first_name VARCHAR(255),                       -- First name of the user
    last_name VARCHAR(255),                        -- Last name of the user
    email VARCHAR(255) UNIQUE NOT NULL,            -- User's email, must be unique
    password VARCHAR(255),                         -- Encrypted password (nullable for social login users)
    location VARCHAR(255),                         -- User's location (optional)
    title VARCHAR(255),                            -- User's title (optional)
    tags JSON,                                     -- Tags for the user
    avatar UUID,                                   -- Avatar image ID
    language VARCHAR(50),                          -- Language preference of the user
    tfa_secret VARCHAR(255),                       -- Two-factor authentication secret
    status VARCHAR(50) DEFAULT 'active',           -- Status (e.g., active, inactive)
    role UUID REFERENCES roles(id),                -- Foreign key to roles table (role UUID)
    token VARCHAR(255),                            -- Token for verification or authentication
    last_access TIMESTAMPTZ,                       -- Last access timestamp
    provider VARCHAR(255),                         -- Social login provider (e.g., google, facebook)
    external_identifier VARCHAR(255),              -- External identifier from the provider (e.g., google ID)
    auth_data JSON,                                -- Authentication data in JSON format
    email_notifications BOOLEAN DEFAULT TRUE,      -- Email notifications flag
    birthday TIMESTAMP,                            -- User's birthday (optional)
    referrer_email VARCHAR(255),                   -- Referring user's email (optional)
    fullname VARCHAR(255),                         -- Full name of the user
    search TEXT,                                   -- Searchable text field
    facebook_url VARCHAR(255),                     -- Facebook URL (optional)
    verify_token VARCHAR(255),                     -- Email verification token
    is_active BOOLEAN DEFAULT TRUE,                -- Flag for whether the user is active
    phone_number VARCHAR(50),                      -- User's phone number
    meta JSONB,                                    -- Additional metadata in JSONB format
    job_title VARCHAR(255),                        -- User's job title (optional)
    appearance VARCHAR(255),                       -- Appearance settings (optional)
    theme_dark VARCHAR(255),                       -- Dark theme settings (optional)
    theme_light VARCHAR(255),                      -- Light theme settings (optional)
    theme_light_overrides JSON,                    -- Light theme overrides in JSON format
    theme_dark_overrides JSON,                     -- Dark theme overrides in JSON format
    theme VARCHAR(255),                            -- Current theme preference
    zalo TEXT,                                     -- Zalo contact info (optional)
    tiktok TEXT,                                   -- TikTok profile URL (optional)
    attachments TEXT,                              -- Attachments (optional)
    long_description TEXT,                         -- Long description (optional)
    note TEXT,                                     -- Note field (optional)
    source JSONB,                                  -- Source of the user data in JSONB
    instagram TEXT,                                -- Instagram profile (optional)
    date_created TIMESTAMPTZ DEFAULT NOW()         -- Account creation timestamp
);



CREATE TABLE public.student_target (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),   -- UUID primary key
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to the users table
    target_study_duration INT,                       -- Study duration (in minutes or hours)
    target_reading FLOAT,                            -- Reading score target
    target_listening FLOAT,                          -- Listening score target
    target_speaking FLOAT,                           -- Speaking score target
    target_writing FLOAT,                            -- Writing score target
    next_exam_date TIMESTAMPTZ                       -- Timestamp for the next exam date
);

CREATE TABLE public.access_control_list (
    id SERIAL PRIMARY KEY,            -- Primary key with auto-increment
    action_id TEXT NOT NULL,           -- Identifier for the action (e.g., CREATE_USER)
    role_id UUID REFERENCES roles(id), -- Foreign key to the roles table
    status SMALLINT DEFAULT 1,         -- Status field (1 for active, 0 for inactive)
    user_id UUID REFERENCES users(id)  -- Foreign key to the users table (optional)
);