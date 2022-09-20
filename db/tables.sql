\c parkingservice

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    role varchar(20) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT '0',
    last_login TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS email_tokens
(
    token VARCHAR(64) NOT NULL,
    valid_from TIMESTAMPTZ NOT NULL,
    valid_to TIMESTAMPTZ NOT NULL,
    user_id INT NOT NULL,

    CONSTRAINT fk_user
        FOREIGN KEY(user_id) REFERENCES users(id)
);
