package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"server/config"
	actorDel "server/internal/Actor/delivery"
	actorRep "server/internal/Actor/repository/postgres"
	actorUsecase "server/internal/Actor/usecase"
	filmDel "server/internal/Film/delivery"
	filmRep "server/internal/Film/repository/postgres"
	filmUsecase "server/internal/Film/usecase"
	sessionDel "server/internal/Session/delivery"
	sessionRep "server/internal/Session/repository/redis"
	sessionUsecase "server/internal/Session/usecase"
	userRep "server/internal/User/repository/postgres"
	"server/internal/middleware"
	"time"
)

// @title Filmoteka API
// @version 1.0
// @license.name Apache 2.0

const PORT = ":8080"

var (
	redisAddr = flag.String("addr", "redis://redis-session:6379/0", "redis addr")

	host     = "test_postgres"
	port     = 5432
	user     = "uliana"
	password = "uliana"
	dbname   = "filmoteka"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

func GetPostgres(psqlInfo string) (*sql.DB, error) {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

func main() {
	flag.Parse()

	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	router := mux.NewRouter()
	authRouter := mux.NewRouter()

	time.Sleep(5 * time.Second)

	db, err := GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err, " ", psqlInfo)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()

	logger := middleware.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())

	actorRepo := actorRep.NewActorRepo(db)
	filmRepo := filmRep.NewFilmRepo(db)
	userRepo := userRep.NewUserRepo(db)
	sessionRepo := sessionRep.NewSessionManager(redisConn)

	actorUC := actorUsecase.NewActorUsecase(actorRepo, filmRepo)
	actorHandler := actorDel.NewActorHandler(actorUC, logger)

	filmUC := filmUsecase.NewFilmUsecase(filmRepo, actorRepo)
	filmHandler := filmDel.NewFilmHandler(filmUC, logger)

	sessionUC := sessionUsecase.NewSessionUsecase(sessionRepo, userRepo)
	sessionHandler := sessionDel.NewSessionHandler(sessionUC, logger)

	actorHandler.RegisterHandler(authRouter)
	filmHandler.RegisterHandler(authRouter)
	sessionHandler.RegisterAuthHandler(authRouter)
	sessionHandler.RegisterHandler(router)
	authMW := middleware.NewSessionMiddleware(sessionUC, logger)

	router.Use(middleware.PanicMiddleware)
	router.Use(logger.ACLogMiddleware)
	authRouter.Use(authMW.AuthMiddleware)

	router.PathPrefix("/api/logout").Handler(authRouter)
	router.PathPrefix("/api/actors").Handler(authRouter)
	router.PathPrefix("/api/films").Handler(authRouter)

	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Println("Server start at port", PORT[1:])
	err = server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

}
