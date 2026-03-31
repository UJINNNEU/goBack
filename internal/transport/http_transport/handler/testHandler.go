package handler

import (
	"backend/internal/service/test"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testService service.TestService
}

func NewTestHandler(testService service.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

func (h *TestHandler) RegisterRoutes(router *gin.Engine) {
	test := router.Group("/api/test")
	{
		//test.GET("/:id", h.GetTestByID)
		test.GET("/:UserId", h.GetAvailableTests)
	}
}

func (h *TestHandler) GetTestByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
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
	id, err := strconv.Atoi(c.Param("UserId"))
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

/*
func (h *TestHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.userService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}*/
