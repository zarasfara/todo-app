CREATE TABLE IF NOT EXISTS posts 
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);