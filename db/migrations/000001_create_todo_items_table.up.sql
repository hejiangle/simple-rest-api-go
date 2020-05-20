CREATE TABLE todo_items (
    id serial PRIMARY KEY ,
    content VARCHAR(225) NOT NULL ,
    status BOOLEAN NOT NULL ,
    created_at TIME NOT NULL ,
    updated_at TIME ,
    deleted_at TIME
);