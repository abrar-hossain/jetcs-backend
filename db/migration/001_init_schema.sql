CREATE TABLE submissions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author_name TEXT NOT NULL,
    email TEXT NOT NULL,
    status TEXT DEFAULT 'under review',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
