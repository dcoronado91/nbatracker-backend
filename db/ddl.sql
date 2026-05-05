DROP TABLE IF EXISTS players;

CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    team TEXT NOT NULL,
    image_url TEXT DEFAULT '',
    championships INT DEFAULT 0 CHECK (championships >= 0),
    mvp INT DEFAULT 0 CHECK (mvp >= 0),
    finals_mvp INT DEFAULT 0 CHECK (finals_mvp >= 0),
    dpoy INT DEFAULT 0 CHECK (dpoy >= 0),
    roty INT DEFAULT 0 CHECK (roty >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);