package main

// package
// import
// var + type
// method + function

import (
	"context"
	"flag"
	"github.com/AbduvokhidovRustamzhon/codeGreen/cmd/crud/app"
	"github.com/AbduvokhidovRustamzhon/codeGreen/pkg/crud/services/burgers"
	"github.com/AbduvokhidovRustamzhon/codeGreen/pkg/crud/services/files"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var (
	hostF   = flag.String("host", "", "Server host")
	portF   = flag.String("port", "", "Server port")
	dsnF    = flag.String("dsn", "", "Postgres DSN")
)
var (
	EHOST   = "HOST"
	EPORT   = "PORT"
	EDSN    = "DATABASE_URL"
)

func main() {
	flag.Parse()

	host, ok := FlagOrEnv(*hostF, EHOST)
	if !ok {
		log.Panic("can't get host")
	}
	port, ok := FlagOrEnv(*portF, EPORT)
	if !ok {
		log.Panic("can't get port")
	}
	dsn, ok := FlagOrEnv(*dsnF, EDSN)
	if !ok {
		log.Panic("can't get dsn")
	}
	addr := net.JoinHostPort(host, port)
	start(addr, dsn)
}

func start(addr string, dsn string) {
	router := app.NewExactMux()
	// Context: <-
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	templatesPath := filepath.Join("web", "templates")
	assetsPath := filepath.Join("web", "assets")
	mediaPath := filepath.Join("web", "media")

	burgersSvc := burgers.NewBurgersSvc(pool)
	filesSvc := files.NewFilesSvc(mediaPath)
	server := app.NewServer(
		router,
		pool,
		burgersSvc, // DI + Containers
		filesSvc,
		templatesPath,
		assetsPath,
		mediaPath,
	)
	server.InitRoutes()


	// server'ы должны работать "вечно"
	panic(http.ListenAndServe(addr, server)) // поднимает сервер на определённом адресе и обрабатывает запросы
}
func FlagOrEnv(flag string, envKey string) (string, bool) {
	if flag != "" {
		return flag, true
	}
	return os.LookupEnv(envKey)
}