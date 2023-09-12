package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data         interface{} `json:"data,omitempty"`
	Message      string      `json:"message,omitempty"`
	Error        error       `json:"-,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func ErrorResponse(c *gin.Context, resp Response) {
	jData, err := json.Marshal(resp)
	if err != nil {
		BadRequest(c)
		return
	}
	c.Header("Content-Type", "application/json")
	_, err = c.Writer.Write(jData)
	if err != nil {
		return
	}
}
func BadRequest(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusBadRequest)
}

func InternalServerError(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusInternalServerError)
}

func CreateResponse(c *gin.Context, resp Response) {
	jData, err := json.Marshal(resp)
	if err != nil {
		BadRequest(c)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)
	_, err = c.Writer.Write(jData)
	if err != nil {
		return
	}
}

func OKResponse(c *gin.Context, resp Response) {
	jData, err := json.Marshal(resp)
	if err != nil {
		BadRequest(c)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
	_, err = c.Writer.Write(jData)
	if err != nil {
		return
	}
}

func DeleteResponse(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusNoContent)
}

func NotFound(c *gin.Context) {
	resp, err := json.Marshal(Response{Message: "404 Not found"})
	if err != nil {
		BadRequest(c)
		return
	}
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusNotFound)
	_, err = c.Writer.Write(resp)
	return
}
