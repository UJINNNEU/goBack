package handler

import (
	"backend/internal/service/serviceT"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testService *serviceT.TestService
}

func NewTestHandler(testService *serviceT.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

func (h *TestHandler) GetTestByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id test"})
		return
	}

	test, err := h.testService.GetTestById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, test)
}

func (h *TestHandler) GetAvailableTests(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "invalid id_user"})
		return
	}
	tests, err := h.testService.GetAvailableTests(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tests)
}
