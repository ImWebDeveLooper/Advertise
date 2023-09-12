package v1

import (
	"agahi/internal/delivery/dto"
	"agahi/internal/delivery/response"
	"agahi/internal/entity/users"
	"agahi/internal/platform/msg"
	"agahi/internal/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RegisterUserAction(repo users.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var uDto dto.RegisterUserRequest
		err := c.ShouldBindJSON(&uDto)
		if err != nil {
			response.BadRequest(c)
			response.ErrorResponse(c, response.Response{ErrorMessage: err.Error()})
			log.Error(err)
			return
		}
		err = usecase.Register(repo, uDto)
		if err != nil {
			response.InternalServerError(c)
			response.ErrorResponse(c, response.Response{ErrorMessage: err.Error()})
			log.Error(err)
			return
		}
		response.CreateResponse(c, response.Response{Message: msg.RegisterMessage})
	}
}
