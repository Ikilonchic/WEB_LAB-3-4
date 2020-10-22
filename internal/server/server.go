package server

import (
	"net/http"
	"net/smtp"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ctxKey int64

// Server ...
type Server struct {
	config 		*Config
	logger 		*logrus.Logger
	router 		*mux.Router
	mailClient  *smtp.Client
}

// Start ...
func (serve *Server) Start() error {
	if err := serve.configureLogger(); err != nil {
		return err
	}

	serve.configureRouter()
	
	if err := serve.configureMail(); err != nil {
		return err
	}

	serve.logger.Infof("Server are starting on port%s ...", serve.config.Port)

	return http.ListenAndServe(serve.config.Port, serve.router)
}

// Configure server //

// Configure logger ...
func (serve *Server) configureLogger() error {
	level, err := logrus.ParseLevel(serve.config.LogLevel)
	if err != nil {
		return err
	}

	serve.logger.SetLevel(level)
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

// Configure mail ...
func (serve *Server) configureMail() error {
	client, err := smtp.Dial(serve.config.MailHost)
	if err != nil {
		return err
	}

	return nil
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}
