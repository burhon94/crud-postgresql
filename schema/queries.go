package schema

const GetBurgers = `SELECT id, name, price FROM burgers WHERE removed = FALSE;`
const AddBurger = `INSERT INTO burgers(name, price) VALUES ($1, $2);`
const RemoveBurger = `UPDATE burgers SET removed = TRUE WHERE id = $1;`
