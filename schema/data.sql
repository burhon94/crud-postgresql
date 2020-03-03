INSERT INTO burgers (name, price)
VALUES ('Big Mac', 16000),
       ('Chicken Mac', 12000),
       ('Roll Mac', 12000);

SELECT id, name, price FROM burgers WHERE removed = FALSE;

INSERT INTO burgers(name, price) VALUES (?, ?);

UPDATE burgers SET removed = TRUE WHERE id = ?;

