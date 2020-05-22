package controller

import (
	"net/http"

	"github.com/dammy001/schgo/api/response"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
