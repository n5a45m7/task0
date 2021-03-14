package main

import (
	api "app/api/impl"
	storageAbst "app/storage"
	storage "app/storage/memory"
	transport "app/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Info("start")
	// create all storages
	userStorage := storage.NewUserStorage()
	accStorage := storage.NewAccountStorage()
	txStorage := storage.NewTransactionStorage()

	// create default user
	userStorage.Create(
		storageAbst.CreateUserDTO{
			Name:    "Alexa",
			Surname: "Amazon",
		},
	)

	// create all api
	userAPI := api.NewUserAPI(userStorage, accStorage, txStorage)
	accAPI := api.NewAccountAPI(accStorage, txStorage)

	// create all handlers
	userGetInfoHandler := transport.NewUserGetInfoHandler(userAPI)
	accCreateHandler := transport.NewAccountCreateHandler(accAPI)

	router := mux.NewRouter()
	router.HandleFunc("/account", accCreateHandler.Handler).Methods(http.MethodPost)
	router.HandleFunc("/userinfo", userGetInfoHandler.Handler).Methods(http.MethodGet)

	handler := cors.Default().Handler(router)
	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	log.Info("done")
}
