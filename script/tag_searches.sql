CREATE TABLE tag_search (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) DEFAULT 'published',
    sort INT,
    user_created UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    user_updated UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255),
    priority INT,
    is_shown BOOLEAN,
    color VARCHAR(255),
    description TEXT,
    code VARCHAR(255),
    icon UUID --files PK later
);

CREATE TABLE tag_position (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) DEFAULT 'draft',
    sort INT,
    user_created UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    user_updated UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255),
    position VARCHAR(255)
);

CREATE TABLE tag_position_tag_search (
    id SERIAL PRIMARY KEY,
    tag_position_id INT REFERENCES tag_position(id) ON DELETE CASCADE,
    tag_search_id INT REFERENCES tag_search(id) ON DELETE CASCADE,
    sort INT
);

