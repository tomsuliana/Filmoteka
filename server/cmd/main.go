package main

import (
	"database/sql"
	"errors"
	"fmt"
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
	"server/server/internal/middleware"
)

const PORT = ":8080"

var (
	host     = "localhost"
	port     = 5432
	user     = "uliana"
	password = "uliana"
	dbname   = "filmoteka"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

//GetPostgres gets postgres connection
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
	router := mux.NewRouter()
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
	actorUC := actorUsecase.NewActorUsecase(actorRepo)
	actorHandler := actorDel.NewActorHandler(actorUC, logger)

	filmRepo := filmRep.NewFilmRepo(db)
	filmUC := filmUsecase.NewFilmUsecase(filmRepo, actorRepo)
	filmHandler := filmDel.NewFilmHandler(filmUC, logger)

	actorHandler.RegisterHandler(router)
	filmHandler.RegisterHandler(router)

	router.Use(middleware.PanicMiddleware)
	router.Use(logger.ACLogMiddleware)

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
