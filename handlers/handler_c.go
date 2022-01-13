package handlers

import (
	"example/web-service-gin/models"
	"net/http"

	_ "example/web-service-gin/models"
	"example/web-service-gin/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Platform *services.Services
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRouter(platform *services.Services) *gin.Engine {
	router := gin.Default()

	h := Handlers{Platform: platform}

	cList := router.Group("/cList")
	cList.GET("/", h.GetAllContacts)
	cList.GET("/:id", h.GetContact)
	cList.POST("/", h.CreateContact)
	cList.PUT("/:id", h.UpdateContact)
	cList.DELETE("/:id", h.DeleteContact)

	router.GET("/", h.HelloWorld)

	return router
}

func (h *Handlers) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

func (h *Handlers) GetAllContacts(ctx *gin.Context) {
	var response Response

	cList, err := h.Platform.GetAllContacts()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = cList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) GetContact(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	cList, err := h.Platform.GetContact(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = cList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) CreateContact(ctx *gin.Context) {
	var response Response
	var request models.Contactlist
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	cList, code, err := h.Platform.CreateContact(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		ctx.JSON(code, response)
		return
	}

	response.Data = cList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) UpdateContact(ctx *gin.Context) {
	var response Response
	var request models.Contactlist
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	request.ID = id

	cList, err := h.Platform.UpdateContact(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = cList
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) DeleteContact(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := h.Platform.DeleteContact(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
