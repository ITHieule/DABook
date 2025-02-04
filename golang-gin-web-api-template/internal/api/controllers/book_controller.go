package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/response"
)

type BookController struct {
	*BaseController
}

var Book = &BookController{}

func (c *BookController) Login(ctx *gin.Context) {
	var requestParams request.User
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Dữ liệu không hợp lệ. Vui lòng kiểm tra lại thông tin đăng nhập.")
		return
	}

	token, err := services.User.Login(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, "Đăng nhập thất bại. Vui lòng thử lại sau.")
		return
	}

	response.OkWithData(ctx, gin.H{
		"message": "Đăng nhập thành công!",
		"data":    token,
	})
}
func (c *BookController) Register(ctx *gin.Context) {
	var requestParams request.User
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Dữ liệu không hợp lệ. Vui lòng kiểm tra lại thông tin đăng ký.")
		return
	}

	result, err := services.User.Register(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, "Đăng ký thất bại. Vui lòng thử lại sau.")
		return
	}

	response.OkWithData(ctx, gin.H{
		"message": "Đăng ký thành công!",
		"data":    result,
	})
}
func (c *BookController) GetDaTa(ctx *gin.Context) {
	var requestParams request.User
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Dữ liệu không hợp lệ. Vui lòng kiểm tra lại thông tin đăng ký.")
		return
	}

	result, err := services.User.Register(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, "Đăng ký thất bại. Vui lòng thử lại sau.")
		return
	}

	response.OkWithData(ctx, gin.H{
		"message": "Đăng ký thành công!",
		"data":    result,
	})
}
