package app

import (
	"crud/pkg/tools/services/burgers"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

// описание сервиса, который хранит зависимости и выполняет работу
type server struct {
	pool          *pgxpool.Pool
	router        http.Handler
	burgersSvc    *burgers.BurgersSvc
	templatesPath string
	assetsPath    string
}

func NewServer(router http.Handler, pool *pgxpool.Pool, burgersSvc *burgers.BurgersSvc, templatesPath string, assetsPath string) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if pool == nil {
		panic(errors.New("pool can't be nil"))
	}
	if burgersSvc == nil {
		panic(errors.New("burgersSvc can't be nil"))
	}
	if templatesPath == "" {
		panic(errors.New("templatesPath can't be empty"))
	}
	if assetsPath == "" {
		panic(errors.New("assetsPath can't be empty"))
	}

	return &server{
		router: router,
		pool: pool,
		burgersSvc: burgersSvc,
		templatesPath: templatesPath,
		assetsPath: assetsPath,
	}
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}
