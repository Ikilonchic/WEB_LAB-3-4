package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"github.com/ikilonchic/WEB_LAB-3-4/internal/models"
)

// Server ...
type Server struct {
	config 		*Config
	logger 		*logrus.Logger
	router 		*mux.Router
	sql			*gorm.DB
}

// Start ...
func (serve *Server) Start() error {
	if err := serve.configureLogger(); err != nil {
		return err
	}

	serve.logger.Info("Logger configured!")

	if err := serve.configureDatabase(); err != nil {
		return err
	}

	serve.logger.Info("Database connected!")

	serve.configureRouter()

	serve.logger.Info("Router configured!")

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
	sqlClient, err := gorm.Open(postgres.Open(serve.config.sqlURL), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlClient.Logger.LogMode(1)

	serve.sql = sqlClient

	if err := serve.sql.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	return nil
}

// Configure router ...
func (serve *Server) configureRouter() {
	serve.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(serve.config.Static))))

	serve.router.Use(serve.setRequestID)
	serve.router.Use(serve.logRequest)
	serve.router.Use(serve.notFound)
	serve.router.Use(serve.jwtAuthentication)

	serve.router.HandleFunc("/", serve.getPage("Home")).Methods("GET", "OPTIONS")
	serve.router.HandleFunc("/signin", serve.getPage("SignIn")).Methods("GET", "OPTIONS")
	serve.router.HandleFunc("/signup", serve.getPage("SignUp")).Methods("GET", "OPTIONS")

	serve.router.HandleFunc("/signin", serve.signIn()).Methods("POST")
	serve.router.HandleFunc("/signup", serve.signUp()).Methods("POST")
	serve.router.HandleFunc("/logout", serve.logout()).Methods("GET")
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}
