package model_response

import "knit2024/application/model_domain"

type KnitResponse struct {
	Pins   []*model_domain.Vector2[int16] `json:"pins"`
	Lines  []*model_domain.Line[int16]    `json:"lines"`
	Width  int16                          `json:"width"`
	Height int16                          `json:"height"`
}

func NewKnitResponse(pins []*model_domain.Vector2[int16], linePairs model_domain.Lines, width, height int16) *KnitResponse {
	lines := make([]*model_domain.Line[int16], 0, len(linePairs))
	for _, line := range linePairs {
		lines = append(lines, &model_domain.Line[int16]{
			Start: line.First(),
			End:   line.Second(),
		})
	}
	return &KnitResponse{
		Pins:   pins,
		Lines:  lines,
		Width:  width,
		Height: height,
	}
}
