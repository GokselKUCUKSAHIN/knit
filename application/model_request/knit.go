package model_request

import "knit2024/infrastructure/custom_error"

type KnitRequest struct {
	Image     [][]uint8 `json:"image"`
	PinCount  int       `json:"pinCount"`
	LineLimit int       `json:"lineLimit"`
}

func (request KnitRequest) Validate() error {
	if len(request.Image) == 0 {
		return custom_error.BadRequestErr("image can-not be empty")
	}
	if request.PinCount < 0 {
		return custom_error.BadRequestErr("pin count can-not be negative")
	}
	if request.PinCount < 100 {
		return custom_error.BadRequestErr("pin count can-not be less than 100")
	}
	if request.LineLimit < 0 {
		return custom_error.BadRequestErr("line limit count can-not be negative")
	}
	if request.LineLimit < 500 {
		return custom_error.BadRequestErr("line limit can-not be less than 500")
	}
	return nil
}
