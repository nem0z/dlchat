package types

type Handler func(params []string) *Response

type Request struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type Response struct {
	Result string `json:"result"`
	Err    string `json:"err"`
}
