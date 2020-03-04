CREATE TABLE burgers (
   id BIGSERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   price INT NOT NULL CHECK ( price > 0 ),
   removed BOOLEAN DEFAULT FALSE
);
