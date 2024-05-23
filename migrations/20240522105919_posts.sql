-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    author VARCHAR(20),
    content TEXT,
    title VARCHAR(100),
    commented BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    author VARCHAR(20),
    path INTEGER[] NOT NULL,
    content VARCHAR(2000)
    -- CONSTRAINT fk_post
    --   FOREIGN KEY(id_post) 
    --     REFERENCES posts(id)
    --     ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_comment(
    post_id INT REFERENCES posts ON DELETE CASCADE,
    comment_id INT REFERENCES comments ON DELETE SET NULL,
    UNIQUE(post_id, comment_id)
);
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS post_comment;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd

-- -- +goose StatementBegin
-- SELECT
--     path,
--     content,
-- FROM comments
-- WHERE id IN (
--     SELECT
--     comment_id
--     FROM post_comments
--     WHERE post_id = 1 
--     )
-- ORDER BY id
-- LIMIT 20
-- OFFSET 0;
-- -- +goose StatementEnd