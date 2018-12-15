package api

import (
	"log"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"github.com/codegangsta/negroni"
	"github.com/fiscaluno/pandorabox"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/loggibox/loggibox-api/api/handler"
	"github.com/loggibox/loggibox-api/pkg/middleware"
	"github.com/loggibox/loggibox-api/pkg/packet"
	"github.com/loggibox/loggibox-api/pkg/user"
	ctext "golang.org/x/net/context"
	"google.golang.org/api/option"
)

// Start ...
func Start() {

	r := mux.NewRouter()

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	dURL := pandorabox.GetOSEnvironment("FDATABASE_URL", "https://loggibox.firebaseio.com")
	ctx := ctext.Background()
	conf := &firebase.Config{
		DatabaseURL: dURL,
	}

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// Packet
	packetRepo := packet.NewFirebaseRepo(client)
	packetService := packet.NewService(packetRepo)
	handler.MakePacketHandlers(r, *n, packetService)

	// User
	// userRepo := user.NewInmemRepository()
	userRepo := user.NewFirebaseRepo(client)
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
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
