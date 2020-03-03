package burgers

import (
	"context"
	"crud/pkg/crud/errors"
	"crud/pkg/crud/models"
	"crud/schema"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type BurgersSvc struct {
	pool *pgxpool.Pool // dependency
}

func NewBurgersSvc(pool *pgxpool.Pool) *BurgersSvc {
	if pool == nil {
		log.Printf("can't pool be nil: %v", pool) // <- be accurate
	}
	return &BurgersSvc{pool: pool}
}

func (service *BurgersSvc) BurgersList() (list []models.Burger, err error) {
	list = make([]models.Burger, 0) // TODO: for REST API
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		err = errors.MyError("can't  tx: ", err)
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), schema.GetBurgers)
	if err != nil {
		err = errors.MyError("can't query: execute pool", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Burger{}
		err := rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			err = errors.MyError("can't scan row: ", err)
			return nil, err
		}
		list = append(list, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *BurgersSvc) Save(model models.Burger) (err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		err = errors.MyError("can't execute pool: ", err)
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		err = errors.MyError("can't begin tx: ", err)
		return err
	}
	defer func() {
		if err != nil {
			err = errors.MyError("can't defer tx: ", err)
			_ = tx.Rollback(context.Background())
			return
		}
		err = tx.Commit(context.Background())
	}()

	_, err = tx.Exec(context.Background(), schema.AddBurger, model.Name, model.Price)
	if err != nil {
		err = errors.MyError("can't execute tx: ", err)
		return err
	}

	return nil
}

func (service *BurgersSvc) RemoveById(id int64) (err error) {

	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		err = errors.MyError("can't execute pool: ", err)
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		err = errors.MyError("can't begin tx: ", err)
		return err
	}
	defer func() {
		if err != nil {
			err = errors.MyError("can't defer tx: ", err)
			_ = tx.Rollback(context.Background())
			return
		}
		err = tx.Commit(context.Background())
		if err != nil {
			err = errors.MyError("can't commit tx: ", err)
			return
		}
	}()

	_, err = tx.Exec(context.Background(), schema.RemoveBurger, id)
	if err != nil {
		err = errors.MyError("can't execute tx: ", err)
		return err
	}

	return nil
}
