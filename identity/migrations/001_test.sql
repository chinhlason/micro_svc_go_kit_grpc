-- +goose Up
CREATE TABLE IF NOT EXISTS db1 (
    id SERIAL,
    username varchar(20),
    password varchar(20),
    PRIMARY KEY(id)
);

INSERT INTO db1 (username, password) VALUES ('admin', 'admin');
INSERT INTO db1 (username, password) VALUES ('user', 'user');

-- +goose Down
DROP TABLE db1;