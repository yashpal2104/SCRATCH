-- +goose Up
-- We Add a new column api_key in the table to authenticate the user(only api_key)
-- which is unique not null and only 64 characters long(the varchar(64)) and set a default api key for every existing record in our table
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea), 'hex')
);
-- random() is a built-in PostgreSQL function that returns a random floating-point number between 0.0 and 1.0 (exclusive).
-- The ::text type cast converts the floating-point number to text (string representation).
-- The ::bytea cast converts the text representation of the random number into a binary format (bytea).
-- sha256(data) computes the SHA-256 hash of the input binary data.
-- encode(data, 'hex') converts binary data into a human-readable hexadecimal string.

-- +goose Down
-- Down migrations just undoes what was done in the Up migrations
ALTER TABLE users DROP COLUMN api_key;