package handlers

import (
	"net/http"

	"DarkSide1710/CRUD-go-gin/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetAllTasks godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         task
// @Router       /task [get]
// @Accept       json
// @Produce      json
// @Success      200  {object}  view.R{data=models.Tasklist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
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

// GetTask godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         task
// @Router       /task/{id} [get]
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  view.R{data=models.Tasklist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
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

// CreateTask godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         task
// @Router       /task [post]
// @Accept       json
// @Produce      json
// @Param        params   body taskRequest true  "Params"
// @Success      200  {object}  view.R{data=models.Tasklist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
func (t *Handlers) CreateTask(ctx *gin.Context) {
	var response Response
	var request taskRequest
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tList, code, err := t.Platform.Task().CreateTask(request.toModel_1())
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

// UpdateTask godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         task
// @Router       /task/{id} [put]
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Contact ID"
// @Success      200  {object}  view.R{data=models.Tasklist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
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

// DeleteTask godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         task
// @Router       /task/{id} [delete]
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  view.R{data=models.Tasklist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
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
