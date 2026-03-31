package http_transport

import (
	"backend/internal/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(errInvalidUserId))
		return
	}
	if id <= 0 {
		c.JSON(http.StatusBadRequest, newErrResp(errInvalidUserId))
		return
	}

	user, err := s.user.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, newErrResp(err))
			return
		}
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) ListUsers(c *gin.Context) {
	users, err := s.user.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	c.JSON(http.StatusOK, users)
}
