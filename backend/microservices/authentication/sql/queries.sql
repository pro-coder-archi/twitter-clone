-- name: FindRegisteredEmail :one
SELECT * FROM authentication.users
    WHERE users.email= ?
        LIMIT 1;