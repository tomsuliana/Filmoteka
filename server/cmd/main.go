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
	"server/server/config"
	actorDel "server/server/internal/Actor/delivery"
	actorRep "server/server/internal/Actor/repository/postgres"
	actorUsecase "server/server/internal/Actor/usecase"
	filmDel "server/server/internal/Film/delivery"
	filmRep "server/server/internal/Film/repository/postgres"
	filmUsecase "server/server/internal/Film/usecase"
	sessionDel "server/server/internal/Session/delivery"
	sessionRep "server/server/internal/Session/repository/redis"
	sessionUsecase "server/server/internal/Session/usecase"
	userRep "server/server/internal/User/repository/postgres"
	"server/server/internal/middleware"
)

const PORT = ":8080"

var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")

	host     = "localhost"
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
