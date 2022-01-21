package handlers

import (
	"net/http"

	"DarkSide1710/CRUD-go-gin/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

// GetAllContacts godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         contact
// @Router       /contact [get]
// @Accept       json
// @Produce      json
// @Success      200  {object}  view.R{data=models.Contactlist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
func (h *Handlers) GetAllContacts(ctx *gin.Context) {
	var response Response

	cList, err := h.Platform.Contact().GetAllContacts()
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

// GetContact    godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         contact
// @Accept       json
// @Produce      json
// @Router 		 /contact/{id} [get]
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  view.R{data=models.Contactlist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
func (h *Handlers) GetContact(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	cList, err := h.Platform.Contact().GetContact(id)
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

// CreateContact godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         contact
// @Router       /contact [post]
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Contact ID"
// @Success      200  {object}  view.R{data=models.Contactlist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
func (h *Handlers) CreateContact(ctx *gin.Context) {
	var response Response
	var request contactRequest
	ctx.Header("Content-Type", "application/json")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	//vizivaem validate
	//if !request.validate() {
	//
	//}

	cList, code, err := h.Platform.Contact().CreateContact(request.toModel())
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

// UpdateContact godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         contact
// @Router       /contact/{id} [put]
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Contact ID"
// @Success      200  {object}  view.R{data=models.Contactlist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
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

	cList, err := h.Platform.Contact().UpdateContact(request)
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

// DeleteContact godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         contact
// @Router       /contact/{id} [delete]
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Contact ID"
// @Success      200  {object}  view.R{data=models.Contactlist}
// @Failure      404  {object}  view.R
// @Failure      500  {object}  view.R
func (h *Handlers) DeleteContact(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := h.Platform.Contact().DeleteContact(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
