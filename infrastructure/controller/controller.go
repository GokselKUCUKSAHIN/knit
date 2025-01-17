package controller

import (
	"github.com/labstack/echo/v4"
	"knit2024/application/knit"
	"knit2024/application/model_request"
	"net/http"
)

type knitController struct {
}

func NewKnitController(
	echo *echo.Echo,
) {
	controller := &knitController{}
	controller.register(echo)
}

func (controller *knitController) register(e *echo.Echo) {
	e.POST("/knit-image", controller.KnitImage)
}

// KnitImage godoc
// @tags knit
// @Accept json
// @Produce json
// @Param request body model_request.KnitRequest true "KnitRequest"
// @Success 200 {array} model_response.KnitResponse
// @Failure 400 {object} custom_error.CustomError
// @Failure 404 {object} custom_error.CustomError
// @Router /knit-image [post]
func (controller *knitController) KnitImage(c echo.Context) error {
	var requestBody model_request.KnitRequest
	err := c.Bind(&requestBody)
	if err != nil {
		return err
	}
	if err = requestBody.Validate(); err != nil {
		return err
	}
	result := knit.Create(
		requestBody.Image,
		requestBody.PinCount,
		requestBody.LineLimit,
		GetImageWidth(requestBody.Image),
		GetImageHeight(requestBody.Image),
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func GetImageHeight(image [][]uint8) int16 {
	return int16(len(image))
}

func GetImageWidth(image [][]uint8) int16 {
	return int16(len(image[0]))
}
