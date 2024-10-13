package dto

type (
	RoomReq struct {
		Name string `json:"name"`
	}

	RoomResp struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
