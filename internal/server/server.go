package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/ikilonchic/WEB_LAB-3-4/internal/database"
)

type ctxKey int64

// Server ...
type Server struct {
	config 		*Config
	logger 		*logrus.Logger
	router 		*mux.Router
	sql			*database.PostgresClient
	redis		*database.RedisClient
}

// Start ...
func (serve *Server) Start() error {
	if err := serve.configureLogger(); err != nil {
		return err
	}

	if err := serve.configureDatabase(); err != nil {
		return err
	}

	serve.configureRouter()

	serve.logger.Infof("Server are starting on port%s ...", serve.config.Port)

	return http.ListenAndServe(serve.config.Port, serve.router)
}

// Configure server: //
// - logger;		 //
// - databases;		 //
// - router.		 //

// Configure logger ...
func (serve *Server) configureLogger() error {
	level, err := logrus.ParseLevel(serve.config.LogLevel)
	if err != nil {
		return err
	}

	serve.logger.SetLevel(level)
	return nil
}

// Configure database ...
func (serve *Server) configureDatabase() error {
	//////////////////////
	return nil
}

// Configure router ...
func (serve *Server) configureRouter() {
	serve.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(serve.config.Static))))

	serve.router.Use(serve.setRequestID)
	serve.router.Use(serve.logRequest)

	serve.router.HandleFunc("/signin", serve.getPage("SignIn")).Methods("GET")
	serve.router.HandleFunc("/signup", serve.getPage("SignUp")).Methods("GET")

	serve.router.HandleFunc("/signin", serve.signIn()).Methods("POST")
	serve.router.HandleFunc("/signup", serve.signUp()).Methods("POST")
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}
