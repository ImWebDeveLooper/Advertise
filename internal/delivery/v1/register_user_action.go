package v1

import (
	"agahi/internal/delivery/dto"
	"agahi/internal/delivery/response"
	"agahi/internal/entity/users"
	"agahi/internal/platform/msg"
	"agahi/internal/usecase"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RegisterUserAction(repo users.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var uDto dto.RegisterUserRequest
		err := json.NewDecoder(req.Body).Decode(&uDto)
		if err != nil {
			response.BadRequest(w)
			response.ErrorResponse(w, response.Response{ErrorMessage: err.Error()})
			log.Error(err)
			return
		}
		err = usecase.Register(repo, uDto)
		if err != nil {
			response.InternalServerError(w)
			response.ErrorResponse(w, response.Response{ErrorMessage: err.Error()})
			log.Error(err)
			return
		}
		response.CreateResponse(w, response.Response{Message: msg.RegisterMessage})
	}
}
