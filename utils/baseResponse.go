package utils

type BaseResp struct {
	StatusCode int32  `json:"status_code"` // 状态码 0：成功；其他：失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}
