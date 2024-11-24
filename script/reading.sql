CREATE TABLE quiz (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    sort INT,
    user_created UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    user_updated UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    type INT,
    content TEXT,
    title VARCHAR(255),
    "order" INT,
    time INT,
    listening UUID, --primary key later
    description TEXT,
    instruction INT,
    quiz_code VARCHAR(255),
    limit_submit INT,
    question TEXT,
    samples TEXT,
    thumbnail UUID, --primary key later
    vote_count INT DEFAULT 0,
    quiz_type SMALLINT,
    full_id INT,
    is_test BOOLEAN,
    mode SMALLINT DEFAULT 0,
    simplified_id INT,
    mock_test_id INT, --primary key later
    mock_test_type INT, --primary key later
    total_submitted INT DEFAULT 0,
    short_description TEXT,
    practice_listing_priority INT DEFAULT 0,
    is_public BOOLEAN DEFAULT TRUE,
    writing_task_type INT,
    meta JSONB
);


CREATE TABLE part (
    id SERIAL PRIMARY KEY,
    sort INT,
    user_created UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255),
    "order" INT,
    content TEXT,
    quiz INT REFERENCES public.quiz(id) ON DELETE CASCADE,
    time INT,
    passage INT DEFAULT 0,
    simplified_content TEXT,
    question_count SMALLINT DEFAULT 0,
    listen_from INT,
    listen_to INT
);

CREATE TABLE question (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL DEFAULT 'draft',
    sort INT,
    user_created UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    user_updated UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255),
    content TEXT,
    type VARCHAR(255),
    single_choice_radio JSON,
    selection JSON,
    multiple_choice JSON,
    gap_fill_in_blank TEXT,
    quiz INT REFERENCES public.quiz(id) ON DELETE CASCADE,
    selection_option JSON,
    locate VARCHAR(255),
    "order" INT,
    explain TEXT,
    description TEXT,
    content_writing INT,
    part INT REFERENCES public.part(id) ON DELETE CASCADE,
    time_to_think INT,
    listen_from INT,
    question_type VARCHAR(255),
    instruction TEXT
);

CREATE TABLE quiz_part (
    id SERIAL PRIMARY KEY,
    quiz_id INT REFERENCES public.quiz(id) ON DELETE CASCADE,
    part_id INT REFERENCES public.part(id) ON DELETE CASCADE,
    sort INT DEFAULT 0
);


CREATE TABLE type (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL DEFAULT 'draft',
    sort INT,
    user_created UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    user_updated UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255),
    public_id TEXT
);

CREATE TABLE type (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL DEFAULT 'draft',
    sort INT,
    user_created UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    user_updated UUID REFERENCES public.users(id) ON DELETE SET NULL,
    date_updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255),
    public_id TEXT
);


ALTER TABLE quiz
ADD CONSTRAINT fk_quiz_type
FOREIGN KEY (type)
REFERENCES type(id)
ON DELETE SET NULL;

CREATE TABLE quiz_tag_search (
    id SERIAL PRIMARY KEY,
    quiz_id INT REFERENCES public.quiz(id) ON DELETE CASCADE,
    tag_search_id INT REFERENCES public.tag_search(id) ON DELETE CASCADE
);

