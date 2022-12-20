CREATE TABLE IF NOT EXISTS authentication.users (

    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);