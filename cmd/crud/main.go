package main

import (
	"context"
	"crud/cmd/crud/app"
	"crud/pkg/crud/services/burgers"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const wildCard = "0.0.0.0"
const dbUrl = "postgres://tagspmotvklkfi:9cb1a3d6f70ad82baecafe26750b184e30e1dfeed0ec884b1f1aee6b119f4f3d@ec2-18-210-51-239.compute-1.amazonaws.com:5432/dcs5aet6f8io8d"
const dbLocal = "postgres://tagspmotvklkfi:9cb1a3d6f70ad82baecafe26750b184e30e1dfeed0ec884b1f1aee6b119f4f3d@localhost:5432/dcs5aet6f8io8d"
const db = "postgres://app:pass@localhost:5432/app"
const numberPort = "9999"

var (
	host = flag.String("host", "", "Server host")
	dsn  = flag.String("dsn", "", "Postgres DSN")
)

func main() {
	*host = wildCard
	log.Println("get port to connect")
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = numberPort
	}
	flag.Parse()
	log.Println("set address to connect")
	addr := net.JoinHostPort(*host, port)
	log.Printf("address to connect: %s", addr)
	log.Println("set database to connect")
	if *dsn == "" {
		*dsn, ok = os.LookupEnv("DATABASE_URL")
		if !ok {
			*dsn = dbUrl
		}
	}
	log.Printf("try start server on: %s, dbUrl %v", addr, dsn)
	start(addr, *dsn)
	log.Printf("server succes on: %s, dbUrl %v", addr, dsn)
}

func start(addr string, dsn string) {
	router := app.NewExactMux()
	log.Println("try creat pool to connect")
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Printf("can't create pool: %v", err)
		panic(err)
	}
	burgersSvc := burgers.NewBurgersSvc(pool)
	log.Println("server upping")
	server := app.NewServer(
		router,
		pool,
		burgersSvc, // DI + Containers
		filepath.Join("web", "templates"),
		filepath.Join("web", "assets"),
	)
	server.InitRoutes()

	panic(http.ListenAndServe(addr, server))
}
