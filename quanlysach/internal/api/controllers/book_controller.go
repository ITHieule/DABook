package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/response"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	*BaseController
}

var Book = &BookController{}

func (c *BookController) GetBook(ctx *gin.Context) {
	result, err := services.Book.GetDataService()
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *BookController) AddBook(ctx *gin.Context) {
	var requestParams request.Book

	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Book.AddDataService(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
func (c *BookController) DeleteBook(ctx *gin.Context) {
	var requestParams request.Book

	fmt.Print("ID : ", requestParams.ID)

	// Kiểm tra và parse request body
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Chuyển đổi ID từ string sang int
	id, err := strconv.Atoi(requestParams.ID)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "ID phải là một số nguyên hợp lệ")
		return
	}
	services.Book.DeleteDataService(id)
	// Phản hồi thành công
	response.Ok(ctx)
}
func (c *BookController) UpdateBook(ctx *gin.Context) {
	var requestParams request.Book
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Book.UpdateDataService(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}
func (c *BookController) OderBook(ctx *gin.Context) {
	var requestParams request.Oder
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Book.OderDataService(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}
func (c *BookController) Oderstat(ctx *gin.Context) {
	result, err := services.Book.GetOrderStatistics()
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
func (c *BookController) SearchBook(ctx *gin.Context) {
	var requestParams request.Book
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Book.SearchIDService(requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}
func (c *BookController) Register(ctx *gin.Context) {
	var requestParams request.User
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.User.Register(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}
func (c *BookController) Login(ctx *gin.Context) {
	var requestParams request.User
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	token, err := services.User.Login(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, token)

}
