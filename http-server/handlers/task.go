package handlers

import (
	"net/http"

	"DarkSide1710/CRUD-go-gin/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetAllTasks(ctx *gin.Context) {
	var response Response

	tList, err := h.Platform.Task().GetAllTasks()
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

func (h *Handlers) GetTask(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	tList, err := h.Platform.Task().GetTask(id)
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

func (t *Handlers) CreateTask(ctx *gin.Context) {
	var response Response
	var request models.Tasklist
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tList, code, err := t.Platform.Task().CreateTask(request)
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

func (t *Handlers) UpdateTask(ctx *gin.Context) {
	var response Response
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

	tList, err := t.Platform.Task().UpdateTask(request)
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

func (t *Handlers) DeleteTask(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := t.Platform.Task().DeleteTask(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
