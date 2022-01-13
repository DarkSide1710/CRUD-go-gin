package handlers

import (
	"example/web-service-gin/models"
	_ "example/web-service-gin/models"
	"example/web-service-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlersT struct {
	Platform *services.ServicesT
}

type ResponseT struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewrouterT(platform *services.ServicesT) *gin.Engine {
	router := gin.Default()

	h := HandlersT{Platform: platform}

	tList := router.Group("/t_list")
	tList.GET("/", h.GetAllTasks)
	tList.GET("/:id", h.GetTask)
	tList.POST("/", h.CreateTask)
	tList.PUT("/:id", h.UpdateTask)
	tList.DELETE("/:id", h.DeleteTask)

	router.GET("/", h.HelloWorld)

	return router
}

func (h *HandlersT) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

func (h *HandlersT) GetAllTasks(ctx *gin.Context) {
	var response ResponseT

	tList, err := h.Platform.GetAllTasks()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = tList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *HandlersT) GetTask(ctx *gin.Context) {
	var response ResponseT
	id := ctx.Param("id")

	tList, err := h.Platform.GetTask(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = tList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *HandlersT) CreateTask(ctx *gin.Context) {
	var response ResponseT
	var request models.Tasklist
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tList, code, err := h.Platform.CreateTask(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		ctx.JSON(code, response)
		return
	}

	response.Data = tList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *HandlersT) UpdateTask(ctx *gin.Context) {
	var response ResponseT
	var request models.Tasklist
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	request.ID = id

	tList, err := h.Platform.UpdateTask(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = tList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *HandlersT) DeleteTask(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := h.Platform.DeleteTask(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
