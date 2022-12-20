-- name: FindRegisteredEmail :one
SELECT email FROM authentication.users
    WHERE users.email= @email
    LIMIT 1;