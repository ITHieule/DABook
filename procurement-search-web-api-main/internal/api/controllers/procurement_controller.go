package controllers

import (
	"net/http"
	"web-api/internal/api/services"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/response"

	"github.com/gin-gonic/gin"
)

type ProcurementController struct {
	*BaseController
}

var Procurement = &ProcurementController{}

//func (c *ProcurementController) GetAll(ctx *gin.Context) {
//	var requestParams request.ProcurementRequest
//	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
//		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
//		return
//	}
//	result, err := services.Procurement.GetDataService(&requestParams)
//	if err != nil {
//		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
//		return
//	}
//	response.OkWithData(ctx, result)
//}

func (c *ProcurementController) GetAverage(ctx *gin.Context) {
	var requestParams request.ProcurementRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Procurement.GetAverageService(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}

func (c *ProcurementController) GetHelloWorld(ctx *gin.Context) {
	// Không cần xử lý request parameters
	result, err := services.Procurement.GetHelloService()
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, "Error fetching data: "+err.Error())
		return
	}

	// Trả về dữ liệu thành công
	response.OkWithData(ctx, result)
}

func (c *ProcurementController) GetAll(ctx *gin.Context) {
	var requestParams request.ProcurementRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Procurement.GetDataService(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}

func (c *ProcurementController) AccoutUser(ctx *gin.Context) {
	var requestParams request.Busers
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Procurement.GetAccoutService(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}

func (c *ProcurementController) SearchUser(ctx *gin.Context) {
	var requestParams request.Busers
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Procurement.SearchDataByUserIDService(requestParams.UserID)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}
func (c *ProcurementController) DeleteUser(ctx *gin.Context) {
	var requestParams request.Busers
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	services.Procurement.DelteDataByUserIDService(requestParams.UserID)
	response.Ok(ctx)
}
func (c *ProcurementController) UpdateUser(ctx *gin.Context) {
	var requestParams request.Busers
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}
	result, err := services.Procurement.UpdataDataByUserIDService(&requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
	}
	response.OkWithData(ctx, result)
}
