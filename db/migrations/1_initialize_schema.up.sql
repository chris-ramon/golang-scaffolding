CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    uuid       TEXT                                               NOT NULL,
    username   TEXT                                               NOT NULL,
    first_name TEXT                                               NULL,
    last_name  TEXT                                               NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
