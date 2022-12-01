package types

type ResponseObject struct {
	Nonce      int64  `json:"nonce"`
	Statuscode int    `json:"status" default:"200"`
	Message    string `json:"message"`
	Payload    any    `json:"payload"`
}

type ErrorObject struct {
	Nonce      int64  `json:"nonce"`
	Statuscode int    `json:"status" default:"400"`
	Message    string `json:"message"`
	Error      any    `json:"error"`
}
