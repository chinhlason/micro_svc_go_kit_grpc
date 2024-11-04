-- +goose Up
CREATE TABLE IF NOT EXISTS db2 (
    id SERIAL,
    username varchar(20),
    password varchar(20),
    PRIMARY KEY(id)
);

INSERT INTO db2 (username, password) VALUES ('admin2', 'admin2');
INSERT INTO db2 (username, password) VALUES ('user2', 'user2');

-- +goose Down
DROP TABLE db2;