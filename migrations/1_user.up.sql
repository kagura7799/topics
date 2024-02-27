create table if not exists users (
	id          BIGSERIAL PRIMARY KEY,
    username    TEXT NOT NULL UNIQUE,
	email       TEXT NOT NULL UNIQUE,
	pass_hash        TEXT NOT NULL,
	name        TEXT,
    about       TEXT,
	avatar_url  TEXT,
	created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	deleted_at  TIMESTAMP WITH TIME ZONE 
)

CREATE INDEX IF NOT EXISTS idx_user ON users (email, username);

CREATE TABLE IF NOT EXISTS topics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    avatar_url  TEXT,	
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    link TEXT,
    content TEXT,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    topic_id INTEGER REFERENCES topics(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE,
    parent_comment_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS votes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    vote_type VARCHAR(20) CHECK (vote_type IN ('up', 'down')),
    votable_id INTEGER NOT NULL,
    votable_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    vote_type VARCHAR(20) CHECK (vote_type IN ('up', 'down')),
    votable_id INTEGER NOT NULL,
    votable_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Index for polymorphic association
CREATE INDEX idx_votes_on_votable ON votes(votable_type, votable_id);

