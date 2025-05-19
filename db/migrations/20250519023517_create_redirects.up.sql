-- 002_create_redirects.sql
CREATE TABLE IF NOT EXISTS redirects (
  id SERIAL PRIMARY KEY,
  short_code TEXT UNIQUE NOT NULL,
  url_id INTEGER REFERENCES urls(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
