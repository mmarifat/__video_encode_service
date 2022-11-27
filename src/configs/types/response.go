package types

type ResponseObject struct {
	Nonce      int    `json:"nonce"`
	Statuscode int    `json:"status"`
	Message    string `json:"msg"`
	Payload    any    `json:"payload"`
}

type ErrorObject struct {
	Nonce      int    `json:"nonce"`
	Statuscode int    `json:"status"`
	Message    string `json:"msg"`
	Error      any    `json:"error"`
}
