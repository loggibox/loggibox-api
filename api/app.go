package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/fiscaluno/pandorabox"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/loggibox/loggibox-api/api/handler"
	"github.com/loggibox/loggibox-api/pkg/middleware"
	"github.com/loggibox/loggibox-api/pkg/user"
)

// Start ...
func Start() {

	r := mux.NewRouter()

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	// user
	userRepo := user.NewInmemRepository()
	userService := user.NewService(userRepo)
	handler.MakeUserHandlers(r, *n, userService)

	http.Handle("/", r)
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	port := pandorabox.GetOSEnvironment("PORT", "5000")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + port,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	logger.Println("Listen on port:" + port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
