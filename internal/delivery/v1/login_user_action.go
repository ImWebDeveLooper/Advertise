package v1

import (
	"agahi/internal/delivery/dto"
	"agahi/internal/delivery/response"
	"agahi/internal/entity/users"
	"agahi/internal/platform/msg"
	"agahi/internal/platform/utilities"
	"agahi/internal/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var loggedInUser = make(map[string]bool)

func LoginUserAction(repo users.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lDto dto.LoginUserRequest
		err := c.ShouldBindJSON(&lDto)
		if err != nil {
			response.BadRequest(c)
			response.ErrorResponse(c, response.Response{ErrorMessage: err.Error()})
			log.Error(err)
			return
		}
		if loggedInUser[lDto.Email] {
			response.BadRequest(c)
			response.ErrorResponse(c, response.Response{ErrorMessage: msg.LoggedInMessage})
			return
		}
		err = usecase.Login(repo, lDto)
		if err != nil {
			response.InternalServerError(c)
			response.ErrorResponse(c, response.Response{ErrorMessage: err.Error()})
			log.Error(err)
			return
		}
		//Generate JWT Token
		tokenString, jErr := utilities.CreateJWTToken(lDto.Email, utilities.JWT_SECRET)
		if jErr != nil {
			response.InternalServerError(c)
			response.ErrorResponse(c, response.Response{ErrorMessage: err.Error()})
			log.Error(err)
			return
		}
		loggedInUser[lDto.Email] = true
		response.OKResponse(c, response.Response{Message: msg.LoginMessage, Data: gin.H{"token": tokenString}})
	}
}
