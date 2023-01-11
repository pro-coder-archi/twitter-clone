-- name: FindRegisteredEmail :one
SELECT * FROM users
    WHERE users.email= @email
        LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users
    (email, password)
        VALUES (@email, @password);

-- name: GetPasswordForEmail :one
SELECT password FROM users
    WHERE users.email= @email
        LIMIT 1;