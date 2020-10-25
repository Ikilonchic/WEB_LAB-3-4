package server

import (
	"time"
	"encoding/json"
	"net/http"

	"github.com/ikilonchic/WEB_LAB-3-4/internal/models"
)

func (serve *Server) signUp() http.HandlerFunc {
	type request struct {
		UserName		string		`json:"username"`
		Email    		string    	`json:"email"`
		Password 		string    	`json:"password"`
		Number			string		`json:"number"`
		Male 	 		bool	  	`json:"male"`
		DateOfBirth	    time.Time 	`json:"dof"`
		About			string		`json:"about"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		decReq := &request{}

		if err := json.NewDecoder(req.Body).Decode(decReq); err != nil {
			serve.error(res, req, http.StatusBadRequest, err)
			return
		}

		user := &models.User{
			Email:    decReq.Email,
			Password: decReq.Password,
		}

		//if err := serve.sql.User().Create(user); err != nil {
		//	serve.error(res, req, http.StatusUnprocessableEntity, err)
		//	return
		//}

		user.Sanitize()
		serve.respond(res, req, http.StatusCreated, user)
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
			serve.error(res, req, http.StatusBadRequest, err)
			return
		}
	}
}