-- +goose Up
CREATE TABLE users
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255),
    password VARCHAR(255),
    PRIMARY KEY (id)

);

CREATE TABLE rooms
(
    id SERIAL NOT NULL,
    name VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE user_room
(
    room_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY (user_id, room_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (room_id) REFERENCES rooms (id)
);

CREATE TABLE messages
(
    id SERIAL NOT NULL,
    room_id INT NOT NULL,
    user_id INT NOT NULL,
    text    VARCHAR(255),
    PRIMARY KEY (user_id, room_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (room_id) REFERENCES rooms (id)
);

-- +goose Down
DROP TABLE users;
DROP TABLE rooms;
DROP TABLE user_room;
DROP TABLE messages;
