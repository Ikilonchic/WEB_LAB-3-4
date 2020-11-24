package server

import (
	"os"
	"io/ioutil"
	"time"
	"net/http"
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/ikilonchic/WEB_LAB-3-4/internal/models"
)

type ctxKey int8

const (
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

func (serve *Server) notFound(next http.Handler) http.Handler {
	pages := []string{
		"/",
		"/signin",
		"/signup",
		"/logout",
		"/static/style/SignIn.css",
		"/static/style/SignUp.css",
		"/static/style/Error.css",
		"/static/img/1.jpg",
		"/static/img/2.jpg",
		"/static/img/3.jpg",
		"/static/img/4.jpg",
		"/static/img/5.jpg",
		"/static/img/6kg.gif",
		"/static/img/YAg6.gif",
		"/static/script/forms.js",
		"/static/script/logout.js",
	}

	file, err := os.Open(serve.config.Templates + "404.html")
	if err != nil {
		serve.logger.Error("Cannot find file ./" + serve.config.Templates + "404.html")
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		serve.logger.Error("Cannot read file ./" + serve.config.Templates + "404.html")
	}
	
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		requestPath := req.URL.Path 

		for _, value := range pages {
			if value == requestPath {
				next.ServeHTTP(res, req)
				return
			}
		}

		if _, err := res.Write(data); err == nil {
			res.Header().Set("Content-Type", "text/html")
		}
	})
}

// Request ID ...
func (serve *Server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		id := uuid.New().String()
		res.Header().Set("X-Request-ID", id)
		next.ServeHTTP(res, req.WithContext(context.WithValue(req.Context(), ctxKeyRequestID, id)))
	})
}

// Logging request ...
func (serve *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := serve.logger.WithFields(logrus.Fields{
			"remote_addr": req.RemoteAddr,
			"request_id":  req.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", req.Method, req.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, req)

		var level logrus.Level
		
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}

		logger.Logf(level, "completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
	})
}

// JWT auth ...
func (serve *Server) jwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		notAuth := []string{
			"/signin",
			"/signup",
			"/static/style/SignIn.css",
			"/static/style/SignUp.css",
			"/static/style/Error.css",
			"/static/img/1.jpg",
			"/static/img/2.jpg",
			"/static/img/3.jpg",
			"/static/img/4.jpg",
			"/static/img/5.jpg",
			"/static/img/6kg.jpg",
			"/static/script/forms.js",
			"/static/script/logout.js",
		}
		
		requestPath := req.URL.Path //текущий путь запроса

		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(res, req)
				return
			}
		}

		tokenHeader, err := req.Cookie("authorization") //Получение токена

		if err != nil { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			http.Redirect(res, req, "/signin", http.StatusSeeOther)
			return
		}

		splitted := strings.Split(tokenHeader.String(), "=") //Токен поставляется в формате `authorization=token`, мы проверяем, соответствует ли полученный токен этому требованию
		if len(splitted) != 2 {
			http.Redirect(res, req, "/signin", http.StatusSeeOther)
			return
		}

		tokenPart := splitted[1] //Получаем вторую часть токена
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(serve.config.tokenPassword), nil
		})

		if err != nil { //Неправильный токен, как правило, возвращает 403 http-код
			http.Redirect(res, req, "/signin", http.StatusSeeOther)
			return
		}

		if !token.Valid { //токен недействителен, возможно, не подписан на этом сервере
			http.Redirect(res, req, "/signin", http.StatusSeeOther)
			return
		}

		//Всё прошло хорошо, продолжаем выполнение запроса
		next.ServeHTTP(res, req.WithContext(context.WithValue(req.Context(), ctxKeyUser, tk.UserID))) //передать управление следующему обработчику!
	});
}