package datastruct

// API 格式
type API struct {
	ErrorCode string      `json:"ErrorCode"`
	ErrorText string      `json:"ErrorText"`
	Data      interface{} `json:"Data"`
}
type APIArray struct {
	ErrorCode string `json:"error_code"`
	ErrorText string `json:"error_text"`
	Data      string `json:"data"`
}
type ErrAPI struct {
	// ErrorCode string      `json:"error_code"`
	ErrorText string `json:"error_text"`
	// Data      interface{} `json:"data"`
}

type ResApi struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}

type DataApi struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type RestaurantApi struct {
	RName int         `json:"RName"`
	Data  interface{} `json:"Data"`
}
