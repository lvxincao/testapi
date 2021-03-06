package lang

// 数据结构
type JsonData struct {
	Routers struct {
		Filter struct {
			Ip struct {
				Info01 string `json:"info_01"`
				Err01  string `json:"err_01"`
			} `json:"ip"`
			Token string `json:"token"`
		} `json:"filter"`
	} `json:"routers"`
	Models struct {
		Users struct {
			SelectInfo01 string `json:"select_info_01"`
			SelectInfo02 string `json:"select_info_02"`
		} `json:"users"`
	} `json:"models"`
	Controllers struct {
		Users struct {
			ReturnInfo01 string `json:"return_info_01"`
		} `json:"users"`
	} `json:"controllers"`
}
