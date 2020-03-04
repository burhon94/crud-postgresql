package burgers

import (
	"context"
	"crud/pkg/crud/errors"
	"crud/pkg/crud/models"
	"crud/pkg/crud/services"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type BurgersSvc struct {
	pool *pgxpool.Pool
}

func NewBurgersSvc(pool *pgxpool.Pool) *BurgersSvc {
	if pool == nil {
		log.Printf("can't pool be nil: %v", pool)
	}
	return &BurgersSvc{pool: pool}
}

func (service *BurgersSvc) BurgersList() (list []models.Burger, err error) {
	list = make([]models.Burger, 0)
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		err = errors.MyError("can't execute pool: ", err)
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), services.GetBurgers)
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

	_, err = conn.Exec(context.Background(), services.SaveBurger, model.Name, model.Price)
	if err != nil {
		err = errors.MyError("can't save burger: ", err)
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

	_, err = conn.Exec(context.Background(), services.RemoveBurger, id)
	if err != nil {
		err = errors.MyError("can't remove burger: ", err)
		return err
	}

	return nil
}