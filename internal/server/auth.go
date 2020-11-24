package server

import (
	"time"
	"encoding/json"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/ikilonchic/WEB_LAB-3-4/internal/models"
)

func (serve *Server) signUp() http.HandlerFunc {
	type request struct {
		UserName		string		`json:"username"`
		Email    		string    	`json:"email"`
		Password 		string    	`json:"password"`
		Number			string		`json:"number"`
		Male 	 		string	  	`json:"male"`
		Country			string		`json:"country"`
		DateOfBirth	    time.Time 	`json:"dob"`
		About			string		`json:"about"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		decReq := &request{}

		if err := json.NewDecoder(req.Body).Decode(decReq); err != nil {
			Error(res, http.StatusBadRequest, err)
			return
		}

		checkUser := &models.User{}

		if result := serve.sql.Where("email = ? OR user_name = ?", decReq.Email, decReq.Email).First(&checkUser); result.Error == nil {
			Respond(res, http.StatusUnauthorized, map[string]interface{}{"error": "User has been creating"})
			return
		}

		male, _ := strconv.ParseBool(decReq.Male)

		user := &models.User{
			UserName:	   decReq.UserName,
			Email:         decReq.Email,
			Password:      decReq.Password,
			Number:   	   decReq.Number,
			Country:	   decReq.Country,	
			Male:		   male,
			DateOfBirth:   decReq.DateOfBirth,
			About:		   decReq.About,
		}

		if result := serve.sql.Create(user); result.Error != nil {
			Error(res, http.StatusUnprocessableEntity, result.Error)
			return
		}

		tk := &models.Token{UserID: user.ID}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(serve.config.tokenPassword))
	
		cookie := &http.Cookie{
			Name: "authorization",
			Value: tokenString,
			Path: "/",
			Expires: time.Now().AddDate(0, 0, 1),
			HttpOnly: true,
			MaxAge: 86400,
		}

		http.SetCookie(res, cookie)
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
}

func (serve *Server) signIn() http.HandlerFunc {
	type request struct {
		Email    		string    	`json:"email"`
		Password 		string    	`json:"password"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		decReq := &request{}

		if err := json.NewDecoder(req.Body).Decode(decReq); err != nil {
			Error(res, http.StatusBadRequest, err)
			return
		}

		user := &models.User{}

		if result := serve.sql.Where("email = ? OR user_name = ?", decReq.Email, decReq.Email).First(&user); result.Error != nil {
			Respond(res, http.StatusUnauthorized, map[string]interface{}{"error": "User not found"})
			return
		}

		if isAuth := user.ComparePassword(decReq.Password); isAuth == true {
			tk := &models.Token{UserID: user.ID}
			token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
			tokenString, _ := token.SignedString([]byte(serve.config.tokenPassword))
		
			cookie := &http.Cookie{
				Name: "authorization",
				Value: tokenString,
				Path: "/",
				Expires: time.Now().AddDate(0, 0, 1),
				HttpOnly: true,
				MaxAge: 86400,
			}

			http.SetCookie(res, cookie)
			http.Redirect(res, req, "/", http.StatusSeeOther)
		}

		Respond(res, http.StatusUnauthorized, map[string]interface{}{"error": "User not found"})
		return
	}
}

func (serve *Server) logout() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		cookie := &http.Cookie{
			Name: "authorization",
			Value: "",
			Path: "/",
			HttpOnly: true,
			MaxAge: -1,
		}

		http.SetCookie(res, cookie)
		http.Redirect(res, req, "/signin", http.StatusSeeOther)

		Respond(res, http.StatusUnauthorized, map[string]interface{}{"error": "User not found"})
		return
	}
}